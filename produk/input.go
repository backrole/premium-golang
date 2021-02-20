package produk

type GetProdukDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
