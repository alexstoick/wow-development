package controllers

import (
	//"encoding/json"
	"fmt"
	"github.com/alexstoick/wow/models"
	//"github.com/alexstoick/wow/web/helpers"
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
	fmt.Printf("DB in fetchItem: %+v\n", db)
	err := db.Debug().Preload("Auctions", func(db *gorm.DB) *gorm.DB {
		return db.Order("(auctions.buyout/auctions.quantity), auctions.imported_at DESC")
	}).Preload("Spells").Preload("Spells.ItemMaterials").Preload("Spells.ItemMaterials.Material").Find(&item, c.Param("id")).Error

	if err != nil {
		panic(err)
	}
	return item
}
func GetItem(c *gin.Context) {
	item := FetchItemFromContext(c)
	c.JSON(200, item.CreateSummary(FetchDatabaseFromContext(c)))
}
func GetItemCrafts(c *gin.Context) {
	item := FetchItemFromContext(c)
	c.JSON(200, item.CreateSpellsForDisplay(FetchDatabaseFromContext(c)))
}

func GetLatestPrice(c *gin.Context) {
	item := FetchItemFromContext(c)
	c.JSON(200, map[string]int{"price": item.GetLatestPrice(FetchDatabaseFromContext(c))})
}
func GetAveragePricesByDay(c *gin.Context) {
	item := FetchItemFromContext(c)
	var prices []models.PriceSummary
	prices = item.GetAveragePrices(FetchDatabaseFromContext(c))

	c.JSON(200, prices)
}

func GetItemAuctions(c *gin.Context) {
	item := FetchItemFromContext(c)
	aucts := item.Auctions
	var presenterAucts []models.PublicAuction
	for _, auct := range aucts {
		presenterAucts = append(presenterAucts, auct.GetPresenter())
	}
	c.JSON(200, presenterAucts)
}
