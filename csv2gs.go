package main

import "flag"

var strCsv = flag.String("file", "", "Filename of the CSV you want to upload")
var sheetsName = flag.String("name", "", "Name of your Google Sheet document")
var parent = flag.String("parent", "", "Id of the parent folder where to place the sheet")

func main() {
	flag.Parse()
	fileName := *strCsv
	fileId := create(*sheetsName, *parent)
	content := prepare(fileName)
	writeSS(fileId, content)

}
