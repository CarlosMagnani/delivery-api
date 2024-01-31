package db

import (
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
type GormAdapter struct {
	db *gorm.DB
}

var (
	db *gorm.DB
	once sync.Once
)


func InitDB() GormAdapter {
	sc := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "process", "processdb", "delivery")

	conn, err := gorm.Open(postgres.Open(sc), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	return GormAdapter{conn}
}


func (adapter GormAdapter) GetDB() *gorm.DB {
	return adapter.db

}