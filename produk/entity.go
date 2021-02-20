package produk

import "time"

type Produk struct {
	ID            int
	UserID        int
	NamaProduk    string
	Judul         string
	Deskripsi     string
	JumlahPembeli int
	Harga         int
	Slug          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	GambarProduks []GambarProduk
}

type GambarProduk struct {
	ID         int
	ProdukID   int
	NamaGambar string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
