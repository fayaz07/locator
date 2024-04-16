package prepare

import "github.com/fayaz07/locator/core/src/models"

func (PrepareData) generateSearchIndexes(asciiIndexes []models.AsciiIndexModel) []models.SearchIndexModel {
	// from this sorted list of Ascii indexes
	// generate single record for similar Ascii codes
	// example: Ascii indexes are follows
	// #, Code, Length, Index
	// 1. 100,2,3
	// 2. 100,4,5
	// 3. 101,6,7
	// 4. 101,8,9
	// 5. 101,10,11
	// 6. 102,12,13
	// 7. 102,14,15
	// 8. 102,16,17
	// 9. 102,18,19
	// 10. 102,20,21
	// for the above Ascii indexes, the search indexes will be
	// # Code, AsciiStart, AsciiEnd
	// 1. 100,1,2
	// 2. 101,3,5
	// 3. 102,6,10

	result := []models.SearchIndexModel{}
	arrIndex := 1
	lineNumber := 1
	result = append(result, models.SearchIndexModel{
		Code:       asciiIndexes[0].Code,
		AsciiStart: lineNumber,
		AsciiEnd:   lineNumber,
	})

	for i := 1; i < len(asciiIndexes); i++ {
		if asciiIndexes[i].Code != asciiIndexes[i-1].Code {
			result[arrIndex-1].AsciiEnd = i
			lineNumber = i + 1
			result = append(result, models.SearchIndexModel{
				Code:       asciiIndexes[i].Code,
				AsciiStart: lineNumber,
				AsciiEnd:   lineNumber,
			})
			arrIndex++
		}
	}
	return result
}
