package models

import (
	"time"
)

type Auction struct {
	AuctionID      int `json:"auc" sql:"type:bigint;unique" gorm:"primary_key"`
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

type PublicAuction struct {
	Auction    `json:"-"`
	AuctionID  int
	Item_id    int `json:"ItemID"`
	Owner      string
	OwnerRealm string
	Bid        int
	Buyout     int
	Quantity   int
	ImportedAt time.Time
}

func (auction Auction) GetPresenter() PublicAuction {
	return PublicAuction{
		AuctionID:  auction.AuctionID,
		Item_id:    auction.Item_id,
		Owner:      auction.Owner,
		OwnerRealm: auction.OwnerRealm,
		Bid:        auction.Bid,
		Buyout:     auction.Buyout,
		Quantity:   auction.Quantity,
		ImportedAt: auction.ImportedAt,
	}
}
