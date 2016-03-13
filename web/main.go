package main

import (
	//"fmt"
	"github.com/alexstoick/wow/database"
	"github.com/alexstoick/wow/web/controllers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
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

	godotenv.Load()

	router := gin.Default()

	router.Use(DatabaseMapper)
	router.Use(CORSMiddleware())

	v1 := router.Group("v1/")
	{
		v1.GET("/items/:id", controllers.GetItem)
		v1.GET("/items/:id/materials", controllers.GetItemMaterials)
		v1.GET("/items/:id/auctions", controllers.GetItemAuctions)
		v1.GET("/items/:id/price", controllers.GetLatestPrice)
		v1.GET("/items/:id/history_price", controllers.GetAveragePricesByDay)
		//v1.POST("/users", controllers.CreateUser)
		//v1.POST("/login", controllers.AuthUser)
		//v1.GET("/verify_token", controllers.VerifyToken)
		//v1.POST("/renew_token", controllers.RenewToken)

		//authentication := v1.Use(controllers.ValidateAuthentication)

		//authentication.POST("/users/me/payments", controllers.CreatePayment)

		//authentication.GET("/users/me/payments", controllers.GetUserPayments)
		//authentication.GET("/users/me/payments/:payment_id", controllers.GetPaymentBeneficiaries)
	}

	port := ":3000"

	router.Run(port)
}
