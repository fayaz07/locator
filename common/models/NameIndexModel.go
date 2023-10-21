package models

type AsciiIndexModel struct {
	Name   string `json:"n" csv:"name"`
	Code   int    `json:"c" csv:"code"`
	Index  int    `json:"i" csv:"index"`
	Length int    `json:"l" csv:"length"`
}
