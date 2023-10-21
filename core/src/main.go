package core

import (
	"log"
	"time"

	"github.com/fayaz07/locator/utils/src/json"
)

const filePath = "./data.json"

func main() {
	start := time.Now()
	json.ParseWithStdJson(filePath)
	end := time.Now()

	log.Println("Time taken by std json: ", end.Sub(start))

	start = time.Now()
	json.ParseWithJsonIterator(filePath)
	end = time.Now()

	log.Println("Time taken by json iterator: ", end.Sub(start))
}
