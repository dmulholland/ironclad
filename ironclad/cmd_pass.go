package main


import "github.com/dmulholl/janus/v2"


import (
    "fmt"
    "os"
    "path/filepath"
)


var passHelp = fmt.Sprintf(`
Usage: %s pass <entry>

  Copy a stored password to the system clipboard or print it to stdout.

  The entry can be specified by its ID or by any unique set of case-insensitive
  substrings of its title.

Arguments:
  <entry>                   Entry ID or title.

Options:
  -f, --file <str>          Database file. Defaults to the last used file.

Flags:
  -h, --help                Print this command's help text and exit.
  -p, --print               Print the password to stdout.
  -r, --readable            Add spaces to the password for readability.
`, filepath.Base(os.Args[0]))


func registerPassCmd(parser *janus.ArgParser) {
    cmd := parser.NewCmd("pass", passHelp, passCallback)
    cmd.NewString("file f")
    cmd.NewFlag("readable r")
    cmd.NewFlag("print p")
}


func passCallback(parser *janus.ArgParser) {
    if !parser.HasArgs() {
        exit("missing entry argument")
    }
    filename, _, db := loadDB(parser)

    // Search for an entry corresponding to the specified arguments.
    list := db.Active().FilterByAll(parser.GetArgs()...)
    if len(list) == 0 {
        exit("no matching entry")
    } else if len(list) > 1 {
        println("Error: the query string matches multiple entries.")
        printCompact(list, len(db.Active()), filepath.Base(filename))
        os.Exit(1)
    }
    entry := list[0]

    // Add spaces if required.
    password := entry.GetPassword()
    if parser.GetFlag("readable") {
        password = addSpaces(password)
    }

    // Print the password to stdout.
    if parser.GetFlag("print") {
        fmt.Print(password)
        if stdoutIsTerminal() {
            fmt.Println()
        }
        return
    }

    // Copy the password to the clipboard.
    writeToClipboard(password)
}
