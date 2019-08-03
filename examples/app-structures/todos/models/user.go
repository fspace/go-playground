package models

type User struct {
	gorm.Model

	Email          string
	HashedPassword []byte
}
