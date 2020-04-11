package main

import (
	"fmt"
	"strings"
)

type ParseResult struct {
}

func getParser() *ParseResult {
	return new(ParseResult)
}

func (p *ParseResult) Parse(content string) string {
	var emptyLinesRemovedContent = getEmptyLinesRemovedContent(content)
	fmt.Println(emptyLinesRemovedContent)
	return string(len(emptyLinesRemovedContent))
}

func getEmptyLinesRemovedContent(content string) []string {
	var result []string
	var lineSeparator string
	var lines = strings.Split(content, "\n")
	for _, line := range lines {
		if isNotString(line) {
			result = append(result, lineSeparator+line)
		}
		lineSeparator = "\n"
	}
	return result
}

func isNotString(line string) bool {
	return len(line) != 0
}
