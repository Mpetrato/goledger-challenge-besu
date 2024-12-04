package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Mpetrato/goledger-challenge-besu/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	var err error
	var db *gorm.DB
	var sqlDB *sql.DB

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, errors.New("error on load env")
	}

	for i := 0; i < 10; i++ {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			sqlDB, err := db.DB()
			if err != nil {
				fmt.Printf("error on get sqlDB instance: %v\n", err)
				return nil, err
			}

			if err := sqlDB.Ping(); err == nil {
				return db, nil
			}

			fmt.Println("retry -> ", i+1)
		}

		time.Sleep(2 * time.Second)
	}

	db.AutoMigrate(&model.ContractModel{})

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)

	return nil, fmt.Errorf("could not connect to the database after 10 retries: %v", err)
}
