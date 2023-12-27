package models

type User struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Password    int64  `json:"password"`
	Status      bool   `json:"status"`
}