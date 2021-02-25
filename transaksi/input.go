package transaksi

import "premium/user"

type GetProdukTransaksisInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

type CreateTransaksiInput struct {
	Harga      int    `json:"harga" binding:"required"`
	ProdukID   int    `json:"produk_id" binding:"required"`
	PaymentURL string `json:"payment_url" binding:"required`
	User       user.User
}
