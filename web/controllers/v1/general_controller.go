package controllers

import (
	"github.com/alexstoick/wow/models"
	"github.com/gin-gonic/gin"
)

func LastUpdate(c *gin.Context) {
	db := FetchDatabaseFromContext(c)
	var ah_file models.AHFile
	db.Order("id DESC").First(&ah_file)
	c.JSON(200, ah_file)
}
