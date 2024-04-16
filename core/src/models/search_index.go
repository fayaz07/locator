package models

type SearchIndexModel struct {
	Code int `json:"code"`
	// afi - ascii file index
	// refers to the line number in the ascii indexes file
	AsciiStart int `json:"afi_start"`
	AsciiEnd   int `json:"afi_end"`
}
