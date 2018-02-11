package main

import (
	"encoding/csv"
	"fmt"
	"os"
)


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
		all = append(all, []interface{}{value[0], value[1]})
	}

	return all
}