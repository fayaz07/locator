package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	core "github.com/fayaz07/locator/core/src"
	"github.com/fayaz07/locator/core/src/models"
	"github.com/fayaz07/locator/core/src/utils/csv"
	"github.com/fayaz07/locator/core/src/utils/json"
	stringUtils "github.com/fayaz07/locator/core/src/utils/string"
)

const (
	fileToRead = "data/ts_p.json"
)

func prepare() {
	// read data from file
	res, err := json.ParseLocationsWithJsonIterator(fileToRead)
	if err != nil {
		panic(err)
	}

	fmt.Println("Total records: ", len(res))
	fmt.Println("First record: ", res[0].PlaceName)

	// sanitized place names
	locations := []models.LocationModel{}

	names := []string{}
	asciiIndexSlice := []models.AsciiIndexModel{}

	for _, record := range res {
		record.PlaceName = strings.ToUpper(stringUtils.SanitizePlaceName(record))
		names = append(names, strings.ToUpper(record.PlaceName))

		locations = append(locations, record)
	}

	sort.SliceStable(names, func(i, j int) bool { return names[i] < names[j] })

	for i, record := range names {
		subStrings := stringUtils.GenerateSubstrings(record, 3, 7, 3)
		for _, subString := range subStrings {
			v := strings.TrimSpace(subString)
			asciiIndexSlice = append(asciiIndexSlice, models.AsciiIndexModel{
				// Name:   v,
				Code:   stringUtils.ConvertToAscii(v),
				Index:  i,
				Length: len(v),
			})
		}
	}

	sort.SliceStable(asciiIndexSlice, func(i, j int) bool { return asciiIndexSlice[i].Code < asciiIndexSlice[j].Code })

	fmt.Println("First record after sanitization: ", locations[0].PlaceName)
	fmt.Println("First record after sanitization(names): ", names[0], len(names))
	fmt.Println("First record after sanitization: ", asciiIndexSlice[0].Code)

	// save names to a file
	// json.SaveStringArray(names, "data/ts_p_names.json")
	// json.SaveIntArray(vectors, "data/ts_p_vectors.json")
	// json.SaveAsciiIndexArray(asciiIndexSlice, "data/csv/ascii.json")
	// json.SaveStringArray(names, "data/csv/names.json")
	csv.SaveToFile(asciiIndexSlice, "data/csv/ascii.csv")

	json.SaveStringArray(names, "data/csv/names.json")

	fmt.Println("saved")

	for {
		initUserBasedSearch(names, asciiIndexSlice)
	}
}

func initUserBasedSearch(names []string, asciiIndexSlice []models.AsciiIndexModel) {
	// ask user to input a place name for searching
	fmt.Print("\n\nEnter a place name: ")
	var search string
	fmt.Scanln(&search)

	start := time.Now()
	search = strings.ToUpper(search)

	fmt.Println("\nsearching for ", search, ", ascii code: ", stringUtils.ConvertToAscii(search))

	asciiCode := stringUtils.ConvertToAscii(search)
	// in2 := sort.SearchInts(vectors, asciiCode)
	in2 := core.SearchOnNames(asciiIndexSlice, asciiCode, len(search))
	fmt.Println("in2: ", in2)
	// fmt.Printf("asciis[in2]: %+v\n", vectors[in2])
	// fmt.Println("index: ", indexes[in2])
	// fmt.Println("names[in2]: ", names[indexes[in2]])

	// collect similar vectors
	// lps := stringUtils.ComputeLPSArray(search)

	possible := []string{}

	i := in2
	searchLen := len(search)
	for ; i < len(asciiIndexSlice); i++ {
		c := asciiIndexSlice[i].Code
		length := asciiIndexSlice[i].Length
		index := asciiIndexSlice[i].Index

		if c == asciiCode && length == searchLen {
			possible = append(possible, names[index])
		}
		if c != asciiCode {
			break
		}
	}
	fmt.Println("end index: ", i)

	suggestions := stringUtils.FindClosestStrings(possible, search)

	end := time.Now()
	timeTakenByVectorSearch := end.Sub(start)

	fmt.Println("possibilites: ", len(possible))
	// fmt.Println("possible: ", possible)
	fmt.Printf("suggestions: %v\n", strings.Join(suggestions, ", "))
	fmt.Println("time taken by vector search: ", timeTakenByVectorSearch)
}
