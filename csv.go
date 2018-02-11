package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

// This function takes a csv and returns it as slice of interfaces
// to be inserted in a Google Sheet

func prepare(filename string) [][]interface{} {

	file, _ := os.Open(filename)

	defer file.Close()

	reader := csv.NewReader(file)
	record, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error", err)
	}

	all := [][]interface{}{}

	for _, value := range record {
		row := []interface{}{}

		for _, item := range value {
			row = append(row, item)
		}

		all = append(all, row)

	}
	return all
}
