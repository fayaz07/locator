package prepare

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/fayaz07/locator/core/src/models"
	"github.com/fayaz07/locator/core/src/utils/csv"
	"github.com/pterm/pterm"

	json "github.com/fayaz07/locator/core/src/utils/json"

	stringUtils "github.com/fayaz07/locator/core/src/utils/string"
)

var (
	srcFilesByCountry      = map[string]*os.File{}
	asciiFilesByCountry    = map[string]*os.File{}
	recordsCount           = map[string]int{}
	totalRecords           = 0
	indexOfCurrentLocation = 0
)

const (
	srcFilePathByCountryTemplate   = "%s/country/%s_src.csv"
	asciiFilePathByCountryTemplate = "%s/country/%s_ascii.csv"
)

// Prepare dataset by countries
func PrepareDatasetByCountry(filePath string, outputPath string) {
	startTime := time.Now()
	log.Println("Checking if the output directory exists")
	if _, err := os.Stat(outputPath + "/country"); os.IsNotExist(err) {
		log.Println("Creating output directory: ", outputPath+"/country")
		err := os.MkdirAll(outputPath+"/country", os.ModePerm)
		if err != nil {
			panic(fmt.Errorf("unable to create output directory: %s, error: %+v", outputPath, err))
		}
	}

	area, _ := pterm.DefaultArea.Start()

	// defer consoleWriter.Stop()
	defer closeAllFiles()

	log.Println("Reading source dataset file: ", filePath)

	logRecordsByCountry(area, "-", 0, 0)

	// read file and parse to RecordModel
	// RecordModel is the native json structure which the json file is expected to have
	err := json.ParseStreamedWithJsonIterator(
		filePath,
		func(location models.LocationModel) {

			// increment total records
			totalRecords++

			// increment records count for the country
			if _, ok := recordsCount[location.CountryCode]; !ok {
				recordsCount[location.CountryCode] = 0
			}
			recordsCount[location.CountryCode]++

			processEachRecord(location, outputPath)

			// update console
			logRecordsByCountry(area, location.CountryCode, recordsCount[location.CountryCode], totalRecords)
		},
	)
	if err != nil {
		panic(fmt.Errorf("unable to read source dataset file: %s, error: %+v", filePath, err))
	}

	// for each country, print stats of records parsed
	for countryCode, countryRecords := range recordsCount {
		log.Println("Records parsed for country: ", countryCode, " count: ", countryRecords)
	}

	endTime := time.Now()
	log.Println("Total time taken: ", endTime.Sub(startTime))

	// close the area
	area.Stop()
}

func logRecordsByCountry(area *pterm.AreaPrinter, countryCode string, countryRecords int, totalRecords int) {
	area.Update(pterm.Sprintf("Total Records Parsed: %d\nRecords parsed for %s, count: %d", totalRecords, countryCode, countryRecords))
}

func closeAllFiles() {
	log.Println("Closing all files")
	for _, file := range srcFilesByCountry {
		file.Close()
	}
	for _, file := range asciiFilesByCountry {
		file.Close()
	}
	log.Println("All files closed")
}

func processEachRecord(location models.LocationModel, outputPath string) {
	// look for country of the record and write it to the file
	// sanitize place name
	location.PlaceName = strings.ToUpper(stringUtils.SanitizePlaceName(location))

	// save to file
	writeLocationRecordToFileByCountry(location, outputPath)

	// generate substrings and ascii index
	subStrings := stringUtils.GenerateSubstrings(location.PlaceName, 3, 7, 3)
	indexOfCurrentLocation++
	asciiIndexSlice := generateAsciiIndexSliceForPlace(subStrings, indexOfCurrentLocation)

	// save ascii indexes list to file
	writeAsciiIndexSliceToFile(asciiIndexSlice, location.CountryCode, outputPath)
}

func writeAsciiIndexSliceToFile(asciiIndexSlice []models.AsciiIndexModel, countryCode string, outputPath string) {
	file := getFileForStoringAsciiCountryRecords(countryCode, outputPath)
	err := csv.SaveAsciiRecordsToFile(file, asciiIndexSlice)
	if err != nil {
		panic(fmt.Errorf(
			"unable to save ascii index data to file: %s, for country: %s, error: %+v",
			outputPath, countryCode, err,
		))
	}
}

func getFileForStoringAsciiCountryRecords(countryCode string, outputPath string) *os.File {
	if file, ok := asciiFilesByCountry[countryCode]; ok {
		return file
	}

	file := createFile(fmt.Sprintf(asciiFilePathByCountryTemplate, outputPath, countryCode))
	asciiFilesByCountry[countryCode] = file
	csv.SaveAsciiHeaderToFile(file, models.AsciiIndexModel{})
	return file
}

func writeLocationRecordToFileByCountry(location models.LocationModel, outputPath string) {
	file := getFileForStoringSrcCountryRecords(location.CountryCode, outputPath)
	err := csv.SaveSingleRecordToFile(file, location)
	if err != nil {
		panic(fmt.Errorf(
			"unable to save location data to file: %s, for country: %s, error: %+v",
			outputPath, location.CountryCode, err,
		))
	}
}

func getFileForStoringSrcCountryRecords(countryCode string, outputPath string) *os.File {
	if file, ok := srcFilesByCountry[countryCode]; ok {
		return file
	}

	file := createFile(fmt.Sprintf(srcFilePathByCountryTemplate, outputPath, countryCode))
	srcFilesByCountry[countryCode] = file
	csv.SaveHeaderToFile(file, models.LocationModel{})
	return file
}

func createFile(fileName string) *os.File {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	return file
}

func generateAsciiIndexSliceForPlace(
	subStrings []string,
	index int,
) []models.AsciiIndexModel {
	asciiIndexSlice := []models.AsciiIndexModel{}
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

func sortAsciiIndexSlice(asciiIndexSlice []models.AsciiIndexModel) {
	sort.SliceStable(
		asciiIndexSlice,
		func(i, j int) bool {
			return asciiIndexSlice[i].Code < asciiIndexSlice[j].Code
		},
	)
}
