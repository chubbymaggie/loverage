package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var (
	reTestFunction = regexp.MustCompile(`^func Test_?([^\(]+)`)
)

func getTestFunctionsFromFiles(files []string) ([]string, error) {
	functions := []string{}
	for _, file := range files {
		parsed, err := getTestFunctionsFromFile(file)
		if err != nil {
			return nil, fmt.Errorf(
				"can't parse test functions from file %s: %s",
				file, err,
			)
		}

		functions = append(functions, parsed...)
	}

	return functions, nil
}

func getTestFunctionsFromFile(file string) ([]string, error) {
	functions := []string{}

	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("can't read file: %s", err)
	}

	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		matches := reTestFunction.FindStringSubmatch(line)
		if len(matches) > 0 {
			functions = append(functions, matches[1])
		}
	}

	return functions, nil
}
