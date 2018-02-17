package main

import (
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
	"io/ioutil"
	"log"
)

//writeSS takes the converted content from the CSV and writes it to
//the Google Sheet created by the function create.
func writeSS(ssid string, content [][]interface{}) {

	ctx := context.Background()
	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(ctx, config)
	sheetsService, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets Client %v", err)
	}

	spreadsheetId := ssid
	rangeData := "A1"
	values := content
	rb := &sheets.BatchUpdateValuesRequest{
		ValueInputOption: "USER_ENTERED",
	}
	rb.Data = append(rb.Data, &sheets.ValueRange{
		Range:  rangeData,
		Values: values,
	})
	_, err = sheetsService.Spreadsheets.Values.BatchUpdate(spreadsheetId, rb).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done.")
}
