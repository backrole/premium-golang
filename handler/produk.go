package handler

import (
	"net/http"
	"premium/helper"
	"premium/produk"
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
