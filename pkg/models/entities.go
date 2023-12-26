package models

type User struct {
	ID          int64  `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	Password    int64  `gorm:"column:password"`
	Status      bool   `gorm:"column:status"`
}