package bd

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"

)

var (
	DB  *gorm.DB
	ERR error
)

func Connect() {
	godotenv.Load(".env")
	DB, ERR = gorm.Open(postgres.Open(os.Getenv("POSTGRES")), &gnorm.Config{})

	if ERR != nil {
		panic("failed to connect database")
	}
}
