package main

import (
	"sort"
	"strings"
)

// SortTestFiles sorts test files by name ignoring suffix _test.go.
func SortTestFiles(files []string) {
	for i, file := range files {
		files[i] = strings.Replace(file, "_test.go", "", -1)
	}

	sort.Strings(files)

	for i, file := range files {
		files[i] = file + "_test.go"
	}
}
