package strings

import "github.com/fayaz07/locator/core/src/models"

func FindClosestLocations(arr []models.LocationModel, q string) []models.LocationModel {
	// use pattern matching to find closest strings
	lps := ComputeLPSArray(q)
	return searchSubstringInLocations(arr, q, lps)
}

func FindClosestAsciis(arr []models.AsciiIndexModel, q string) []models.AsciiIndexModel {
	// use pattern matching to find closest strings
	lps := ComputeLPSArray(q)
	return searchSubstringInAsciis(arr, q, lps)
}

func searchSubstringInAsciis(locations []models.AsciiIndexModel, pattern string, lps []int) []models.AsciiIndexModel {
	var results []models.AsciiIndexModel

	for _, location := range locations {
		startIndex := KMPSearch(location.Name, pattern, lps)
		if startIndex != -1 {
			results = append(results, location)
		}
	}

	return results
}

func searchSubstringInLocations(locations []models.LocationModel, pattern string, lps []int) []models.LocationModel {
	var results []models.LocationModel

	for _, location := range locations {
		startIndex := KMPSearch(location.PlaceName, pattern, lps)
		if startIndex != -1 {
			results = append(results, location)
		}
	}

	return results
}
