package services

import (
	"fmt"
	"github.com/alexstoick/wow/models"
	"github.com/jinzhu/gorm"
	"time"
)

type ItemFetcher struct {
	Db gorm.DB
}

func (fetcher ItemFetcher) LoadItem(id string) (item models.Item, err Error) {
	db := fetcher.Db

	err := db.Find(&item, id).Error

	return item, err
}

func (fetcher ItemFetcher) LoadAuctions() (item models.Item, err Error) {
	db := fetcher.Db
	err = db.Preload("Auctions", func(db *gorm.DB) *gorm.DB {
		return db.Where("present = ?", true).Where("buyout > 0").Order("(auctions.buyout/auctions.quantity), auctions.imported_at DESC")
	}).Find(&item).Error

	return item, err
}

func (fetcher ItemFetcher) LoadSpells() (item models.Item, err Error) {
	db := fetcher.Db
	err = db.Preload("Spells").Preload("Spells.ItemMaterials").Preload("Spells.ItemMaterials.Material").Find(&item).Error

	return item, err
}
