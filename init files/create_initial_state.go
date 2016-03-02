package main

import (
	"fmt"
	"github.com/alexstoick/wow/database"
	"github.com/alexstoick/wow/models"
	"github.com/empatica/csvparser"
	"github.com/jinzhu/gorm"
	"github.com/oleiade/reflections"
	"time"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func ImportItemsCSV() {
	var csvParser = parser.CsvParser{
		CsvFile:         "items.csv",
		CsvSeparator:    ',',
		SkipFirstLine:   true, //default:false
		SkipEmptyValues: true, //default:false. It will skip empty values and won't try to parse them
	}
	db := database.ConnectToDb()
	t0 := time.Now()
	parsedCSV, err := csvParser.Parse(models.Item{})
	handleError(err)
	t1 := time.Now()
	fmt.Printf("The CSV call took %v to run.\n", t1.Sub(t0))
	fmt.Printf("%v", len(parsedCSV))

	for i := 15000; i < len(parsedCSV); i++ {
		item := parsedCSV[i].(*models.Item)
		fmt.Printf("Inserting %ld\n", i)
		db.Create(&item)
	}
}

func CreateCraftableItems() {
	type CraftableItem struct {
		ItemID   int
		ItemName string
	}

	db := database.ConnectToDb()
	var csvParser = parser.CsvParser{
		CsvFile:         "craftable.csv",
		CsvSeparator:    ',',
		SkipFirstLine:   true, //default:false
		SkipEmptyValues: true, //default:false. It will skip empty values and won't try to parse them
	}

	t0 := time.Now()
	parsedCSV, err := csvParser.Parse(CraftableItem{})
	handleError(err)
	t1 := time.Now()
	fmt.Printf("The CSV call took %v to run.\n", t1.Sub(t0))

	fmt.Printf("%v", len(parsedCSV))

	for i := 0; i < len(parsedCSV); i++ {
		craft_item := parsedCSV[i].(*CraftableItem)
		fmt.Printf("Inserting %d\n", i)

		item := models.Item{ItemID: craft_item.ItemID, ItemName: craft_item.ItemName}
		fmt.Printf("Real item: %+v", item)
		db.Create(&item)
	}
}

func CreateItemMats() {
	db := database.ConnectToDb()
	var csvParser = parser.CsvParser{
		CsvFile:         "./craftable_mats.csv",
		CsvSeparator:    ',',
		SkipFirstLine:   true, //default:false
		SkipEmptyValues: true, //default:false. It will skip empty values and won't try to parse them
	}

	type ItemAndMaterials struct {
		ItemID         int
		CreatedBySpell int
		Mat1           int
		Mat1Qty        int
		Mat2           int
		Mat2Qty        int
		Mat3           int
		Mat3Qty        int
		Mat4           int
		Mat4Qty        int
		Mat5           int
		Mat5Qty        int
		Mat6           int
		Mat6Qty        int
		Mat7           int
		Mat7Qty        int
		Mat8           int
		Mat8Qty        int
	}

	t0 := time.Now()
	parsedCSV, err := csvParser.Parse(ItemAndMaterials{})
	handleError(err)
	t1 := time.Now()
	fmt.Printf("The CSV call took %v to run.\n", t1.Sub(t0))

	fmt.Printf("CSV lenght %v\n", len(parsedCSV))

	for i := 0; i < len(parsedCSV); i++ {
		item_and_mats := parsedCSV[i].(*ItemAndMaterials)
		fmt.Printf("Inserting %d\n", i)

		for i := 1; i < 9; i++ {
			fieldName1 := fmt.Sprintf("Mat%d", i)
			fieldName2 := fmt.Sprintf("Mat%dQty", i)
			value1, _ := reflections.GetField(item_and_mats, fieldName1)
			value2, _ := reflections.GetField(item_and_mats, fieldName2)
			item_mat := models.ItemMaterial{
				ItemID:     item_and_mats.ItemID,
				MaterialID: value1.(int),
				Quantity:   value2.(int),
			}
			fmt.Printf("Real item: %+v\n", item_mat)
			db.Create(&item_mat)
			if value1 == 0 {
				break
			}
		}
	}
}

func main() {

	db := database.ConnectToDb()
	database.AutoMigrateModels(db)
	db.Close()
	ImportItemsCSV()
	CreateCraftableItems()
	CreateItemMats()
}
