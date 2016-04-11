package v2

import (
	//"github.com/alexstoick/wow/models"
	"github.com/gin-gonic/gin"
)

func GetItem(c *gin.Context) {
	var item Item
	db := FetchDatabaseFromContext(c)
	item.Load(c.Param("id"), db).LoadAuctions(db)
	item := FetchItemFromContext(c)
	c.JSON(200, item.CreateSummary(FetchDatabaseFromContext(c)))
}
