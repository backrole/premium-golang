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

type UserTransaksiFormatter struct {
	ID         int             `json:"id"`
	Harga      int             `json:"harga"`
	Status     string          `json:"status"`
	PaymentURL string          `json:"payment_url"`
	CreatedAt  time.Time       `json:"created_at"`
	Produk     ProdukFormatter `json:"produk"`
}

type ProdukFormatter struct {
	Nama      string `json:"nama"`
	GambarURL string `json:"gambar_url"`
}

func FormatUserTransaksi(transaksi Transaksi) UserTransaksiFormatter {
	formatter := UserTransaksiFormatter{}
	formatter.ID = transaksi.ID
	formatter.Harga = transaksi.Harga
	formatter.Status = transaksi.Status
	formatter.CreatedAt = transaksi.CreatedAt
	formatter.Produk.Nama = transaksi.Produk.NamaProduk
	produkFormatter := ProdukFormatter{}
	produkFormatter.GambarURL = ""

	if len(transaksi.Produk.GambarProduks) > 0 {
		produkFormatter.GambarURL = transaksi.Produk.GambarProduks[0].NamaGambar
	}
	formatter.Produk = produkFormatter
	return formatter

}

func FormatUserTransaksis(transaksi []Transaksi) []UserTransaksiFormatter {
	if len(transaksi) == 0 {
		return []UserTransaksiFormatter{}
	}

	var transaksisFormatter []UserTransaksiFormatter

	for _, transaksi := range transaksi {
		formatter := FormatUserTransaksi(transaksi)
		transaksisFormatter = append(transaksisFormatter, formatter)
	}

	return transaksisFormatter
}

type TransaksiFormatter struct {
	ID         int    `json:"id"`
	ProdukID   int    `json:"produk_id"`
	UserID     int    `json:"user_id"`
	Harga      int    `json:"harga"`
	Status     string `json:"status"`
	Code       string `json:"code"`
	PaymentURL string `json:"payment_url"`
}

func FormatTransaksi(transaksi Transaksi) TransaksiFormatter {
	formatter := TransaksiFormatter{}
	formatter.ID = transaksi.ID
	formatter.ProdukID = transaksi.ProdukID
	formatter.UserID = transaksi.UserID
	formatter.Harga = transaksi.Harga
	formatter.Status = transaksi.Status
	formatter.Code = transaksi.Code
	formatter.PaymentURL = transaksi.PaymentURL
	return formatter
}
