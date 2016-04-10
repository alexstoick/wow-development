package v2

import (
	//"fmt"
	"github.com/alexstoick/wow/models"
	"github.com/alexstoick/wow/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func FetchDatabaseFromContext(c *gin.Context) gorm.DB {
	fake_db, _ := c.Get("db")
	db := fake_db.(gorm.DB)
	return db
}

func FetchItemFromContext(c *gin.Context) models.Item {
	var item models.Item
	db := FetchDatabaseFromContext(c)
	fetcher := services.ItemFetcher{db}
	item, err := fetcher.LoadItem(c.Param("id"))

	if err != nil {
		panic(err)
	}
	return item
}
