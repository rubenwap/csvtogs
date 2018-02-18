# CSV to GSheets

Package that takes your CSV and uploads it to Google Sheets. It doesn´t upload the pure CSV, but it creates a Google Sheet document instead. 

Ideally this is useful for batch processes where a CSV is generated elsewhere and you need to offer it via Google Sheets. 

This tool has been created for my very own use case, so your needs might vary. 

Also, this was a learning project for me to learn `Go`, so expect the code to don´t be as clean as it should ideally be. 

# Usage

Get it with 

`go get github.com/rubenwap/csvtogs`

And use it with 

`import github.com/rubenwap/csvtogs`

and...

`csvtogs.Transform("path/to/file.csv", "Name of File", "parent_folder_Gdrive_id")`

The first time you run it, you will be prompted to visit a Google URL to authenticate your credentials. You will be given a key that you must paste back into the terminal from where your code is running. Then you are good to go. 
    
# Requirements

You must have a `client_secret.json` file from a Google project with permissions to use the Drive and Sheets API, in the same folder as the executable. 
Please see here for details: https://console.developers.google.com/start/api?id=drive



