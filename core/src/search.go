package core

import (
	"sort"

	"github.com/fayaz07/locator/core/src/models"
)

func SearchOnNames(names []models.AsciiIndexModel, query int, length int) int {
	return sort.Search(len(names), func(i int) bool { return names[i].Code >= query })
}
