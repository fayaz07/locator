package search

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/fayaz07/locator/core/src/models"
	stringUtils "github.com/fayaz07/locator/core/src/utils/string"
	"github.com/gocarina/gocsv"
)

func SearchByCountry(countryCode string, query string) ([]models.LocationModel, error) {
	// search if country exists
	if _, ok := indexesMap[countryCode]; !ok {
		return nil, fmt.Errorf("country not found")
	}

	// search for query
	searchIndexes := indexesMap[countryCode]

	// binary search
	log.Println("Searching for: ", query, " in country: ", countryCode, " with indexes: ", len(searchIndexes))

	asciiCode := stringUtils.ConvertToAscii(query)

	log.Println("Ascii code: ", asciiCode)

	// search for ascii code
	index := sort.Search(len(searchIndexes), func(i int) bool { return searchIndexes[i].Code >= asciiCode })
	if index >= len(searchIndexes) {
		return nil, fmt.Errorf("no records found")
	}

	searchIndexRecord := searchIndexes[index]
	log.Println("Found record: ", searchIndexRecord.Code, " with start: ", searchIndexRecord.AsciiStart, " and end: ", searchIndexRecord.AsciiEnd)

	// get all records, starting from index to ending index
	// asciiRecordsSlice := []models.AsciiIndexModel{}
	// csv.ReadFromFile(inputDirPath+"/"+countryCode+"_ascii.csv", &asciiRecordsSlice)

	// slicedRecords := asciiRecordsSlice[searchIndexRecord.AsciiStart:searchIndexRecord.AsciiEnd]
	// clear(asciiRecordsSlice)

	// slicedRecords := []models.AsciiIndexModel{}

	// for i := searchIndexRecord.AsciiStart; i < searchIndexRecord.AsciiEnd; i++ {
	// 	if asciiRecordsSlice[i].Length == len(query) {
	// 		slicedRecords = append(slicedRecords, asciiRecordsSlice[i])
	// 	}
	// }

	log.Println("Fetching sliced records")
	slicedRecords := readAsciiRecords(countryCode, int64(searchIndexRecord.AsciiStart), int64(searchIndexRecord.AsciiEnd))

	log.Println("Sliced records: ", len(slicedRecords))

	file, err := os.Create("acii.csv")
	if err != nil {
		panic(err)
	}
	gocsv.MarshalFile(slicedRecords, file)

	// for _, record := range slicedRecords {
	// 	log.Print("Index: ", record.GetRow())
	// }

	// search in sliced-ascii-records

	closestAsciis := stringUtils.FindClosestAsciis(slicedRecords, query)

	file, err = os.Create("closest.csv")
	if err != nil {
		panic(err)
	}
	gocsv.MarshalFile(closestAsciis, file)

	// get all records from location file
	locations := []models.LocationModel{}
	filePath := inputDirPath + "/" + countryCode + "_src.csv"

	log.Println("Reading location file: ", filePath)

	// csv.ReadFromFile(filePath, &locations)

	file, err = os.Open(filePath)
	if err != nil {
		panic(err)
	}

	err = gocsv.UnmarshalFile(file, &locations)
	if err != nil {
		panic(err)
	}

	log.Println("Total locations: ", len(locations))

	slicedLocations := []models.LocationModel{}

	testFile, _ := os.Create("test.csv")
	defer testFile.Close()

	testFile.WriteString(models.LocationModel{}.GetHeader())

	for _, record := range slicedRecords {
		// 	// locations = append(locations, models.LocationModel{})
		// 	file.Seek(int64(record.Index), 0)
		// 	line := scanner.Text()
		// 	locations = append(locations, models.GetLocationFromCSVRecord(line))
		slicedLocations = append(slicedLocations, locations[record.Index])
		testFile.WriteString(locations[record.Index].GetRow())
	}

	log.Println("Sliced locations: ", len(slicedLocations))

	return stringUtils.FindClosestLocations(slicedLocations, query), nil
}
