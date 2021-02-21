package transaksi

import "premium/user"

type GetProdukTransaksisInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
