package models

import (
	"fmt"
	"strings"
)

type LocationModel struct {
	CountryCode string  `json:"country_code" csv:"country_code"`
	PostalCode  string  `json:"postal_code" csv:"postal_code"`
	PlaceName   string  `json:"place_name" csv:"place_name"`
	State       string  `json:"admin_name1" csv:"state"`
	City        string  `json:"admin_name2" csv:"city"`
	Block       string  `json:"admin_name3" csv:"block"`
	Latitude    float64 `json:"latitude" csv:"latitude"`
	Longitude   float64 `json:"longitude" csv:"longitude"`
	Accuracy    int     `json:"accuracy" csv:"accuracy"`
}

func RemoveCommas(str string) string {
	return strings.ReplaceAll(str, ",", " ")
}

func (location *LocationModel) StripCommas() {
	location.CountryCode = RemoveCommas(location.CountryCode)
	location.PostalCode = RemoveCommas(location.PostalCode)
	location.PlaceName = RemoveCommas(location.PlaceName)
	location.State = RemoveCommas(location.State)
	location.City = RemoveCommas(location.City)
	location.Block = RemoveCommas(location.Block)
}

func MapRecordToLocationModel(record RecordModel) LocationModel {
	return LocationModel{
		CountryCode: record.CountryCode,
		PostalCode:  record.PostalCode,
		PlaceName:   record.PlaceName,
		State:       record.AdminName1,
		City:        record.AdminName2,
		Block:       record.AdminName3,
		Latitude:    record.Latitude,
		Longitude:   record.Longitude,
		Accuracy:    record.Accuracy,
	}
}

func (location LocationModel) GetHeader() string {
	return "country_code,postal_code,place_name,state,city,block,latitude,longitude,accuracy\n"
}

func (location LocationModel) GetRow() string {
	return fmt.Sprintf(
		"%s,%s,%s,%s,%s,%s,%f,%f,%d\n",
		location.CountryCode, location.PostalCode, location.PlaceName,
		location.State, location.City, location.Block, location.Latitude,
		location.Longitude, location.Accuracy)
}

func GetLocationFromCSVRecord(record string) LocationModel {
	var location LocationModel
	fmt.Sscanf(
		record,
		"%s,%s,%s,%s,%s,%s,%f,%f,%d",
		&location.CountryCode, &location.PostalCode, &location.PlaceName,
		&location.State, &location.City, &location.Block, &location.Latitude,
		&location.Longitude, &location.Accuracy)
	return location
}
