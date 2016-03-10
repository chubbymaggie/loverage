package main

import (
	"strings"
	"unicode"
)

var (
	// copy from golang/lint
	initialisms = map[string]bool{
		"API":   true,
		"ASCII": true,
		"CPU":   true,
		"CSS":   true,
		"DNS":   true,
		"EOF":   true,
		"GUID":  true,
		"HTML":  true,
		"HTTP":  true,
		"HTTPS": true,
		"ID":    true,
		"IP":    true,
		"JSON":  true,
		"LHS":   true,
		"QPS":   true,
		"RAM":   true,
		"RHS":   true,
		"RPC":   true,
		"SLA":   true,
		"SMTP":  true,
		"SQL":   true,
		"SSH":   true,
		"TCP":   true,
		"TLS":   true,
		"TTL":   true,
		"UDP":   true,
		"UI":    true,
		"UID":   true,
		"UUID":  true,
		"URI":   true,
		"URL":   true,
		"UTF8":  true,
		"VM":    true,
		"XML":   true,
		"XSRF":  true,
		"XSS":   true,
	}
)

const (
	ShortestInitialismLength = 2
	LongestInitialismLength  = 5
)

type Meaning struct {
	object   string
	function string
	suffixes []string
}

func getMeanings(functions []string) []Meaning {
	meanings := []Meaning{}

	for _, function := range functions {
		meaning := getMeaning(function)
		meanings = append(meanings, meaning)
	}

	return meanings
}

func getMeaning(function string) Meaning {
	parts := strings.Split(function, "_")

	// Algrorithm depends on number of parts
	// 1 - Function|Object
	// 2 - Function|Object + Description
	// 3+ - Object + ObjectMethod + Description

	var meaning Meaning
	switch len(parts) {
	case 1:
		meaning.function = parts[0]

	case 2:
		meaning.function = parts[0]
		meaning.suffixes = []string{expandCamelCase(parts[1])}

	default:
		meaning.object = parts[0]
		meaning.function = parts[1]
		meaning.suffixes = expandCamelCaseMany(parts[2:])
	}

	return meaning
}

func expandCamelCaseMany(words []string) []string {
	expanded := []string{}
	for _, word := range words {
		expanded = append(expanded, expandCamelCase(word))
	}

	return expanded
}

func expandCamelCase(str string) string {
	letters := []rune(str)
	words := []string{}

	lastPosition := 0
	for position := 0; position < len(letters); position++ {
		if position > 0 && unicode.IsUpper(letters[position]) {
			initialism, ok := isBeginsWithInitialism(str[lastPosition:])
			if ok {
				words = append(words, initialism)

				position += len(initialism) - 1
				lastPosition = position

				continue
			}

			words = append(words, str[lastPosition:position])
			lastPosition = position
		}
	}

	if str[lastPosition:] != "" {
		words = append(words, str[lastPosition:])
	}

	return strings.Join(words, " ")
}

func isBeginsWithInitialism(str string) (string, bool) {
	for i := ShortestInitialismLength; i <= LongestInitialismLength; i++ {
		if len(str) > i-1 && initialisms[str[:i]] {
			return str[:i], true
		}
	}

	return "", false
}
