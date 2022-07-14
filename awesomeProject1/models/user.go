package models

type User struct {
	ID    string `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

const ActiveUser string = "active_user"
