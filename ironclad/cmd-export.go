package main


import "github.com/dmulholland/janus-go/janus"


import (
    "fmt"
    "os"
    "path/filepath"
)


var exportHelp = fmt.Sprintf(`
Usage: %s export [FLAGS] [OPTIONS] [ARGUMENTS]

  Export a list of entries in JSON format. Entries can be specified by ID or
  by title. (Titles are checked for a case-insensitive substring match.) If no
  entries are specified, all entries will be exported.

  Output is written to stdout.

Arguments:
  [entries]                 List of entries to export by ID or title.

Options:
  -f, --file <str>          Database file. Defaults to the last used file.
  -t, --tag <str>           Filter entries using the specified tag.

Flags:
  -h, --help                Print this command's help text and exit.
`, filepath.Base(os.Args[0]))


func registerExportCmd(parser *janus.ArgParser) {
    cmd := parser.NewCmd("export", exportHelp, exportCallback)
    cmd.NewString("file f")
    cmd.NewString("tag t")
}


func exportCallback(parser *janus.ArgParser) {

    // Load the database.
    _, _, db := loadDB(parser)

    // Default to exporting all active entries.
    list := db.Active()

    // Do we have query strings to filter on?
    if parser.HasArgs() {
        list = list.FilterByAny(parser.GetArgs()...)
    }

    // Are we filtering by tag?
    if parser.GetString("tag") != "" {
        list = list.FilterByTag(parser.GetString("tag"))
    }

    // Create the JSON dump.
    dump, err := list.Export()
    if err != nil {
        exit(err)
    }

    // Print the JSON to stdout.
    fmt.Println(dump)
}
