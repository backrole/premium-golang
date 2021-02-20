package produk

import "strings"

type ProdukFormatter struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	NamaProduk string `json:"nama_produk"`
	Judul      string `json:"judul"`
	GambarURL  string `json:"gambar_url"`
	Harga      int    `json:"harga"`
	Slug       string `json:"slug"`
}

func FormatProduk(produk Produk) ProdukFormatter {
	produkFormatter := ProdukFormatter{}
	produkFormatter.ID = produk.ID
	produkFormatter.UserID = produk.UserID
	produkFormatter.NamaProduk = produk.NamaProduk
	produkFormatter.Judul = produk.Judul
	produkFormatter.Harga = produk.Harga
	produkFormatter.Slug = produk.Slug
	produkFormatter.GambarURL = ""

	if len(produk.GambarProduks) > 0 {
		produkFormatter.GambarURL = produk.GambarProduks[0].NamaGambar
	}

	return produkFormatter
}

func FormatProduks(produks []Produk) []ProdukFormatter {
	produksFormatter := []ProdukFormatter{}
	for _, produk := range produks {
		produkFormatter := FormatProduk(produk)
		produksFormatter = append(produksFormatter, produkFormatter)
	}

	return produksFormatter
}

type ProdukDetailFormatter struct {
	ID            int                    `json:"id"`
	NamaGambar    string                 `json:"namagambar"`
	Judul         string                 `json:"judul"`
	Deskripsi     string                 `json:"deskripsi"`
	GambarURL     string                 `json:"gambar_url"`
	JumlahPembeli int                    `json:"jumlahpembeli"`
	Harga         int                    `json:"harga"`
	UserID        int                    `json:"user_id"`
	Slug          string                 `json:"slug"`
	Perks         []string               `json:"perks"`
	User          ProdukUserFormatter    `json:"use"`
	Images        []ProdukImageFormatter `json:"images"`
}

type ProdukUserFormatter struct {
	Nama         string `json:"nama"`
	GambarProduk string `json:"gambar_url"`
}

type ProdukImageFormatter struct {
	GambarProduk string `json:"gambar_url"`
	IsPrimary    bool   `json:"is_primary"`
}

func FormatProdukDetail(produk Produk) ProdukDetailFormatter {
	produkDetailFormatter := ProdukDetailFormatter{}
	produkDetailFormatter.ID = produk.ID
	produkDetailFormatter.NamaGambar = produk.NamaProduk
	produkDetailFormatter.Judul = produk.Judul
	produkDetailFormatter.Deskripsi = produk.Deskripsi
	produkDetailFormatter.JumlahPembeli = produk.JumlahPembeli
	produkDetailFormatter.Harga = produk.Harga
	produkDetailFormatter.UserID = produk.UserID
	produkDetailFormatter.Slug = produk.Slug
	produkDetailFormatter.GambarURL = ""

	if len(produk.GambarProduks) > 0 {
		produkDetailFormatter.GambarURL = produk.GambarProduks[0].NamaGambar
	}

	var perks []string

	for _, perk := range strings.Split(produk.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	produkDetailFormatter.Perks = perks

	user := produk.User
	produkUserFormatter := ProdukUserFormatter{}
	produkUserFormatter.Nama = user.Nama
	produkUserFormatter.GambarProduk = user.Gambar
	produkDetailFormatter.User = produkUserFormatter

	images := []ProdukImageFormatter{}

	for _, image := range produk.GambarProduks {
		produkImageFormatter := ProdukImageFormatter{}
		produkImageFormatter.GambarProduk = image.NamaGambar

		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		produkImageFormatter.IsPrimary = isPrimary

		images = append(images, produkImageFormatter)

	}

	produkDetailFormatter.Images = images

	return produkDetailFormatter
}
