package produk

import (
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetProduks(userID int) ([]Produk, error)
	GetProdukByID(input GetProdukDetailInput) (Produk, error)
	CreateProduk(input CreateProdukInput) (Produk, error)
}

type service struct {
	repo Repo
}

func NewService(repo Repo) *service {
	return &service{repo}
}

func (s *service) GetProduks(userID int) ([]Produk, error) {
	if userID != 0 {
		produks, err := s.repo.FindByUserID(userID)
		if err != nil {
			return produks, err
		}

		return produks, nil
	}

	produks, err := s.repo.FindAll()
	if err != nil {
		return produks, err
	}

	return produks, nil
}

func (s *service) GetProdukByID(input GetProdukDetailInput) (Produk, error) {
	produk, err := s.repo.FindByID(input.ID)

	if err != nil {
		return produk, err
	}

	return produk, nil
}

func (s *service) CreateProduk(input CreateProdukInput) (Produk, error) {
	produk := Produk{}
	produk.NamaProduk = input.NamaProduk
	produk.Judul = input.Judul
	produk.Deskripsi = input.Deskripsi
	produk.Harga = input.Harga
	produk.Perks = input.Perks
	produk.UserID = input.User.ID

	slugCandidate := fmt.Sprintf("%s %d", input.NamaProduk, input.User.ID)
	produk.Slug = slug.Make(slugCandidate)

	newProduk, err := s.repo.Save(produk)
	if err != nil {
		return newProduk, err
	}

	return newProduk, nil
}
