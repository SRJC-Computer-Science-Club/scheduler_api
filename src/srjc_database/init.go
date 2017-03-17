package srjc_database

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"fmt"
	"log"
	"github.com/joho/godotenv"
	"os"
)







var database *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil{
		log.Panic("error in loading .env")
	}

	dbname := os.Getenv("DBNAME")
	url := os.Getenv("URL")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	var params string = "" +
		" user= "+ username +
		" password="+password +
		" host="+url +
		" dbname="+dbname +
		" sslmode=disable"

	database , err := gorm.Open("postgres", params)

	if err != nil {
		log.Println("error for connect")
		log.Panic(err)
	}
	fmt.Print("connected to database")
	database.Where("section_id = ?", 8007).Find(&sections{})
}

