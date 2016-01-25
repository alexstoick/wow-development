package main

//
//func main() {
//	connectToDb()
//	autoMigrateModels()
//
//	t0 := time.Now()
//	resp, _ := http.Get("https://eu.api.battle.net/wow/auction/data/quelthalas?locale=en_GB&apikey=5kuxc3d7rjwk75dvds22egepcwajwtqx")
//	defer resp.Body.Close()
//	body, _ := ioutil.ReadAll(resp.Body)
//	t1 := time.Now()
//	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
//
//	type DataDump struct {
//		Files []models.AHFile
//	}
//	var file DataDump
//	ffjson.Unmarshal(body, &file)
//	fmt.Printf("%+v\n", file)
//	fmt.Printf("%+v\n", file.Files[0].URL)
//	ahfile := file.Files[0]
//	db.Create(&ahfile)
//}
