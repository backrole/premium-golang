package produk

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetProduks(userID int) ([]Produk, error)
	GetProdukByID(input GetProdukDetailInput) (Produk, error)
	CreateProduk(input CreateProdukInput) (Produk, error)
	UpdateProduk(inputID GetProdukDetailInput, inputData CreateProdukInput) (Produk, error)
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
func (s *service) UpdateProduk(inputID GetProdukDetailInput, inputData CreateProdukInput) (Produk, error) {
	produk, err := s.repo.FindByID(inputID.ID)
	if err != nil {
		return produk, err
	}

	if produk.UserID != inputData.User.ID {
		return produk, errors.New("Anda bukan pemilik produk")
	}

	produk.NamaProduk = inputData.NamaProduk
	produk.Judul = inputData.Judul
	produk.Deskripsi = inputData.Deskripsi
	produk.Harga = inputData.Harga
	produk.Perks = inputData.Perks

	updateProduk, err := s.repo.Update(produk)
	if err != nil {
		return updateProduk, err
	}

	return updateProduk, nil
}
