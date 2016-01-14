package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type JSONFile struct {
	Realms   []Realm
	Auctions []Auction
}
type Realm struct {
	Name string
	Slug string
}
type Auction struct {
	ID           int `json:"auc"`
	Item_id      int `json:"item"`
	Owner        string
	OwnerRealm   string
	Bid          int
	Buyout       int
	Quantity     int
	TimeLeft     string
	Rand         int
	Seed         int
	Context      int
	BonusList    []Bonus `json:"bonusLists"`
	Modifiers    []Modifier
	PetSpeciesId int
	PetBreedId   int
	PetLevel     int
	PetQualityId int
}
type Modifier struct {
	Type  int
	Value int
}
type Bonus struct {
	Id int `json:"bonusListId"`
}

func main() {
	fmt.Printf("lol\n")
	file, _ := ioutil.ReadFile("./auctions.json")

	var jsonfile JSONFile
	t0 := time.Now()
	json.Unmarshal(file, &jsonfile)
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
	//fmt.Printf("res: %+v\n", jsonfile)

}
