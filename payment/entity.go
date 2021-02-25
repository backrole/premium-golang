package payment

import "premium/produk"

type Transaksi struct {
	ID     int
	Harga  int
	Produk produk.Produk
}
