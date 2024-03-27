package models

type CoordinatesModel struct {
	Lon float64 `json:"k" csv:"longitude"`
	Lat float64 `json:"l" csv:"latitude"`
}
