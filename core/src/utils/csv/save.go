package csv

import (
	"os"

	"github.com/fayaz07/locator/core/src/models"
	"github.com/gocarina/gocsv"
)

func SaveToFile(data interface{}, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	err = gocsv.MarshalFile(data, file)
	if err != nil {
		return err
	}
	return nil
}

func SaveHeaderToFile(file *os.File, location models.LocationModel) error {
	file.WriteString(location.GetHeader())
	return nil
}

func SaveSingleRecordToFile(file *os.File, location models.LocationModel) error {
	file.WriteString(location.GetRow())
	return nil
}

func SaveAsciiHeaderToFile(file *os.File, asciiIndex models.AsciiIndexModel) error {
	file.WriteString(asciiIndex.GetHeader())
	return nil
}

func SaveSingleAsciiRecordToFile(file *os.File, asciiIndex models.AsciiIndexModel) error {
	file.WriteString(asciiIndex.GetRow())
	return nil
}

func SaveAsciiRecordsToFile(file *os.File, asciiIndex []models.AsciiIndexModel) error {
	for _, record := range asciiIndex {
		file.WriteString(record.GetRow())
	}
	return nil
}
