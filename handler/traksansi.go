package handler

import (
	"net/http"
	"premium/helper"
	"premium/transaksi"
	"premium/user"

	"github.com/gin-gonic/gin"
)

type transaksiHandler struct {
	service transaksi.Service
}

func NewTransaksiHandler(service transaksi.Service) *transaksiHandler {
	return &transaksiHandler{service}
}

func (h *transaksiHandler) GetProdukTransaksis(c *gin.Context) {
	var input transaksi.GetProdukTransaksisInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Gagal get transaksi produk", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transaksis, err := h.service.GetTransaksiByProdukID(input)
	if err != nil {
		response := helper.APIResponse("Gagal get transaksi produk", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Sukses get transaksi produk", http.StatusOK, "sukses", transaksi.FormatProdukTransaksis(transaksis))
	c.JSON(http.StatusOK, response)
}
func (h *transaksiHandler) GetUserTransaksis(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	transaksis, err := h.service.GetTransaksiByUserID(userID)
	if err != nil {
		response := helper.APIResponse("User gagal transaksi produk", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("User sukses transaksi produk", http.StatusOK, "sukses", transaksi.FormatUserTransaksis(transaksis))
	c.JSON(http.StatusOK, response)
}

func (h *transaksiHandler) CreateTransaksi(c *gin.Context) {
	var input transaksi.CreateTransaksiInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Gagal create transaksi", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newTransaksi, err := h.service.CreateTransaksi(input)
	if err != nil {
		response := helper.APIResponse("Gagal create transaksis", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Sukses create transaksi", http.StatusOK, "Sukses", transaksi.FormatTransaksi(newTransaksi))
	c.JSON(http.StatusOK, response)
	return

}

func (h *transaksiHandler) GetNotifikasi(c *gin.Context) {
	var input transaksi.TransaksiNotifikasiInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		response := helper.APIResponse("Gagal proses notifikasi", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.service.ProsesPayment(input)
	if err != nil {
		response := helper.APIResponse("Gagal proses notifikasi", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	c.JSON(http.StatusOK, input)
}
