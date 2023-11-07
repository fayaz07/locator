package main

import (
	"fmt"
	"sort"

	"github.com/fayaz07/locator/common/models"
	"github.com/fayaz07/locator/utils/src/json"
)

const (
	filePath                         = "data/data.json"
	outputDir                        = "data/output"
	filePathTemplate                 = outputDir + "/%s.json"
	placeToRecordFilePathTemplate    = outputDir + "/%s_p.json"
	placeToRecordFilePathTemplateCSV = outputDir + "/csv/%s_p.csv"
)

func main() {

	// run operations on telangana data set
	prepare()

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

func getByState(res []models.Record) {

	// let's separate data by countries
	telangana := []models.Record{}
	for _, record := range res {
		if record.AdminName1 == "Telangana" {
			telangana = append(telangana, record)
		}
	}

	sort.SliceStable(telangana,
		func(i, j int) bool {
			return telangana[i].PlaceName < telangana[j].PlaceName
		},
	)

	locations := []models.LocationModel{}
	for _, record := range telangana {
		locations = append(locations, models.MapRecordToLocationModel(record))
	}

	// save to file
	err := json.SaveToFileL(locations, fmt.Sprintf(placeToRecordFilePathTemplate, "ts"))
	if err != nil {
		panic(err)
	}
}
