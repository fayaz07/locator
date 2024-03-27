package json

import (
	"os"

	"github.com/fayaz07/locator/core/src/models"
	jsoniter "github.com/json-iterator/go"
)

func SaveToFile(records []models.RecordModel, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := jsoniter.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(records)
	if err != nil {
		return err
	}

	return nil
}

func SaveToFileL(records []models.LocationModel, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := jsoniter.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(records)
	if err != nil {
		return err
	}

	return nil
}

func SaveStringArray(records []string, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := jsoniter.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(records)
	if err != nil {
		return err
	}

	return nil
}

func SaveIntArray(records []int, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := jsoniter.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(records)
	if err != nil {
		return err
	}

	return nil
}

func SaveAsciiIndexArray(records []models.AsciiIndexModel, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := jsoniter.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(records)
	if err != nil {
		return err
	}

	return nil
}
