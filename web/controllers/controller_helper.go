package controllers

import (
	//"fmt"
	"github.com/alexstoick/wow/models"
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
	err := db.Debug().Preload("Auctions", func(db *gorm.DB) *gorm.DB {
		return db.Where("present = ?", true).Where("buyout > 0").Order("(auctions.buyout/auctions.quantity), auctions.imported_at DESC")
	}).Preload("Spells").Preload("Spells.ItemMaterials").Preload("Spells.ItemMaterials.Material").Find(&item, c.Param("id")).Error

	if err != nil {
		panic(err)
	}
	return item
}
