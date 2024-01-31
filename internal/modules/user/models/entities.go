package models

type User struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Password    string `json:"password"`
	Status      bool   `json:"status"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}