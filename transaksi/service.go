package transaksi

import (
	"errors"
	"premium/payment"
	"premium/produk"
	"strconv"
)

type service struct {
	repo           Repo
	produkRepo     produk.Repo
	paymentService payment.Service
}

type Service interface {
	GetTransaksiByProdukID(input GetProdukTransaksisInput) ([]Transaksi, error)
	GetTransaksiByUserID(userID int) ([]Transaksi, error)
	CreateTransaksi(input CreateTransaksiInput) (Transaksi, error)
	ProsesPayment(input TransaksiNotifikasiInput) error
}

func NewService(repo Repo, produkRepo produk.Repo, paymentService payment.Service) *service {
	return &service{repo, produkRepo, paymentService}
}

func (s *service) GetTransaksiByProdukID(input GetProdukTransaksisInput) ([]Transaksi, error) {

	produk, err := s.produkRepo.FindByID(input.ID)
	if err != nil {
		return []Transaksi{}, err
	}

	if produk.UserID != input.User.ID {
		return []Transaksi{}, errors.New("Anda bukan pemilik produk")
	}

	transaksis, err := s.repo.GetByProdukID(input.ID)
	if err != nil {
		return transaksis, err
	}

	return transaksis, nil
}

func (s *service) GetTransaksiByUserID(userID int) ([]Transaksi, error) {
	transaksis, err := s.repo.GetByUserID(userID)
	if err != nil {
		return transaksis, err
	}

	return transaksis, nil
}

func (s *service) CreateTransaksi(input CreateTransaksiInput) (Transaksi, error) {
	transaksi := Transaksi{}
	transaksi.ProdukID = input.ProdukID
	transaksi.Harga = input.Harga
	transaksi.UserID = input.User.ID
	transaksi.Status = "pending"

	newTransaksi, err := s.repo.Save(transaksi)
	if err != nil {
		return newTransaksi, err
	}

	// trans, _ := s.repo.GetByProdukID(transaksi.ProdukID)

	paymentTransaksi := payment.Transaksi{
		ID:    newTransaksi.ID,
		Harga: newTransaksi.Harga,
	}

	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransaksi, input.User)
	if err != nil {
		return newTransaksi, err
	}

	newTransaksi.PaymentURL = paymentURL

	newTransaksi, err = s.repo.Update(newTransaksi)
	if err != nil {
		return newTransaksi, err
	}

	return newTransaksi, nil
}
func (s *service) ProsesPayment(input TransaksiNotifikasiInput) error {
	transaction_id, _ := strconv.Atoi(input.OrderID)

	transaction, err := s.repo.GetByID(transaction_id)
	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		input.TransactionStatus = "paid"
	} else if input.TransactionStatus == "settlement" {
		input.TransactionStatus = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	updatedTransaksi, err := s.repo.Update(transaction)
	if err != nil {
		return err
	}

	produk, err := s.produkRepo.FindByID(updatedTransaksi.ProdukID)
	if err != nil {
		return err
	}
	if updatedTransaksi.Status == "paid" {
		produk.JumlahPembeli = produk.JumlahPembeli + 1

		_, err := s.produkRepo.Update(produk)

		if err != nil {
			return err
		}
	}
	return nil
}
