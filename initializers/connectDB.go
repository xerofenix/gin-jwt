package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	//getting db url from env
	dsn := os.Getenv("DB_URL")

	var err error
	//connecting to db via gorm
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error! connecting to database ", err)
	}
	fmt.Println("Databse connected successfully")
}
