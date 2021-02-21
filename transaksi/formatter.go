package transaksi

import "time"

type ProdukTransaksiFormatter struct {
	ID        int       `json:"id"`
	Nama      string    `json:"nama"`
	Harga     int       `json:"harga"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatProdukTransaksi(transaksi Transaksi) ProdukTransaksiFormatter {
	formatter := ProdukTransaksiFormatter{}
	formatter.ID = transaksi.ID
	formatter.Nama = transaksi.User.Nama
	formatter.Harga = transaksi.Harga
	formatter.CreatedAt = transaksi.CreatedAt
	return formatter
}

func FormatProdukTransaksis(transaksi []Transaksi) []ProdukTransaksiFormatter {
	if len(transaksi) == 0 {
		return []ProdukTransaksiFormatter{}
	}

	var transaksisFormatter []ProdukTransaksiFormatter

	for _, transaksi := range transaksi {
		formatter := FormatProdukTransaksi(transaksi)
		transaksisFormatter = append(transaksisFormatter, formatter)
	}

	return transaksisFormatter
}
