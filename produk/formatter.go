package produk

type ProdukFormatter struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	NamaProduk string `json:"nama_produk"`
	Judul      string `json:"judul"`
	GambarURL  string `json:"gambar_url"`
	Harga      int    `json:"harga"`
}

func FormatProduk(produk Produk) ProdukFormatter {
	produkFormatter := ProdukFormatter{}
	produkFormatter.ID = produk.ID
	produkFormatter.UserID = produk.UserID
	produkFormatter.NamaProduk = produk.NamaProduk
	produkFormatter.Judul = produk.Judul
	produkFormatter.Harga = produk.Harga
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
