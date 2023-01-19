package main

import (
	"fmt"
	"reflect"
	"strings"
)

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb" json:"name"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches"`
}

func describePlant(new_plant interface{}) {
	s := reflect.ValueOf(new_plant)
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		if len(s.Type().Field(i).Tag) > 0 {
			v := strings.ReplaceAll(string(s.Type().Field(i).Tag), ":", "=")
			new := strings.ReplaceAll(v, "\"", "")
			new = strings.ReplaceAll(new, " ", ", ")
			fmt.Print(s.Type().Field(i).Name, "(", new, ")", ":", f.Interface(), "\n")
		} else {
			fmt.Print(s.Type().Field(i).Name, ":", f.Interface(), "\n")
		}
	}
	fmt.Println()
}

func main() {
	fmt.Println()
	up1 := UnknownPlant{"rose", "triangle", 12}
	describePlant(up1)

	fmt.Println("----------------------")

	up2 := AnotherUnknownPlant{255, "triangle", 12}
	describePlant(up2)
	fmt.Println()
}
