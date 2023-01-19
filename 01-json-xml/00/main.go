package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type Recipes struct {
	XMLName xml.Name `xml:"recipes"    json:"-"`
	Cakes   []struct {
		Name        string `xml:"name"       json:"name"`
		Stovetime   string `xml:"stovetime"  json:"time"`
		Ingredients []struct {
			IngredientName  string  `xml:"itemname"   json:"ingredient_name"`
			IngredientCount float64 `xml:"itemcount"  json:"ingredient_count"`
			IngredientUnit  string  `xml:"itemunit"   json:"ingredient_unit"`
		} `xml:"ingredients>item" json:"ingredients"`
	} `xml:"cake" json:"cake"`
}

type Json struct{}
type Xml struct{}
type DBReader interface {
	Convert(*Recipes) ([]byte, error)
}

func (s *Json) Convert(storage *Recipes) (result []byte, err error) {
	result, err = xml.MarshalIndent(storage, "", "\t")
	return
}
func (s *Xml) Convert(storage *Recipes) (result []byte, err error) {
	result, err = json.MarshalIndent(storage, "", "\t")
	return
}

const (
	JsonFile = 1
	XmlFile  = 2
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("Error: ")
}
func main() {
	var filename string

	flag.StringVar(&filename, "f", "", "-f filename.{json or xml}")
	flag.Parse()
	if isFlagPassed("f") {
		recipesStruct, inputType := parseFile(filename)
		convertedBytes, _ := convertFile(&recipesStruct, inputType)
		fmt.Println(string(convertedBytes))
	} else {
		fmt.Println("No filename flag specified")
	}
}

func convertFile(recipes *Recipes, inputType uint8) (result []byte, err error) {
	var dbr DBReader
	var sJson *Json
	var sXml *Xml
	if inputType == JsonFile {
		dbr = sJson
	} else if inputType == XmlFile {
		dbr = sXml
	}
	result, err = dbr.Convert(recipes)
	return
}

func parseFile(filename string) (r Recipes, filetype uint8) {

	inputBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln("Error reading file")
	}
	if strings.HasSuffix(filename, ".json") {
		filetype = JsonFile
		json.Unmarshal(inputBytes, &r)
	} else if strings.HasSuffix(filename, ".xml") {
		filetype = XmlFile
		xml.Unmarshal(inputBytes, &r)
	} else {
		log.Fatalln("Not valid file extension")
	}
	return
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
