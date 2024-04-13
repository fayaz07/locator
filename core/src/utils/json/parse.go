package json

import (
	"encoding/json"
	"log"
	"os"

	"github.com/fayaz07/locator/core/src/models"
	jsoniter "github.com/json-iterator/go"
)

func ParseWithStdJson(filePath string) {
	var records []models.RecordModel

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

func ParseWithJsonIterator(filePath string) ([]models.RecordModel, error) {
	var records []models.RecordModel

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

func ParseStreamedWithJsonIterator(
	filePath string,
	onEachRecord func(location models.LocationModel),
) error {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error opening file:", err)
		return err
	}
	defer file.Close()
	// Create a JSON decoder using jsoniter
	decoder := json.NewDecoder(file)

	decoder.UseNumber() // Handle large integer values

	//	decoder.Buffered().Read([]byte{100}) // Read the first byte to force the decoder to read the first JSON object
	decoder.Token() // Decode the first JSON object

	// Iterate over each JSON object in the file
	for decoder.More() {
		var item models.LocationModel

		// Decode the next JSON object
		if err := decoder.Decode(&item); err != nil {
			log.Println("Error decoding JSON:", err)
			return err
		}

		// Process the item
		onEachRecord(item)
	}

	decoder.Token()
	return nil
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
