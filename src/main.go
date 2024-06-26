package main

// import (
// 	prepareModule "github.com/fayaz07/locator/core/src/prepare"
// )

const (
	filePath  = "data/data.json"
	outputDir = "data/output"

	filePathTemplate                 = outputDir + "/%s.json"
	placeToRecordFilePathTemplate    = outputDir + "/%s_p.json"
	placeToRecordFilePathTemplateCSV = outputDir + "/csv/%s_p.csv"
)

func main() {
	//prepareModule.PrepareDatasetByCountry(filePath, outputDir)

	// prepByCountry := prepareModule.PrepareData{
	// 	DataConfig: prepareModule.DataConfig{
	// 		InputFilePath:  filePath,
	// 		OutputPath:     outputDir,
	// 		MinQueryLength: 3,
	// 		MaxQueryLength: 7,
	// 	},
	// 	Mode: prepareModule.ByCountry,
	// }

	// prepByCountry.PrepareDatasetByCountry()

	// var records []models.LocationModel
	// gocsv.UnmarshalFile(file, &records)

	// log.Println("Records count: ", len(records))
	// log.Println("Gutpa: ", records[60323].PlaceName)
	// log.Println("Gutpa: ", records[60324].PlaceName)
	// log.Println("Gutpa: ", records[60325].PlaceName)

	// search.InitSearch(outputDir)

	// for {
	// 	// read input from user
	// 	fmt.Print("Enter place name to search: ")
	// 	var placeName string
	// 	fmt.Scanln(&placeName)

	// 	results, err := search.SearchByCountry("IN", strings.ToUpper(placeName))
	// 	if err != nil {
	// 		log.Println("Error: ", err)
	// 		continue
	// 	}

	// 	// print results
	// 	log.Println("Results, count:", len(results))
	// 	for i := 0; i < len(results); i++ {
	// 		log.Println(results[i].GetRow())
	// 	}
	// }

	// run operations on telangana data set
	// prepare()

	// prepare data
	// res, err := json.ParseWithJsonIterator(filePath)
	// if err != nil {
	// 	panic(err)
	// }

	// getByState(res)

	// let's separate data by countries
	// countries := make(map[string][]models.Record)
	// for _, record := range res {
	// 	countries[record.CountryCode] = append(countries[record.CountryCode], record)
	// }

	// save data to files
	// for countryCode, records := range countries {
	// 	fmt.Println("Saving data for country: ", countryCode)
	// 	err := json.SaveToFile(records, fmt.Sprintf(filePathTemplate, countryCode))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// go through all places and sort them in alphabetical order

	// for countryCode, records := range countries {
	// 	fmt.Println("Processing data for country: ", countryCode)
	// 	sort.SliceStable(records,
	// 		func(i, j int) bool {
	// 			return records[i].PlaceName < records[j].PlaceName
	// 		},
	// 	)
	// 	locations := []models.LocationModel{}
	// 	for _, record := range records {
	// 		locations = append(locations, models.MapRecordToLocationModel(record))
	// 	}
	// 	// save to file
	// 	// err = json.SaveToFileL(locations, fmt.Sprintf(placeToRecordFilePathTemplate, countryCode))
	// 	// if err != nil {
	// 	// 	panic(err)
	// 	// }
	// 	err = csv.SaveToFile(locations, fmt.Sprintf(placeToRecordFilePathTemplateCSV, countryCode))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
}

// func getByState(res []models.RecordModel) {

// 	// let's separate data by countries
// 	telangana := []models.RecordModel{}
// 	for _, record := range res {
// 		if record.AdminName1 == "Telangana" {
// 			telangana = append(telangana, record)
// 		}
// 	}

// 	sort.SliceStable(telangana,
// 		func(i, j int) bool {
// 			return telangana[i].PlaceName < telangana[j].PlaceName
// 		},
// 	)

// 	locations := []models.LocationModel{}
// 	for _, record := range telangana {
// 		locations = append(locations, models.MapRecordToLocationModel(record))
// 	}

// 	// save to file
// 	err := json.SaveToFileL(locations, fmt.Sprintf(placeToRecordFilePathTemplate, "ts"))
// 	if err != nil {
// 		panic(err)
// 	}
// }
