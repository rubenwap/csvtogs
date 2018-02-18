// Package csvtogs implements functionality to take a CSV file as input
// and create a Google Sheet with its contents.
// It is required that the user provides a client_secret.json for their
// own Google account.
package csvtogs

//Transform executes the pipeline of actions: Create the GSheets document,
//transforming the CSV contents and writing the final online content.
func Transform(file string, name string, parent string) {

	fileID := create(name, parent)
	content := prepare(file)
	writeSS(fileID, content)

}
