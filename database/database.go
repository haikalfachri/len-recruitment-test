package database

import (
	"len-test/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	DB_USERNAME string `mapstructure:"POSTGRES_USER"`
	DB_PASSWORD string `mapstructure:"POSTGRES_PASSWORD"`
	DB_NAME     string `mapstructure:"POSTGRES_DB"`
	DB_HOST     string `mapstructure:"POSTGRES_HOST"`
	DB_PORT     string `mapstructure:"POSTGRES_PORT"`
}

func (c *Config) ConnectDB() {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		c.DB_HOST,
		c.DB_USERNAME,
		c.DB_PASSWORD,
		c.DB_NAME,
		c.DB_PORT,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Printf("successfully connected to database\n")
}

func MigrateDB() {

	err := DB.AutoMigrate(&models.Student{}, &models.Book{}, &models.Borrowing{})

	if err != nil {
		log.Fatalf("failed to perform database migration: %s\n", err)
	}

	log.Printf("successfully database migration\n")
}


