package json

import (
	"encoding/json"
	"log"
	"os"

	"github.com/fayaz07/locator/core/src/models"
	jsoniter "github.com/json-iterator/go"
)

func ParseWithStdJson(filePath string) {
	var records []models.Record

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	err = json.Unmarshal(content, &records)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	log.Println(len(records))
}

func ParseWithJsonIterator(filePath string) ([]models.Record, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var records []models.Record

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
		return nil, err
	}

	err = json.Unmarshal(content, &records)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
		return nil, err
	}

	return records, nil
}

func ParseLocationsWithJsonIterator(filePath string) ([]models.LocationModel, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var records []models.LocationModel

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
		return nil, err
	}

	err = json.Unmarshal(content, &records)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
		return nil, err
	}

	return records, nil
}
