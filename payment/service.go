package payment

import (
	"premium/user"
	"strconv"

	midtrans "github.com/veritrans/go-midtrans"
)

type service struct {
}

type Service interface {
	GetPaymetURL(transaksi Transaksi, user user.User) (string, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymetURL(transaksi Transaksi, user user.User) (string, error) {
	// midclient := midtrans.NewClient()
	// midclient.ServerKey = "SB-Mid-client-N-mLFyOCuq1R0FJb"
	// midclient.ClientKey = "SB-tMid-server-lVSb2InwtN6Cv8vy6DqzvZzG"
	// midclient.APIEnvType = midrans.Sandbox

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
	midclient := midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-client-N-mLFyOCuq1R0FJb"
	midclient.ClientKey = "SB-tMid-server-lVSb2InwtN6Cv8vy6DqzvZzG"
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
			Email: user.Email,
			FName: user.Nama,
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}
