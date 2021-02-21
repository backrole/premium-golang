package main

import (
	"log"
	"net/http"
	"premium/auth"
	"premium/handler"
	"premium/helper"
	"premium/produk"
	"premium/transaksi"
	"premium/user"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/premium?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepo := user.NewRepo(db)
	produkRepo := produk.NewRepo(db)
	transaksiRepo := transaksi.NewRepo(db)

	userService := user.NewService(userRepo)
	produkService := produk.NewService(produkRepo)
	authService := auth.NewService()
	transaksiService := transaksi.NewService(transaksiRepo, produkRepo)

	userHandler := handler.NewUserHandler(userService, authService)
	produkHandler := handler.NewProdukHandelr(produkService)
	transaksiHandler := handler.NewTransaksiHandler(transaksiService)

	router := gin.Default()
	router.Static("/images", "./images")

	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_ceks", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.GET("/produks", produkHandler.GetProduks)
	api.GET("/produks/:id", produkHandler.GetProduk)
	api.POST("/produks", authMiddleware(authService, userService), produkHandler.CreateProduk)
	api.PUT("/produks/:id", authMiddleware(authService, userService), produkHandler.UpdateProduk)
	api.POST("/gambar-produks", authMiddleware(authService, userService), produkHandler.UploadGambar)
	api.GET("/produk/:id/transaksi", authMiddleware(authService, userService), transaksiHandler.GetProdukTransaksis)
	api.GET("/transaksi", authMiddleware(authService, userService), transaksiHandler.GetUserTransaksis)

	router.Run()
}
func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))
		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
