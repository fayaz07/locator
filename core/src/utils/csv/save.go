package csv

import (
	"os"

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
