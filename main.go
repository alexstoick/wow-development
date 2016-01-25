package main

import (
	"fmt"
	"github.com/alexstoick/wow/database"
	"github.com/alexstoick/wow/models"
	"github.com/jinzhu/gorm"
	"github.com/pquerna/ffjson/ffjson"
	"io/ioutil"
	"net/http"
	"time"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func SaveAuction(auction models.Auction) {
	db := database.ConnectToDb()
	db.Create(&auction)
}

func ProcessAuctions(ah_file models.AHFile) {

	t0 := time.Now()
	resp, _ := http.Get(ah_file.URL)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t1 := time.Now()
	fmt.Printf("The call to DOWNLOAD AH FILE took %v to run.\n", t1.Sub(t0))
	fmt.Printf("ah_file ID %d\n", ah_file.ID)

	type JSONFile struct {
		Realms   []models.Realm
		Auctions []models.Auction
	}

	var jsonfile JSONFile
	t0 = time.Now()
	ffjson.Unmarshal(body, &jsonfile)

	//concurrency := 10
	//sem := make(chan bool, concurrency)

	for i := 0; i < 2; i++ { //i < len(jsonfile.Auctions); i++ { //
		auction := jsonfile.Auctions[i]
		auction.ImportedFromId = ah_file.ID
		SaveAuction(auction)
	}
	t1 = time.Now()
	fmt.Printf("The call to SAVE AUCTIONS took %v to run.\n", t1.Sub(t0))
}

var db gorm.DB

func GetLatestAHFilelist() models.AHFile {

	t0 := time.Now()
	resp, _ := http.Get("https://eu.api.battle.net/wow/auction/data/quelthalas?locale=en_GB&apikey=5kuxc3d7rjwk75dvds22egepcwajwtqx")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))

	type DataDump struct {
		Files []models.AHFile
	}
	var file DataDump
	ffjson.Unmarshal(body, &file)
	fmt.Printf("%+v\n", file)
	fmt.Printf("%+v\n", file.Files[0].URL)
	ahfile := file.Files[0]
	db.Create(&ahfile)
	return ahfile
}

func main() {

	db := database.ConnectToDb()
	database.AutoMigrateModels(db)

	//ah_file := GetLatestAHFilelist()
	var ah_file models.AHFile
	db.Last(&ah_file)
	ProcessAuctions(ah_file)
}
