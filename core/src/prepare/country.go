package prepare

import (
	"fmt"
	"sort"
	"strings"

	"github.com/fayaz07/locator/core/src/models"
	"github.com/fayaz07/locator/core/src/utils/csv"
	"github.com/fayaz07/locator/core/src/utils/json"

	stringUtils "github.com/fayaz07/locator/core/src/utils/string"
)

// Prepare dataset by countries
func PrepareDatasetByCountry(filePath string, outputPath string) {
	fmt.Println("Reading source dataset file: ", filePath)
	records := readRecords(filePath)

	// separate data by countries
	fmt.Println("Mapping records by country")
	countries := mapRecordsByCountry(records)
	fmt.Println("Total countries: ", len(countries))

	// sort locations of each country by name
	for countryCode, records := range countries {
		fmt.Println("Processing data for country: ", countryCode)

		processRecordsByCountry(records, countryCode, outputPath)
	}
}

func readRecords(filePath string) []models.RecordModel {
	// read file and parse to RecordModel
	// RecordModel is the native json structure which the json file is expected to have
	records, err := json.ParseWithJsonIterator(filePath)
	if err != nil {
		panic(fmt.Errorf("unable to read source dataset file: %s, error: %+v", filePath, err))
	}
	return records
}

func mapRecordsByCountry(records []models.RecordModel) map[string][]models.LocationModel {
	countries := make(map[string][]models.LocationModel)

	for _, record := range records {
		// map to LocationModel
		location := models.MapRecordToLocationModel(record)

		// sanitize place name
		location.PlaceName = strings.ToUpper(stringUtils.SanitizePlaceName(location))

		// save location by country
		countries[record.CountryCode] = append(countries[record.CountryCode], location)
	}
	return countries
}

// process records by country
func processRecordsByCountry(
	records []models.LocationModel,
	countryCode string,
	outputPath string,
) {
	// sort locations by name
	sortLocations(records)

	// clear ascii index slice
	asciiIndexSlice := []models.AsciiIndexModel{}

	// generate ascii index for each location
	fmt.Println("Generating searchable index for each location")
	asciiIndexSlice = generateAsciiIndexSlice(records, asciiIndexSlice)

	// sort ascii index slice by code
	sortAsciiIndexSlice(asciiIndexSlice)

	// save ascii indexes list to file
	fmt.Printf("Saving searchable index list to file: %s/country/%s_ascii.csv\n", outputPath, countryCode)
	saveAsciiIndexSliceToFile(asciiIndexSlice, countryCode, outputPath)

	// save locations list to file
	fmt.Printf("Saving location list to file: %s/country/%s_src.csv\n", outputPath, countryCode)
	saveLocationsToFile(records, countryCode, outputPath)

	fmt.Println("Processing data for country: ", countryCode, " completed\n")
}

func saveLocationsToFile(locations []models.LocationModel, countryCode string, outputPath string) {
	err := csv.SaveToFile(locations, fmt.Sprintf("%s/country/%s_src.csv", outputPath, countryCode))
	if err != nil {
		panic(fmt.Errorf("unable to save location data to file: %s, for country: %s, error: %+v", outputPath, countryCode, err))
	}
}

func saveAsciiIndexSliceToFile(asciiIndexSlice []models.AsciiIndexModel, countryCode string, outputPath string) {
	err := csv.SaveToFile(asciiIndexSlice, fmt.Sprintf("%s/country/%s_ascii.csv", outputPath, countryCode))
	if err != nil {
		panic(fmt.Errorf("unable to save location data to file: %s, for country: %s, error: %+v", outputPath, countryCode, err))
	}
}

func sortAsciiIndexSlice(asciiIndexSlice []models.AsciiIndexModel) {
	sort.SliceStable(
		asciiIndexSlice,
		func(i, j int) bool {
			return asciiIndexSlice[i].Code < asciiIndexSlice[j].Code
		},
	)
}

func generateAsciiIndexSlice(
	records []models.LocationModel,
	asciiIndexSlice []models.AsciiIndexModel,
) []models.AsciiIndexModel {
	for i, record := range records {
		subStrings := stringUtils.GenerateSubstrings(record.PlaceName, 3, 7, 3)

		asciiIndexSlice = append(
			asciiIndexSlice,
			generateAsciiIndexSliceForPlace(subStrings, asciiIndexSlice, i)...,
		)
	}
	return asciiIndexSlice
}

func generateAsciiIndexSliceForPlace(
	subStrings []string,
	asciiIndexSlice []models.AsciiIndexModel,
	index int,
) []models.AsciiIndexModel {
	for _, subString := range subStrings {
		placeName := strings.TrimSpace(subString)
		asciiIndexSlice = append(asciiIndexSlice, models.AsciiIndexModel{
			Code:   stringUtils.ConvertToAscii(placeName),
			Index:  index,
			Length: len(placeName),
		})
	}
	return asciiIndexSlice
}

// sort locations by name
func sortLocations(locations []models.LocationModel) {
	sort.SliceStable(
		locations,
		func(i, j int) bool {
			return locations[i].PlaceName < locations[j].PlaceName
		},
	)
}
