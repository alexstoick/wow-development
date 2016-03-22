package controllers

import (
	"github.com/alexstoick/wow/models"
	"github.com/gin-gonic/gin"
)

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
