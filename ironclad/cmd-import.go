package main


import "github.com/dmulholland/clio/go/clio"


import (
    "fmt"
    "os"
    "path/filepath"
    "io/ioutil"
)


// Help text for the 'import' command.
var importHelp = fmt.Sprintf(`
Usage: %s import [FLAGS] [OPTIONS] ARGUMENTS

  Import a list of entries in JSON format.

Arguments:
  <file>                    File to import.

Options:
  -f, --file <str>          Database file. Defaults to the last used file.

Flags:
      --help                Print this command's help text and exit.
`, filepath.Base(os.Args[0]))


// Callback for the 'import' command.
func importCallback(parser *clio.ArgParser) {

    // Make sure an argument has been specified.
    if !parser.HasArgs() {
        exit("you must supply the name of a file to import")
    }

    // Load the database.
    filename, password, db := loadDB(parser)

    // Read the JSON input file.
    input, err := ioutil.ReadFile(parser.GetArgs()[0])
    if err != nil {
        exit(err)
    }

    // Import the entries into the database.
    err = db.Import(input)
    if err != nil {
        exit(err)
    }
    // Save the updated database to disk.
    saveDB(filename, password, db)
}
