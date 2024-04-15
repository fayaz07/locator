package prepare

import (
	"sort"

	"github.com/fayaz07/locator/core/src/models"
)

func sortAsciiIndexSlice(asciiIndexSlice []models.AsciiIndexModel) {
	sort.SliceStable(
		asciiIndexSlice,
		func(i, j int) bool {
			return asciiIndexSlice[i].Code < asciiIndexSlice[j].Code
		},
	)
}
