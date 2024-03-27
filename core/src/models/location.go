package models

type LocationModel struct {
	CountryCode string           `json:"cc" csv:"country_code"`
	PostalCode  string           `json:"z" csv:"postal_code"`
	PlaceName   string           `json:"a" csv:"place_name"`
	State       string           `json:"b" csv:"state"`
	City        string           `json:"d" csv:"city"`
	Block       string           `json:"f" csv:"block"`
	Latitude    float64          `json:"h" csv:"latitude"`
	Longitude   float64          `json:"i" csv:"longitude"`
	Accuracy    int              `json:"j" csv:"accuracy"`
	Coordinates CoordinatesModel `json:"m"`
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
		Coordinates: CoordinatesModel{
			Lon: record.Coordinates.Lon,
			Lat: record.Coordinates.Lat,
		},
	}
}
