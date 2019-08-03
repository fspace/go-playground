package models

// ... snip ...

type Todo struct {
	gorm.Model

	Name string
	Done bool

	User   User
	UserID int
}
