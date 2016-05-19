package main

import (
	//"fmt"
	"github.com/alexstoick/wow/database"
	controllers_v1 "github.com/alexstoick/wow/web/controllers/v1"
	controllers_v2 "github.com/alexstoick/wow/web/controllers/v2"
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
		v1.GET("/items/:id", controllers_v1.GetItem)
		v1.GET("/items/:id/crafts", controllers_v1.GetItemCrafts)
		v1.GET("/items/:id/auctions", controllers_v1.GetItemAuctions)
		v1.GET("/items/:id/price", controllers_v1.GetLatestPrice)
		v1.GET("/items/:id/history_price", controllers_v1.GetAveragePricesByDay)
		v1.GET("/last_update", controllers_v1.LastUpdate)
	}

	v2 := router.Group("v2/")
	{
		v2.GET("/items/:id", controllers_v2.GetItem)
		//v1.GET("/items/:id/price", controllers_v2.GetLatestPrice)
		v2.GET("/items/:id/history_price", controllers_v2.GetPricesByDay)
		v2.GET("/last_update", controllers_v2.LastUpdate)
		v2.GET("/items/:id/auctions/:user", controllers_v2.GetItemAuctionsByPlayer)
	}

	port := ":3000"

	router.Run(port)
}
