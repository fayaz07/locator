package search

import (
	"os"
	"strings"

	"github.com/fayaz07/locator/core/src/models"
	"github.com/fayaz07/locator/core/src/utils/csv"
)

var (
	inputDirPath string
	indexesMap   map[string][]models.SearchIndexModel = make(map[string][]models.SearchIndexModel)
)

func InitSearch(datasetPath string) {
	// read input directory for all files
	// read files with specific suffix
	inputDirPath = datasetPath + "/country"

	dir, err := os.Open(inputDirPath)
	if err != nil {
		panic(err)
	}

	files, err := dir.Readdir(-1)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), "_search.csv") {
			// read file
			countryName := strings.Split(file.Name(), "_")[0]

			searchIndexes := []models.SearchIndexModel{}
			csv.ReadFromFile(dir.Name()+"/"+file.Name(), &searchIndexes)
			indexesMap[countryName] = searchIndexes
		}
	}

	openAsciiFile("IN", datasetPath)
}
