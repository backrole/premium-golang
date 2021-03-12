package user

type UserFormatter struct {
	ID     int    `json:"id"`
	Nama   string `json:"nama"`
	Kampus string `json:"kampus"`
	Email  string `json:"email"`
	Token  string `json:"token"`
	Gambar string `json:"gambar"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:     user.ID,
		Nama:   user.Nama,
		Kampus: user.Kampus,
		Email:  user.Email,
		Token:  token,
		Gambar: user.Gambar,
	}

	return formatter
}
