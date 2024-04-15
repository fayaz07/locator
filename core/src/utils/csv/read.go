package csv

import (
	"log"
	"os"

	"github.com/fayaz07/locator/core/src/models"
	"github.com/gocarina/gocsv"
)

func ReadAsciiIndexesCSVFile(file *os.File) []models.AsciiIndexModel {
	log.Println("Unmarshalling Ascii Indexes CSV file into list...")
	records := []models.AsciiIndexModel{}
	err := gocsv.UnmarshalFile(file, &records)
	if err != nil {
		panic(err)
	}
	log.Println("Unmarshal complete!")
	return records
}
