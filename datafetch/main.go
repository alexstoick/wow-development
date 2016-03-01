package main

import (
	"encoding/json"
	"fmt"
	"github.com/alexstoick/wow/database"
	"github.com/alexstoick/wow/models"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"time"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func SaveAuction(auction models.Auction, i int, db gorm.DB) {
	db.Create(&auction)
}

func DownloadAHFile(url string) []byte {
	t0 := time.Now()
	fmt.Println(url)
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		panic(err1)
	}
	t1 := time.Now()
	fmt.Printf("The call to DOWNLOAD AH FILE took %v to run.\n", t1.Sub(t0))
	return body
}

func ProcessAuctions(ah_file models.AHFile) {

	body := DownloadAHFile(ah_file.URL)
	type JSONFile struct {
		Realms   []models.Realm
		Auctions []models.Auction
	}

	jsonfile := JSONFile{}
	err := json.Unmarshal(body, &jsonfile)
	if err != nil {
		panic(err)
	}

	db := database.ConnectToDb()
	fmt.Printf("AUCTIONS LENGTH %d\n", len(jsonfile.Auctions))

	t0 := time.Now()
	for i := 0; i < len(jsonfile.Auctions); i++ { //i < 5; i++ { //
		auction := jsonfile.Auctions[i]
		auction.ImportedFromId = ah_file.ID
		auction.ImportedAt = time.Now()
		t2 := time.Now()
		SaveAuction(auction, i, db)
		t3 := time.Now()
		fmt.Printf("The call to INSERT (%d) AUCTIOS took %v to run.\n", i, t3.Sub(t2))
	}

	t1 := time.Now()
	fmt.Printf("The call to SAVE AUCTIONS took %v to run.\n", t1.Sub(t0))
}

func GetLatestAHFilelist() models.AHFile {

	t0 := time.Now()
	resp, err1 := http.Get("https://eu.api.battle.net/wow/auction/data/quelthalas?locale=en_GB&apikey=5kuxc3d7rjwk75dvds22egepcwajwtqx")
	if err1 != nil {
		panic(err1)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))

	type DataDump struct {
		Files []models.AHFile
	}
	var file DataDump
	err := json.Unmarshal(body, &file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", file)
	fmt.Printf("%+v\n", file.Files[0].URL)
	ahfile := file.Files[0]
	db := database.ConnectToDb()
	db.Create(&ahfile)
	return ahfile
}

var running = false

func PullData() {
	fmt.Println("starting pull data")
	if running == true {
		return
	}
	running = true
	db := database.ConnectToDb()
	database.AutoMigrateModels(db)

	ah_file := GetLatestAHFilelist()
	ProcessAuctions(ah_file)
	running = false
}

func main() {
	fmt.Println("starting datafetch")
	db := database.ConnectToDb()
	database.AutoMigrateModels(db)

	ah_file := GetLatestAHFilelist()
	fmt.Println(ah_file)
	for {
	}
	fmt.Println("ending datafetch")
}
