package main

import (
	"github.com/LaGregance/gorm-save-bug/database"
	"log"
)

func main() {
	db, err := database.OpenDB()
	if err != nil {
		log.Fatalln("Can't open the DB: ", err)
	}

	err = db.AutoMigrate(&database.User{})
	if err != nil {
		log.Fatalln("AutoMigrate: ", err)
	}
}
