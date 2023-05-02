package mysql

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
}

func SqlInit() (*DB, error) {
	db, err := InitDatabase()
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil

}

func InitDatabase() (*gorm.DB, error) {
	DataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"))

	db, err := gorm.Open(mysql.Open(DataSourceName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatal("Failed to open database!")
	}

	return db, nil
}
