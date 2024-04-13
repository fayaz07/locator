package models

import "fmt"

type AsciiIndexModel struct {
	// Name   string `json:"n" csv:"name"`
	Code   int `json:"c" csv:"code"`
	Index  int `json:"i" csv:"index"`
	Length int `json:"l" csv:"length"`
}

func (asciiIndex AsciiIndexModel) GetHeader() string {
	return "code,index,length\n"
}

func (asciiIndex AsciiIndexModel) GetRow() string {
	return fmt.Sprintf(
		"%d,%d,%d\n",
		asciiIndex.Code, asciiIndex.Index, asciiIndex.Length)
}
