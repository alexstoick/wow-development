package models

import (
	"github.com/jinzhu/gorm"
)

type Spell struct {
	SpellID       int `gorm:"primary_key" sql:"unique_index:spell_and_item" json:"-"`
	SpellName     string
	Profession    string
	ItemID        int            `sql:"unique_index:spell_and_item" json:"-"`
	ItemMaterials []ItemMaterial `gorm:"foreignkey:SpellID"`
}

type SpellSummary struct {
	SpellName  string
	Profession string
	Items      []ItemSummaryWithoutCrafts
}

func (spell Spell) CreateSummary(db gorm.DB) SpellSummary {

	var sum []ItemSummaryWithoutCrafts
	db.Debug().Preload("ItemMaterials").Preload("ItemMaterials.Material").Preload("ItemMaterials.Material.Spells").Preload("ItemMaterials.Material.Spells.ItemMaterials").Preload("ItemMaterials.Material.Spells.ItemMaterials.Material").Find(&spell)

	for _, mat := range spell.ItemMaterials {
		sum = append(sum, mat.Material.CreateSummaryWithoutCrafts(db))
	}
	return SpellSummary{spell.SpellName, spell.Profession, sum}
}

func (spell Spell) GetLatestCraftPrice(db gorm.DB) int {
	sum := 0
	for _, itemMat := range spell.ItemMaterials {
		sum = sum + itemMat.Material.GetLatestPrice(db)
	}
	return sum
}
