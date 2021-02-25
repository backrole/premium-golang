package payment

import (
	"log"
	"premium/user"
	"strconv"

	"github.com/veritrans/go-midtrans"
)

type service struct {
}

type Service interface {
	GetPaymentURL(transaksi Transaksi, user user.User) (string, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentURL(transaksi Transaksi, user user.User) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = ""
	midclient.ClientKey = ""
	midclient.APIEnvType = midtrans.Sandbox

	var snapGateway midtrans.SnapGateway
	snapGateway = midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaksi.ID),
			GrossAmt: int64(transaksi.Harga),
		},
		CustomerDetail: &midtrans.CustDetail{
			FName: user.Nama,
			Email: user.Email,
		},
	}

	log.Println("GetToken:")
	snapTokenResp, _ := snapGateway.GetToken(snapReq)

	return snapTokenResp.RedirectURL, nil
	// snapGateway := midtrans.SnapGateway{
	// 	Client: midclient,
	// }
	// snapReq := &midtrans.SnapReq{
	// 	CustomerDetail: &midtrans.CustDetail{
	// 		Email: user.Email,
	// 		FName: user.Nama,
	// 	},
	// 	TransactionDetails: midtrans.TransactionDetails{
	// 		OrderID:  strconv.Itoa(transaksi.ID),
	// 		GrossAmt: int64(transaksi.Harga),
	// 	},
	// }

	// snapTokenResp, err := snapGateway.GetToken(snapReq)
	// if err != nil {
	// 	return "", err
	// }

	// return snapTokenResp.RedirectURL, nil

}
