package repo

import (
	"DTS/Chapter-3/final-project-myGram/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "gume98"
	port     = "5432"
	dbname   = "db-my-gram"
	db       *gorm.DB
	err      error
)

func StartDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", host, user, password, port, dbname)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Invalid connect to database")
	}

	log.Println("Success connect to database")

	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
	// db.Debug().Migrator().DropTable(models.Photo{})
}

func GetDB() *gorm.DB {
	return db
}
