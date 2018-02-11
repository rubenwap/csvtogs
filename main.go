package main

func main() {

	fileName := "hacker_news.csv"
	crawl(fileName)
	fileId := create("Upload Test")
	content := prepare(fileName)
	writeSS(fileId, content)

}
