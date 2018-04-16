package main


import "github.com/dmulholland/args"


import (
    "fmt"
    "os"
    "path/filepath"
    "io/ioutil"
)


var importHelp = fmt.Sprintf(`
Usage: %s import [FLAGS] [OPTIONS] ARGUMENTS

  Import a list of entries in JSON format.

  You can specify the name of a file to import. If no filename is specified,
  input is read from stdin.

Arguments:
  <file>                    File to import. Defaults to stdin.

Options:
  -f, --file <str>          Database file. Defaults to the last used file.

Flags:
  -h, --help                Print this command's help text and exit.
`, filepath.Base(os.Args[0]))


func importCallback(parser *args.ArgParser) {

    // Read the JSON input.
    var input []byte
    var err error
    if parser.HasArgs() {
        input, err = ioutil.ReadFile(parser.GetArg(0))
        if err != nil {
            exit(err)
        }
    } else {
        input = []byte(inputViaStdin())
    }

    // Load the database.
    filename, password, db := loadDB(parser)

    // Import the entries into the database.
    err = db.Import(input)
    if err != nil {
        exit(err)
    }

    // Save the updated database to disk.
    saveDB(filename, password, db)
}
