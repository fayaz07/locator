package main

import (
	"fmt"

	console "github.com/fayaz07/locator/cli/src/console"
)

const cliFunctionality = `
  This CLI lets you prepare the dataset that can be used by the Go Locator API(https://github.com/fayaz07/locator).
  If you haven't already downloaded the source dataset, please download it from the following link:
  https://data.opendatasoft.com/api/explore/v2.1/catalog/datasets/geonames-postal-code@public/exports/json 
  If the above link doesn't work, try exploring for updated dataset from the website's home page: 
  https://data.opendatasoft.com/
`

func intro() {
	console.InBox(100, "Welcome to Go Locator CLI")
	console.Info(cliFunctionality)
	console.Print("  Once you are ready to proceed, press any key to proceed or Q to Quit.")
}

func isUserRead() bool {
	input := console.ReadOneChar()
	if input == 'q' || input == 'Q' {
		console.Println("\n  Bye...")
		return false
	}
	return true
}

func reqSourceDatasetPath() {
	console.Print("\n\n  Please enter the path to the source dataset file: ")
	var path string
	n, err := fmt.Scanln(&path)
	if err != nil {
		console.Error("  Error reading input, please raise an issue at https://github.com/fayaz07/locator/issues")
	}
	if n == 0 {
		console.Error("  Path cannot be empty")
		reqSourceDatasetPath()
	}
	console.Info(fmt.Sprintf("  Source dataset path: %s", path))
}

func main() {
	intro()

	if !isUserRead() {
		return
	}

	reqSourceDatasetPath()
}
