package main

import (
	"flag"
	"fmt"
)

var strCsv = flag.String("file", "", "Filename of the CSV you want to upload")
var sheetsName = flag.String("name", "", "Name of your Google Sheet document")
var parent = flag.String("parent", "", "Id of the parent folder where to place the sheet")

//main executes the pipeline of actions: Create the GSheets document,
//transforming the CSV contents and writing the final online content.
func main() {

	fmt.Println(strCsv)
	fmt.Println(sheetsName)
	fmt.Println(parent)

	flag.Parse()
	fileName := *strCsv
	fileId := create(*sheetsName, *parent)
	content := prepare(fileName)
	writeSS(fileId, content)

}
