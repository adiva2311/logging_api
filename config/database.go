package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var DB *gorm.DB

func ConnectDB() *gorm.DB{
	//  konfigurasi PostgreSQL

	dsn := "host=117.53.47.52 user=aditia password=jaringan123 dbname=automobi port=1821 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	fmt.Println("Database connected successfully!")

	return DB
}

