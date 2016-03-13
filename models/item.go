package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Item struct {
	ItemID         int       `gorm:"primary_key"`
	ItemName       string    `sql:"index:items_item_name_index"`
	InBlizzardAPI  int       `json:"-"`
	InWowheadAPI   int       `json:"-"`
	IsAuctionable  int       `json:"-"`
	GlobalMedianEU int       `json:"-"`
	Alchemy        int       `json:"-"`
	Archaeology    int       `json:"-"`
	Blacksmithing  int       `json:"-"`
	Cooking        int       `json:"-"`
	Disenchanting  int       `json:"-"`
	Enchanting     int       `json:"-"`
	Engineering    int       `json:"-"`
	Firstaid       int       `json:"-"`
	Herbalism      int       `json:"-"`
	Inscription    int       `json:"-"`
	Jewelcrafting  int       `json:"-"`
	Leatherworking int       `json:"-"`
	Milling        int       `json:"-"`
	Mining         int       `json:"-"`
	Prospecting    int       `json:"-"`
	Skinning       int       `json:"-"`
	Tailoring      int       `json:"-"`
	StackSize      int       `json:"-"`
	BuyPrice       int       `json:"-" sql:"type:bigint"`
	SellPrice      int       `json:"-" sql:"type:bigint"`
	ItemClass      string    `json:"-"`
	ItemSubClass   string    `json:"-"`
	ItemType       string    `json:"-"`
	InventoryType  string    `json:"-"`
	Equippable     int       `json:"-"`
	Source         string    `json:"-"`
	SourceId       int       `json:"-"`
	SourceType     string    `json:"-"`
	SourceDesc     string    `json:"-"`
	DeprecatedMsg  string    `json:"-"`
	Spells         []Spell   `json:"-" gorm:"foreignkey:ItemID"`
	Auctions       []Auction `gorm:"foreignkey:Item_id"`
}

type PriceSummary struct {
	Average int
	Count   int
	Date    time.Time
}

type ItemSummary struct {
	Item       Item
	BuyPrice   int
	CraftPrice int
}

func (item Item) GetLatestPrice(db gorm.DB) int {
	prices := item.GetAveragePrices(db)

	return prices[len(prices)-1].Average
}
func (item Item) GetAveragePrices(db gorm.DB) []PriceSummary {
	rows, _ := db.Raw("select count(auctions.id), avg(buyout), imported_at::date, extract(hour from imported_at) from auctions where item_id =? group by 3,4 order by 3,4", item.ItemID).Rows()
	var summary []PriceSummary
	for rows.Next() {
		var average float64
		var count int
		var date time.Time
		var hour int
		rows.Scan(&count, &average, &date, &hour)
		date = date.Add(time.Duration(int(time.Hour) * hour))
		prcSum := PriceSummary{Average: int(average), Date: date, Count: count}
		summary = append(summary, prcSum)
	}
	return summary
}

func (item Item) GetLatestCraftPrice(db gorm.DB) int {
	var materials []Item
	spell := item.Spells[0]
	for _, itemMat := range spell.ItemMaterials {
		materials = append(materials, itemMat.Material)
	}
	sum := 0
	for _, item := range materials {
		sum = sum + item.GetLatestPrice(db)
	}
	return sum
}

func (item Item) CreateItemSummary(db gorm.DB) ItemSummary {
	buyPrice := item.GetLatestPrice(db)
	craftPrice := item.GetLatestCraftPrice(db)
	summary := ItemSummary{item, buyPrice, craftPrice}
	return summary
}
