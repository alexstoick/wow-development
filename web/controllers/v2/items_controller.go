package v2

import (
	"github.com/alexstoick/wow/models"
	"github.com/gin-gonic/gin"
)

func GetPricesByDay(c *gin.Context) {
	var item models.Item

	db := FetchDatabaseFromContext(c)

	item.Load(c.Param("id"), db)
	summary := item.LoadHourlySummary(10, db)

	c.JSON(200, summary)

}

func GetItem(c *gin.Context) {
	var item models.Item
	db := FetchDatabaseFromContext(c)

	item.Load(c.Param("id"), db)
	item.LoadAuctions(1, db)
	c.JSON(200, item.CreateSummary(FetchDatabaseFromContext(c)))
}
