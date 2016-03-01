package database

import (
	"fmt"
	"github.com/alexstoick/wow/models"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"time"
)

func AutoMigrateModels(db gorm.DB) {
	t0 := time.Now()
	db.AutoMigrate(&models.Item{})
	db.AutoMigrate(&models.ItemMaterial{})
	db.AutoMigrate(&models.AHFile{})
	db.AutoMigrate(&models.Auction{})
	t1 := time.Now()
	fmt.Printf("The DB MIGRATE call took %v to run.\n", t1.Sub(t0))
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
func ConnectToDb() (db gorm.DB) {
	const (
		DB_USER     = ""
		DB_PASSWORD = ""
		DB_NAME     = "wow_development"
	)
	dbinfo := fmt.Sprintf("host=wow_db_1 user=postgres dbname=%s sslmode=disable", DB_NAME)
	fmt.Println(dbinfo)
	pq.ParseURL(dbinfo)
	var err error

	db, err = gorm.Open("postgres", dbinfo)
	handleError(err)
	return db
}
