package transaksi

import (
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

type Repo interface {
	GetByProdukID(produkID int) ([]Transaksi, error)
	GetByUserID(userID int) ([]Transaksi, error)
}

func NewRepo(db *gorm.DB) *repo {
	return &repo{db}
}

func (r *repo) GetByProdukID(produkID int) ([]Transaksi, error) {
	var transansis []Transaksi

	err := r.db.Preload("User").Where("produk_id", produkID).Order("id desc").Find(&transansis).Error

	if err != nil {
		return transansis, err
	}

	return transansis, nil
}

func (r *repo) GetByUserID(userID int) ([]Transaksi, error) {
	var transaksis []Transaksi

	err := r.db.Preload("Produk.GambarProduks", "gambar_produks.is_primary = 1").Where("user_id = ?", userID).Find(&transaksis).Error
	if err != nil {
		return transaksis, err
	}
	return transaksis, nil

}
