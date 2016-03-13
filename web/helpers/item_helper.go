package helpers

import (
	"github.com/alexstoick/wow/models"
	//"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetMaterialsForItem(item models.Item, db gorm.DB) []models.ItemMaterial {
	var itemMaterials []models.ItemMaterial
	db.Debug().Where(models.ItemMaterial{ItemID: item.ItemID}).Preload("Material").Find(&itemMaterials)
	//db.Model(&itemMaterials[0]).Preload("Material")
	return itemMaterials
}
