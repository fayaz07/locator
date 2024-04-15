package prepare

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fayaz07/locator/core/src/models"
	"github.com/fayaz07/locator/core/src/utils/csv"
	stringUtils "github.com/fayaz07/locator/core/src/utils/string"
)

func generateAsciiIndexSliceForPlace(
	subStrings []string,
	index int,
) []models.AsciiIndexModel {
	asciiIndexSlice := []models.AsciiIndexModel{}
	for _, subString := range subStrings {
		placeName := strings.TrimSpace(subString)
		asciiIndexSlice = append(asciiIndexSlice, models.AsciiIndexModel{
			Name:   placeName,
			Code:   stringUtils.ConvertToAscii(placeName),
			Index:  index,
			Length: len(placeName),
		})
	}
	return asciiIndexSlice
}

func (p PrepareData) readAsciiFilesAndSort(files map[string]*os.File, basePath string, subDir string) error {
	for countryCode := range files {
		file := openFile(fmt.Sprintf("%s/%s/%s_ascii.csv", basePath, subDir, countryCode))
		log.Println("Processing sort of Ascii Indexes file for country code:", countryCode)
		err := p.readAsciiFileAndSort(countryCode, file)
		closeFile(file)
		if err != nil {
			log.Println("Failed to sort Ascii Indexes file for country code:", countryCode)
			return err
		}
	}
	return nil
}

func (PrepareData) readAsciiFileAndSort(countryCode string, file *os.File) error {
	log.Println("Reading Ascii Indexes file for country code:", countryCode)

	indexes := csv.ReadAsciiIndexesCSVFile(file)

	log.Println("Sorting Ascii Indexes file for country code:", countryCode, ", indexes:", len(indexes))
	sortAsciiIndexSlice(indexes)

	log.Println("Clearing contents of Ascii Indexes file for country code: ", countryCode)
	clearFile(file)

	file.Seek(0, 0)

	log.Println("Sort complete...")
	csv.SaveAsciiHeaderToFile(file, models.AsciiIndexModel{})

	log.Println("Saving sorted Ascii Indexes file for country code: ", countryCode)
	err := csv.SaveAsciiRecordsToFile(file, indexes)
	if err != nil {
		return err
	}

	log.Println("Saved Ascii Indexes file for country code: ", countryCode)
	return nil
}
