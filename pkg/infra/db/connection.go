package db

import (
	"delivery_api/cmd/configs"
	"delivery_api/pkg/models"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	once sync.Once
)

func initDB() {
	conf := configs.GetDatabase()
	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disabled", conf.Host, conf.Port, conf.User, conf.Password, conf.Database)
	conn, err := gorm.Open(postgres.Open(sc), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = conn.AutoMigrate(&models.User{})
	if err != nil{
		panic(err)
	}

	db = conn
}


func OpenConnection() *gorm.DB {
	once.Do(initDB)
	return db
}