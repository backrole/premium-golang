package produk

import "premium/user"

type GetProdukDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateProdukInput struct {
	NamaProduk string `json:"nama_produk" binding:"required"`
	Judul      string `json:"judul" binding:"required"`
	Deskripsi  string `json:"deskripsi" binding:"required"`
	Harga      int    `json:"harga" binding:"required"`
	Perks      string `json:"perks" binding:"required"`
	User       user.User
}
