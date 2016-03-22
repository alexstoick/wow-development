package main

import (
	//"fmt"
	"github.com/alexstoick/wow/database"
	"github.com/alexstoick/wow/web/controllers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db gorm.DB

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func DatabaseMapper(c *gin.Context) {
	c.Set("db", db)

	c.Next()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {

	db = database.ConnectToDb()
	database.AutoMigrateModels(db)

	router := gin.Default()

	router.Use(DatabaseMapper)
	router.Use(CORSMiddleware())

	v1 := router.Group("v1/")
	{
		v1.GET("/items/:id", controllers.GetItem)
		v1.GET("/items/:id/crafts", controllers.GetItemCrafts)
		v1.GET("/items/:id/auctions", controllers.GetItemAuctions)
		v1.GET("/items/:id/price", controllers.GetLatestPrice)
		v1.GET("/items/:id/history_price", controllers.GetAveragePricesByDay)
		v1.GET("/last_update", controllers.LastUpdate)
	}

	port := ":3000"

	router.Run(port)
}
