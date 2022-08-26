package bd

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var (
	DB  *gorm.DB
	ERR error
)

func Connect() {
	godotenv.Load(".env")
	DB, ERR = gorm.Open("postgres", os.Getenv("POSTGRES"))

	if ERR != nil {
		panic("failed to connect database")
	}
}
