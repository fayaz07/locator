package prepare

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fayaz07/locator/core/src/models"
	"github.com/fayaz07/locator/core/src/utils/csv"
	"github.com/pterm/pterm"

	json "github.com/fayaz07/locator/core/src/utils/json"

	stringUtils "github.com/fayaz07/locator/core/src/utils/string"
)

var (
	srcFilesByCountry   = map[string]*os.File{}
	asciiFilesByCountry = map[string]*os.File{}
	recordsCount        = map[string]int{}
	totalRecords        = 0
)

const (
	srcFilePathByCountryTemplate   = "%s/country/%s_src.csv"
	asciiFilePathByCountryTemplate = "%s/country/%s_ascii.csv"
)

// Prepare dataset by countries
func (p PrepareData) PrepareDatasetByCountry() {
	startTime := time.Now()
	log.Println("Checking if the output directory exists")

	basePath := fmt.Sprintf("%s/%s", p.OutputPath, p.Mode.SubFolderPath())

	if _, err := os.Stat(basePath); os.IsNotExist(err) {
		log.Println("Creating output directory: ", basePath)
		err := os.MkdirAll(basePath, os.ModePerm)
		if err != nil {
			panic(fmt.Errorf("unable to create output directory: %s, error: %+v", basePath, err))
		}
	}

	area, _ := pterm.DefaultArea.Start()

	log.Println("Reading source dataset file: ", p.InputFilePath)

	logRecordsByCountry(area, "-", 0, 0)

	// read file and parse to RecordModel
	// RecordModel is the native json structure which the json file is expected to have
	err := json.ParseStreamedWithJsonIterator(
		p.InputFilePath,
		func(location models.LocationModel) {
			// increment total records
			totalRecords++

			// increment records count for the country
			if _, ok := recordsCount[location.CountryCode]; !ok {
				recordsCount[location.CountryCode] = 0
			}
			recordsCount[location.CountryCode]++

			p.processRecordForCountry(location, p.OutputPath, recordsCount[location.CountryCode])

			// update console
			logRecordsByCountry(area, location.CountryCode, recordsCount[location.CountryCode], totalRecords)
		},
	)
	if err != nil {
		panic(fmt.Errorf("unable to read source dataset file: %s, error: %+v", p.InputFilePath, err))
	}

	// for each country, print stats of records parsed
	for countryCode, countryRecords := range recordsCount {
		log.Println("Records parsed for country: ", countryCode, " count: ", countryRecords)
	}

	// sort the ascii index slice and save to file
	log.Println("Sorting Ascii Indexes files...")
	err = p.readAsciiFilesSortAndGenIndexes(asciiFilesByCountry, p.OutputPath, p.Mode.SubFolderPath())
	if err != nil {
		panic(err)
	}

	endTime := time.Now()
	log.Println("Total time taken: ", endTime.Sub(startTime))

	// close the area
	area.Stop()

	defer closeFiles("Country records source", srcFilesByCountry)
	defer closeFiles("Ascii records source", asciiFilesByCountry)
}

func logRecordsByCountry(area *pterm.AreaPrinter, countryCode string, countryRecords int, totalRecords int) {
	area.Update(pterm.Sprintf("Total Records Parsed: %d\nRecords parsed for %s, count: %d", totalRecords, countryCode, countryRecords))
}

func (p PrepareData) processRecordForCountry(location models.LocationModel, outputPath string, index int) {
	// look for country of the record and write it to the file
	// sanitize place name
	location.PlaceName = strings.ToUpper(stringUtils.SanitizePlaceName(location))

	// save to file
	writeLocationRecordToFileByCountry(location, outputPath)

	// generate substrings and ascii index
	subStrings := stringUtils.GenerateSubstrings(location.PlaceName, p.MinQueryLength, p.MaxQueryLength, 3)

	asciiIndexSlice := generateAsciiIndexSliceForPlace(subStrings, index)

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
