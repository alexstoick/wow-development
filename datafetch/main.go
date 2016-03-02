package main

import (
	"encoding/json"
	"fmt"
	"github.com/alexstoick/wow/database"
	"github.com/alexstoick/wow/models"
	"github.com/jinzhu/gorm"
	//"github.com/robfig/cron"
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
	var db_auct models.Auction
	db.Where(models.Auction{AuctionID: auction.AuctionID}).First(&db_auct)
	if db_auct.AuctionID == 0 {
		db.FirstOrCreate(&auction, auction)
	}
}

func DownloadAHFile(url string) []byte {
	t0 := time.Now()
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

	t0 := time.Now()
	// for i := 0; i < len(jsonfile.Auctions); i++ { //i < 5; i++ { //
	for i := 0; i < 5; i++ {
		auction := jsonfile.Auctions[i]
		auction.ImportedFromId = ah_file.ID
		auction.ImportedAt = time.Now()
		SaveAuction(auction, i, db)
	}

	t1 := time.Now()
	fmt.Printf("The call to SAVE AUCTIONS took %v to run.\n", t1.Sub(t0))
}

func GetLatestAHFilelist() (models.AHFile, bool) {

	resp, err1 := http.Get("https://eu.api.battle.net/wow/auction/data/quelthalas?locale=en_GB&apikey=5kuxc3d7rjwk75dvds22egepcwajwtqx")
	if err1 != nil {
		panic(err1)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	type DataDump struct {
		Files []models.AHFile
	}
	var file DataDump
	err := json.Unmarshal(body, &file)
	if err != nil {
		panic(err)
	}
	ahfile := file.Files[0]
	db := database.ConnectToDb()
	var db_file models.AHFile
	db.Where(models.AHFile{LastModified: ahfile.LastModified}).First(&db_file)
	if db_file.ID == 0 {
		db.Create(&ahfile)
		return ahfile, true
	}
	return db_file, false
}

var running = false

func PullData() {
	fmt.Println("starting to fetch time: " + time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))
	if running == true {
		return
	}
	running = true
	ah_file, new := GetLatestAHFilelist()
	if new {
		fmt.Println("processing auctions")
		ProcessAuctions(ah_file)
	}
	fmt.Println("no new ah file")
	running = false
}

func main() {
	fmt.Println("starting datafetch")
	db := database.ConnectToDb()
	database.AutoMigrateModels(db)
	PullData()

	// c := cron.New()
	// c.AddFunc("@every 5m", func() {
	// 	PullData()
	// })
	// c.Start()
	// select {}
	fmt.Println("ending datafetch")
}
