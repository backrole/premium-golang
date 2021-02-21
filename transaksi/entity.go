package transaksi

import (
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
	CreatedAt time.Time
	UpdatedAt time.Time
}
