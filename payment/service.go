package payment

import (
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
	midclient.ServerKey = "SB-Mid-server-lVSb2InwtN6Cv8vy6DqzvZzG"
	midclient.ClientKey = "SB-Mid-client-N-mLFyOCuq1R0FJb"
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

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}
