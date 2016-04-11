package v2

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func FetchDatabaseFromContext(c *gin.Context) gorm.DB {
	fake_db, _ := c.Get("db")
	db := fake_db.(gorm.DB)
	return db
}
