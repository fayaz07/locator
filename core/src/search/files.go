package search

import (
	"fmt"
	"log"
	"os"
)

var (
	asciiFiles  = map[string]*os.File{}
	sourceFiles = map[string]*os.File{}
)

func openAsciiFile(countryCode string, inputPath string) error {
	if _, ok := asciiFiles[countryCode]; ok {
		return nil
	}

	file, err := os.Open(fmt.Sprintf("%s/country/%s_ascii.csv", inputPath, countryCode))
	if err != nil {
		log.Println("Unable to open file", err)
		return err
	}
	asciiFiles[countryCode] = file
	return nil
}

func openSourceFile(countryCode string, inputPath string) error {
	if _, ok := sourceFiles[countryCode]; ok {
		return nil
	}

	file, err := os.Open(fmt.Sprintf("%s/country/%s_src.csv", inputPath, countryCode))
	if err != nil {
		log.Println("Unable to open file", err)
		return err
	}
	sourceFiles[countryCode] = file
	return nil
}
