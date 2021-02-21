package transaksi

import (
	"premium/produk"
	"premium/user"
	"time"
)

type Transaksi struct {
	ID        int
	ProdukID  int
	UserID    int
	Harga     int
	Status    string
	Code      string
	User      user.User
	Produk    produk.Produk
	CreatedAt time.Time
	UpdatedAt time.Time
}
