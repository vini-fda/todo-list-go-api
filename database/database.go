package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func init() {
	godotenv.Load()

	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v",
		os.Getenv("PG_HOST"), os.Getenv("PG_USER"), os.Getenv("PG_PASS"),
		os.Getenv("PG_DBNM"), os.Getenv("PG_PORT"),
	)

	fmt.Println(dsn)

	var err error
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connection successfully opened!!!!")
}
