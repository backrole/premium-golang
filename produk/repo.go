package produk

import "gorm.io/gorm"

type Repo interface {
	FindAll() ([]Produk, error)
	FindByUserID(userID int) ([]Produk, error)
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repo {
	return &repo{db}
}

func (r *repo) FindAll() ([]Produk, error) {
	var produks []Produk

	err := r.db.Preload("GambarProduks", "gambar_produks.is_primary = 1").Find(&produks).Error
	if err != nil {
		return produks, err
	}

	return produks, nil
}

func (r *repo) FindByUserID(userID int) ([]Produk, error) {
	var produks []Produk
	err := r.db.Where("user_id = ?", userID).Preload("GambarProduks", "gambar_produks.is_primary = 1").Find(&produks).Error

	if err != nil {
		return produks, err
	}

	return produks, nil

}
