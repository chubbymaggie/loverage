package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"

	"github.com/docopt/docopt.go"
)

var (
	version = `1.0`
	usage   = `Loverage ` + version + `

Loverage is the utility for creating literal Go test coverage using function
names from _test.go files.

Usage:
    loverage [options]

Options:
    -f --files <path>   Specify test files that should be examined.
                         [default: *_test.go]
    -o --output <path>  Specify output file, use "-" value for stdout.
                         [default: -]
    -t --table          Output as markdown table.
    -h --help           Show this screen.
    -v --version        Show version.
`
)

const (
	ObjectMethodDelimiter = "."
)

func main() {
	args, err := docopt.Parse(usage, nil, true, version, true, true)
	if err != nil {
		panic(err)
	}

	var (
		sources     = args["--files"].(string)
		destination = args["--output"].(string)
		asTable     = args["--table"].(bool)
	)

	output := os.Stdout
	if destination != "-" {
		output, err = os.OpenFile(
			destination,
			os.O_CREATE|os.O_TRUNC|os.O_WRONLY,
			0666,
		)
		if err != nil {
			log.Fatalf("can't open output file: %s", err)
		}
	}
	defer func() {
		err = output.Close()
		if err != nil {
			log.Fatalf("can't close output file: %s", err)
		}
	}()

	files, err := filepath.Glob(sources)
	if err != nil {
		log.Fatalf("can't glob files: %s", err)
	}

	SortTestFiles(files)

	functions, err := getTestFunctionsFromFiles(files)
	if err != nil {
		log.Fatal(err)
	}

	table := tabwriter.NewWriter(output, 5, 4, 2, ' ', tabwriter.TabIndent)

	meanings := getMeanings(functions)
	for _, meaning := range meanings {
		columns := []string{}
		if meaning.object == "" {
			columns = append([]string{meaning.function}, meaning.suffixes...)
		} else {
			columns = append(
				[]string{
					meaning.object + ObjectMethodDelimiter + meaning.function,
				},
				meaning.suffixes...,
			)
		}

		if asTable {
			fmt.Fprintln(table, "| "+strings.Join(columns, "\t | ")+" |")
		} else {
			fmt.Fprintln(table, strings.Join(columns, "\t"))
		}
	}

	err = table.Flush()
	if err != nil {
		log.Fatalf("can't write table to stdout: %s", err)
	}
}
