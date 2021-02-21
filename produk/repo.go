package produk

import "gorm.io/gorm"

type Repo interface {
	FindAll() ([]Produk, error)
	FindByUserID(userID int) ([]Produk, error)
	FindByID(ID int) (Produk, error)
	Save(produk Produk) (Produk, error)
	Update(produk Produk) (Produk, error)
	CreateGambar(gambarProduk GambarProduk) (GambarProduk, error)
	MarkAllGambarAsNonPrimary(produkID int) (bool, error)
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

func (r *repo) FindByID(ID int) (Produk, error) {
	var produk Produk

	err := r.db.Preload("User").Preload("GambarProduks").Where("id = ?", ID).Find(&produk).Error
	if err != nil {
		return produk, err
	}

	return produk, nil
}

func (r *repo) Save(produk Produk) (Produk, error) {
	err := r.db.Create(&produk).Error

	if err != nil {
		return produk, err
	}
	return produk, nil
}

func (r *repo) Update(produk Produk) (Produk, error) {
	err := r.db.Save(&produk).Error
	if err != nil {
		return produk, err
	}

	return produk, nil
}

func (r *repo) CreateGambar(gambarProduk GambarProduk) (GambarProduk, error) {
	err := r.db.Create(&gambarProduk).Error

	if err != nil {
		return gambarProduk, err
	}
	return gambarProduk, nil
}

func (r *repo) MarkAllGambarAsNonPrimary(produkID int) (bool, error) {
	err := r.db.Model(&GambarProduk{}).Where("produk_id = ?", produkID).Update("is_primary", false).Error

	if err != nil {
		return false, err
	}
	return true, nil
}
