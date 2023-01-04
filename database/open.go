package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"strconv"
)

func OpenDB() (*gorm.DB, error) {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("OpenPostgres, bad port %s\n", os.Getenv("DB_PORT"))
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_DATABASE"),
			port,
		),
	}), &gorm.Config{
		NamingStrategy: FieldCamelCaseNamer{Default: schema.NamingStrategy{}},
	})

	if err != nil {
		log.Fatalln("OpenPostgres Error: ", err)
	}

	schema.RegisterSerializer("bigInt", BigIntSerializer{})
	return db, db.Error
}
