# CSV to GSheets

Command line tool that takes your CSV and uploads it to Google Sheets. It doesn´t upload the pure CSV, but it creates a Google Sheet document instead. 

Ideally this is useful for batch processes where a CSV is generated elsewhere and you need to offer it via Google Sheets. 

This tool has been created for my very own use case, so your needs might vary. 

Also, this was a learning project for me to learn `Go`, so expect the code to don´t be as clean as it should ideally be. 

# Usage

Download a binary release or clone and compile with `go build`. Then do: 

    csv2gs -file "mysheet.csv" -name "Monthly Report" -parent "id of the targetGoogle Drive folder"
    
The first time you run it, you will be prompted to visit a Google URL to authenticate your credentials. You will be given a key that you must paste back into the terminal from where you ran this program. Then you are good to go. 
    
# Requirements

You must have a `client_secret.json` file from a Google project with permissions to use the Drive and Sheets API, in the same folder as the executable. 
Please see here for details: https://console.developers.google.com/start/api?id=drive



