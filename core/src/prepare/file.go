package prepare

import (
	"fmt"
	"log"
	"os"

	"github.com/fayaz07/locator/core/src/models"
	"github.com/fayaz07/locator/core/src/utils/csv"
)

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

func closeFile(file *os.File) {
	file.Close()
}

func closeFiles(context string, files map[string]*os.File) {
	log.Println("Closing files for ", context)
	for _, file := range files {
		file.Close()
	}
	log.Println("All files closed for ", context)
}

func clearFile(file *os.File) {
	file.Truncate(0)
}

func openFile(filePath string) *os.File {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	return file
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
