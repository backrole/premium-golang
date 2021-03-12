package transaksi

import (
	"premium/produk"
	"premium/user"
)

type GetProdukTransaksisInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

type CreateTransaksiInput struct {
	Harga    int `json:"harga" binding:"required"`
	ProdukID int `json:"produk_id" binding:"required"`
	Produk   produk.Produk
	User     user.User
}

type TransaksiNotifikasiInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
