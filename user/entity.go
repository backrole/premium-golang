package user

import "time"

type User struct {
	ID         int
	Nama       string
	Kampus     string
	Email      string
	Password   string
	Gambar     string
	Role       string
	Created_at time.Time
	Updated_at time.Time
}
