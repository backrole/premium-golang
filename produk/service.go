package produk

type Service interface {
	GetProduks(userID int) ([]Produk, error)
	GetProdukByID(input GetProdukDetailInput) (Produk, error)
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
