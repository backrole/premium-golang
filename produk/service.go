package produk

type Service interface {
	FindProduks(userID int) ([]Produk, error)
}

type service struct {
	repo Repo
}

func NewService(repo Repo) *service {
	return &service{repo}
}

func (s *service) FindProduks(userID int) ([]Produk, error) {
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
