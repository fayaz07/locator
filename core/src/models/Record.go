package models

type RecordModel struct {
	CountryCode string      `json:"country_code"`
	PostalCode  string      `json:"postal_code"`
	PlaceName   string      `json:"place_name"`
	AdminName1  string      `json:"admin_name1"` // state
	AdminCode1  string      `json:"admin_code1"` // state code
	AdminName2  string      `json:"admin_name2"` // city
	AdminCode2  string      `json:"admin_code2"` // city code
	AdminName3  string      `json:"admin_name3"` // block
	AdminCode3  interface{} `json:"admin_code3"` // block code
	Latitude    float64     `json:"latitude"`
	Longitude   float64     `json:"longitude"`
	Accuracy    int         `json:"accuracy"`
	Coordinates struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coordinates"`
}
