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

type CreateGambarProdukInput struct {
	ProdukID  int  `form:"produk_id" binding:"required"`
	IsPrimary bool `form:"is_primary"`
	User      user.User
}
