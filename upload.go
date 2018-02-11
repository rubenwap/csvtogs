package main

import (
	"io/ioutil"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

func create(newFile string) string {
	ctx := context.Background()

	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved credentials
	// at ~/.credentials/drive-go-quickstart.json
	config, err := google.ConfigFromJSON(b, drive.DriveScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(ctx, config)

	srv, err := drive.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve drive Client %v", err)
	}

	// If we are uploading a file rather than creating, we need to add .Media(file)
	// The returned ID of the file is what we will need to write using the Sheets api
	driveFile, err := srv.Files.Create(&drive.File{Name: newFile, MimeType: "application/vnd.google-apps.spreadsheet"}).Do()
	if err != nil {
		log.Fatalf("Unable to create file: %v", err)
	}

	//log.Printf("file: %+v", driveFile)
	return driveFile.Id

}
