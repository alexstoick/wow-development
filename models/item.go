package models

import (
	"fmt"
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

type ItemSummaryWithoutCrafts struct {
	Item       Item
	BuyPrice   int
	CraftPrice int
}

type ItemSummary struct {
	Item      Item
	BuyPrice  int
	UpdatedAt time.Time
	Crafts    []CraftSummary
}

type CraftSummary struct {
	SpellID    int
	Price      int
	Name       string
	Profession string
}

func (item Item) GetLatestPrice(db gorm.DB) int {
	prices := item.GetAveragePrices(db)

	if len(prices) > 0 {
		return prices[0].Average
	}
	return 0
}

func (item Item) GetAveragePrices(db gorm.DB) []PriceSummary {
	rows, err := db.Debug().Raw("select count(auctions.auction_id), avg(buyout/quantity), imported_at::date, extract(hour from imported_at) from auctions where item_id =? group by 3,4 order by 3,4", item.ItemID).Rows()
	var summary []PriceSummary
	fmt.Printf("DB in GetAvgPrices: %+v\n", db)
	fmt.Printf("rows in GetAvgPrices: %+v\n", rows)
	fmt.Printf("err in GetAvgPrices: %+v\n", err)
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

func (item Item) CheapestCraftPrice(db gorm.DB) int {
	min := 9999999999

	for _, spell := range item.Spells {
		price := spell.GetLatestCraftPrice(db)
		if min > price {
			min = price
		}
	}

	return min
}

func (item Item) CreateSpellsForDisplay(db gorm.DB) []SpellSummary {
	var summary []SpellSummary

	for _, spell := range item.Spells {
		summary = append(summary, spell.CreateSummary(db))
	}
	return summary
}

func (item Item) CreateSummaryWithoutCrafts(db gorm.DB) ItemSummaryWithoutCrafts {
	return ItemSummaryWithoutCrafts{item, item.GetLatestPrice(db), item.CheapestCraftPrice(db)}
}

func (item Item) CreateSummary(db gorm.DB) ItemSummary {
	//buyPrice := item.GetLatestPrice(db)
	var crafts []CraftSummary
	for _, spell := range item.Spells {
		crafts = append(crafts, CraftSummary{
			SpellID:    spell.SpellID,
			Price:      spell.GetLatestCraftPrice(db),
			Profession: spell.Profession,
			Name:       spell.SpellName,
		})
	}
	var updated_at time.Time
	var buyPrice int
	if len(item.Auctions) > 0 {
		updated_at = item.Auctions[0].ImportedAt
		buyPrice = item.Auctions[0].Buyout / item.Auctions[0].Quantity
	} else {
		buyPrice = 999999999
		updated_at = time.Now().AddDate(0, 0, -3)
	}

	item.Auctions = []Auction{}
	summary := ItemSummary{item, buyPrice, updated_at, crafts}
	return summary
}
