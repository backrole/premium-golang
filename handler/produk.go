package handler

import (
	"net/http"
	"premium/helper"
	"premium/produk"
	"premium/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type produkHandler struct {
	service produk.Service
}

func NewProdukHandelr(service produk.Service) *produkHandler {
	return &produkHandler{service}
}

func (h *produkHandler) GetProduks(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	produks, err := h.service.GetProduks(userID)

	if err != nil {
		response := helper.APIResponse("Error get produk", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List produk", http.StatusOK, "error", produk.FormatProduks(produks))
	c.JSON(http.StatusOK, response)
}

func (h *produkHandler) GetProduk(c *gin.Context) {
	var input produk.GetProdukDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Gagal get detail produk", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	produkDetail, err := h.service.GetProdukByID(input)
	if err != nil {
		response := helper.APIResponse("Gagal get detail produk", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Detail produk", http.StatusOK, "sukses", produk.FormatProdukDetail(produkDetail))

	c.JSON(http.StatusOK, response)
	return
}

func (h *produkHandler) CreateProduk(c *gin.Context) {
	var input produk.CreateProdukInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("gagal create produk", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	newProduk, err := h.service.CreateProduk(input)
	if err != nil {

		response := helper.APIResponse("gagal create produk", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Sukses create produk", http.StatusOK, "Sukses", produk.FormatProduk(newProduk))
	c.JSON(http.StatusOK, response)

}
