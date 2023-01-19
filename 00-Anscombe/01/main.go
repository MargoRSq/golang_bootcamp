package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/r3labs/diff/v3"
	"golang.org/x/exp/slices"
)

type Ingredient struct {
	IngredientName  string `xml:"itemname"   json:"ingredient_name"`
	IngredientCount string `xml:"itemcount"  json:"ingredient_count"`
	IngredientUnit  string `xml:"itemunit"   json:"ingredient_unit"`
}
type Cake struct {
	Name        string       `xml:"name"       json:"name"`
	Stovetime   string       `xml:"stovetime"  json:"time"`
	Ingredients []Ingredient `xml:"ingredients>item" json:"ingredients"`
}
type Recipes struct {
	XMLName xml.Name `xml:"recipes"    json:"-"`
	Cakes   []Cake   `xml:"cake" json:"cake"`
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
	var oldFilename string
	var newFilename string

	flag.StringVar(&oldFilename, "old", "", "--old old_filename.{json or xml}")
	flag.StringVar(&newFilename, "new", "", "--new new_filename.{json or xml}")
	flag.Parse()
	if flag.NFlag() == 2 {
		oldFile, _, err := parseFile(oldFilename)
		if err != nil {
			log.Fatalln("Not valid old file")
		}
		newFile, _, err := parseFile(newFilename)
		if err != nil {
			log.Fatalln("Not valid new file")
		}
		diffs := compareStructs(&oldFile, &newFile)
		for _, diff := range diffs {
			fmt.Println(diff)
		}
	} else {
		fmt.Println("Usage: ./compareDB --old original_database.xml --new stolen_database.json")
	}
}

func getCakesNames(recipes *Recipes) (names []string) {
	for _, cake := range recipes.Cakes {
		names = append(names, cake.Name)
	}
	return
}
func createDiffMessageCakes(firstNames []string, secondNames []string, whatHappend string) (messages []string) {
	for _, name := range firstNames {
		if !slices.Contains(secondNames, name) {
			messages = append(messages, fmt.Sprintf("%s cake \"%s\"", whatHappend, name))
		}
	}
	return
}

func createDiffMessageIngs(firstNames []string, secondNames []string, whatHappend string, cakeName string) (messages []string) {
	for _, name := range firstNames {
		if !slices.Contains(secondNames, name) {
			messages = append(messages, fmt.Sprintf("%s ingredient \"%s\" for cake \"%s\"", whatHappend, name, cakeName))
		}
	}
	return
}

func fetchDiffCakes(oldRecipes *Recipes, newRecipes *Recipes) (messages []string) {
	oldNames := getCakesNames(oldRecipes)
	newNames := getCakesNames(newRecipes)
	messages = append(messages, createDiffMessageCakes(newNames, oldNames, "ADDED")...)
	messages = append(messages, createDiffMessageCakes(oldNames, newNames, "REMOVED")...)
	return
}

func fetchDiffIngs(firstNames []string, secondNames []string, cakeName string) (messages []string) {
	messages = append(messages, createDiffMessageIngs(secondNames, firstNames, "ADDED", cakeName)...)
	messages = append(messages, createDiffMessageIngs(firstNames, secondNames, "REMOVED", cakeName)...)
	return
}

func compareIngredients(oldCake Cake, newCake Cake) (messages []string) {
	var oldIngsNames []string
	var newIngsNames []string
	for _, oldIng := range oldCake.Ingredients {
		oldIngsNames = append(oldIngsNames, oldIng.IngredientName)
		for _, newIng := range newCake.Ingredients {
			if !slices.Contains(newIngsNames, newIng.IngredientName) {
				newIngsNames = append(newIngsNames, newIng.IngredientName)
			}
			if oldIng.IngredientName == newIng.IngredientName {
				changelog, _ := diff.Diff(oldIng, newIng)
				for _, ch := range changelog {
					isTo := ch.To != ""
					isFrom := ch.From != ""
					switch fieldName := ch.Path[len(ch.Path)-1]; fieldName {
					case "IngredientCount":
						if isTo && isFrom {
							messages = append(messages, fmt.Sprintf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"", oldIng.IngredientName, oldCake.Name, ch.To, ch.From))
						} else if isFrom {
							messages = append(messages, fmt.Sprintf("REMOVED unit count \"%s\" for ingredient \"%s\" for cake \"%s\"", ch.From, oldIng.IngredientName, oldCake.Name))
						} else if isTo {
							messages = append(messages, fmt.Sprintf("ADDDED unit count \"%s\" for ingredient \"%s\" for cake \"%s\"", ch.To, oldIng.IngredientName, oldCake.Name))
						}
					case "IngredientUnit":
						if isTo && isFrom {
							messages = append(messages, fmt.Sprintf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"", oldIng.IngredientName, oldCake.Name, ch.To, ch.From))
						} else if isFrom {
							messages = append(messages, fmt.Sprintf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"", ch.From, oldIng.IngredientName, oldCake.Name))
						} else if isTo {
							messages = append(messages, fmt.Sprintf("ADDDED unit \"%s\" for ingredient \"%s\" for cake \"%s\"", ch.To, oldIng.IngredientName, oldCake.Name))
						}
					}
				}
			}
		}
	}
	messages = append(messages, fetchDiffIngs(oldIngsNames, newIngsNames, oldCake.Name)...)
	return
}

func compareCakes(oldCakes []Cake, newCakes []Cake) (messages []string) {
	for _, oldCake := range oldCakes {
		for _, newCake := range newCakes {
			if oldCake.Name == newCake.Name {
				changelog, _ := diff.Diff(oldCake, newCake)
				for _, ch := range changelog {
					if ch.Path[len(ch.Path)-1] == "Stovetime" {
						messages = append(messages, fmt.Sprintf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"", oldCake.Name, ch.To, ch.From))
					}
				}
				messages = append(messages, compareIngredients(oldCake, newCake)...)
			}
		}
	}
	return messages
}

func compareStructs(oldRecipes *Recipes, newRecipes *Recipes) (messages []string) {
	messages = append(messages, fetchDiffCakes(oldRecipes, newRecipes)...)
	messages = append(messages, compareCakes(oldRecipes.Cakes, newRecipes.Cakes)...)
	return
}

func parseFile(filename string) (r Recipes, filetype uint8, err error) {

	inputBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln("Error reading file")
	}
	if strings.HasSuffix(filename, ".xml") {
		filetype = XmlFile
		err = xml.Unmarshal(inputBytes, &r)
	} else if strings.HasSuffix(filename, ".json") {
		filetype = JsonFile
		err = json.Unmarshal(inputBytes, &r)
	} else {
		log.Fatalln("Not valid file extension")
	}
	return
}
