package produk

import (
	"premium/user"
	"time"
)

type Produk struct {
	ID            int
	UserID        int
	NamaProduk    string
	Judul         string
	Deskripsi     string
	JumlahPembeli int
	Harga         int
	Perks         string
	Slug          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	GambarProduks []GambarProduk
	User          user.User
}

type GambarProduk struct {
	ID         int
	ProdukID   int
	NamaGambar string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
