package transaksi

import "time"

type Transaksi struct {
	ID        int
	ProdukID  int
	UserID    int
	Harga     int
	Status    string
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
