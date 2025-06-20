package models

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"  validate:"required"`
	Email    string `json:"email"     validate:"required,email"`
	Premium  bool   `json:"premium"`
}
