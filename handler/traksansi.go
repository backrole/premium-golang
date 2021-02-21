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
