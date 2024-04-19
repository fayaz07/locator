package search

import (
	"bufio"
	"fmt"
	"log"

	"github.com/fayaz07/locator/core/src/models"
	"github.com/gocarina/gocsv"
)

func readAsciiRecords(
	countryCode string,
	startIndex int64,
	endIndex int64,
) []models.AsciiIndexModel {
	file, ok := asciiFiles[countryCode]
	if !ok {
		panic(fmt.Errorf("file not read for countryCode: %s", countryCode))
	}

	// as file will also include headers
	file.Seek(startIndex+1, 0)

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	currentLine := startIndex + 1

	result := []models.AsciiIndexModel{}
	end := currentLine + 10
	for currentLine <= end {
		record := models.AsciiIndexModel{}
		line := scanner.Text()
		log.Println("line", currentLine, "is", line)
		gocsv.UnmarshalString(line, &record)
		result = append(result, record)
		currentLine++
	}
	return result
}
