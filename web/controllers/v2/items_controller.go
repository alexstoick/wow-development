package v2

import (
	//"github.com/alexstoick/wow/models"
	"github.com/gin-gonic/gin"
)

func GetItem(c *gin.Context) {
	item := FetchItemFromContext(c)
	c.JSON(200, item.CreateSummary(FetchDatabaseFromContext(c)))
}
