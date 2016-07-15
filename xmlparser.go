package main

import (
	"encoding/xml"
	"flag"
	"github.com/tealeg/xlsx"
	"os"
)

var filePath = flag.String("excel", "./toparse.xlsx", "Path to xlsx file")
var outputPath = flag.String("outputdir", "./parsed.xml", "Path to output folder")
var rootName = flag.String("topelement", "records", "The root element")
var recordName = flag.String("recordelement", "record", "The element name of each record")

type HeaderSlice []string

func generateTokens(xlFile *xlsx.File, topLevelName string, recordName string) []xml.Token {
	headers := make([]string, 0)
	startToken := xml.StartElement{Name: xml.Name{Local: topLevelName}}
	tokens := []xml.Token{startToken}
	sheet := xlFile.Sheets[0]
	for rowIndex, row := range sheet.Rows {
		if rowIndex == 0 {
			for _, cell := range row.Cells {
				headers = append(headers, cell.String())
			}
		} else {
			tokens = append(tokens, xml.StartElement{Name: xml.Name{"", recordName}})
			for cellIndex, cell := range row.Cells {
				t := xml.StartElement{Name: xml.Name{"", headers[cellIndex]}}
				tokens = append(tokens, t, xml.CharData(cell.String()), xml.EndElement{t.Name})
			}
			tokens = append(tokens, xml.EndElement{xml.Name{"", recordName}})
		}
	}
	tokens = append(tokens, xml.EndElement{startToken.Name})
	return tokens
}

func encodeTokens(tokens []xml.Token, file *os.File) error {
	enc := xml.NewEncoder(file)
	enc.Indent("", "  ")
	for _, t := range tokens {
		err := enc.EncodeToken(t)
		if err != nil {
			return err
		}
	}
	enc.Flush()
	return nil
}

func main() {
	flag.Parse()
	xlFile, _ := xlsx.OpenFile(*filePath)
	tokens := generateTokens(xlFile, *rootName, *recordName)
	fileOut, _ := os.Create(*outputPath)
	encodeTokens(tokens, fileOut)
}
