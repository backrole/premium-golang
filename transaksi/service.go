package transaksi

import (
	"errors"
	"premium/produk"
)

type service struct {
	repo       Repo
	produkRepo produk.Repo
}

type Service interface {
	GetTransaksiByProdukID(input GetProdukTransaksisInput) ([]Transaksi, error)
}

func NewService(repo Repo, produkRepo produk.Repo) *service {
	return &service{repo, produkRepo}
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
