package models

import "time"

type Realm struct {
	Name string
	Slug string
}
type Auction struct {
	//ID             int `gorm:"primary_key"`
	AuctionID      int `json:"auc" sql:"type:bigint;unique"`
	Item_id        int `json:"item" sql:"index:auction_item_id_index;type:bigint"`
	Owner          string
	OwnerRealm     string
	Bid            int `sql:"type:bigint"`
	Buyout         int `sql:"type:bigint"`
	Quantity       int
	TimeLeft       string
	Rand           int `sql:"type:bigint"`
	Seed           int `sql:"type:bigint"`
	Context        int
	BonusList      []Bonus `json:"bonusLists"`
	Modifiers      []Modifier
	PetSpeciesId   int
	PetBreedId     int
	PetLevel       int
	PetQualityId   int
	ImportedFrom   AHFile
	ImportedFromId int       `json:"-" sql:"index:auction_imported_from_index"`
	ImportedAt     time.Time `sql:"index:auction_imported_at_index"`
}
type Modifier struct {
	Type  int
	Value int
}
type Bonus struct {
	Id int `json:"bonusListId"`
}

type ItemMaterial struct {
	ID         int `gorm:"primary_key"`
	ItemID     int `sql:"unique_index:item_and_material"`
	MaterialID int `sql:"unique_index:item_and_material"`
	Quantity   int
}

type Item struct {
	ItemID         int    `gorm:"primary_key"`
	ItemName       string `sql:"index:items_item_name_index"`
	InBlizzardAPI  int
	InWowheadAPI   int
	IsAuctionable  int
	GlobalMedianEU int
	Alchemy        int
	Archaeology    int
	Blacksmithing  int
	Cooking        int
	Disenchanting  int
	Enchanting     int
	Engineering    int
	Firstaid       int
	Herbalism      int
	Inscription    int
	Jewelcrafting  int
	Leatherworking int
	Milling        int
	Mining         int
	Prospecting    int
	Skinning       int
	Tailoring      int
	StackSize      int
	BuyPrice       int `sql:"type:bigint"`
	SellPrice      int `sql:"type:bigint"`
	ItemClass      string
	ItemSubClass   string
	ItemType       string
	InventoryType  string
	Equippable     int
	Source         string
	SourceId       int
	SourceType     string
	SourceDesc     string
	DeprecatedMsg  string
}
type AHFile struct {
	ID           int    `gorm:"primary_key"`
	URL          string `sql:"not_null"`
	LastModified int    `sql:"type:bigint;unique"`
	CreatedAt    time.Time
}
