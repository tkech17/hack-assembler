package main

import (
	"strconv"
	"strings"
)

type ParseResult struct {
	SymbolTable
	AssemblyLines []string
}

func getParser() *ParseResult {
	return &ParseResult{*GetSymbolTable(), nil}
}

func (p *ParseResult) Parse(content string) {
	p.AssemblyLines = strings.Split(content, "\n")
	p.AssemblyLines = Filter(p.AssemblyLines, isNotEmptyLine)
	p.AssemblyLines = Filter(p.AssemblyLines, isNotComment)
	p.AssemblyLines = MapString(p.AssemblyLines, removeCommentLine)
	p.AssemblyLines = MapString(p.AssemblyLines, strings.TrimSpace)
	p.AssemblyLines = MapString(p.AssemblyLines, removeSpaces)
	p.removeLabelsAndSaveVariables()
	p.changeVariableValuesIntoAssemblyLines()
}

func removeSpaces(str string) string {
	return strings.ReplaceAll(str, " ", "")
}

func (p *ParseResult) removeLabelsAndSaveVariables() {
	labelLines := Filter(p.AssemblyLines, isLabel)
	labelNames := MapString(labelLines, parseLabelName)
	labelsMap := mapFrom(labelNames)
	p.unionLabelWithLine()
	p.saveLabelIndexes()
	p.saveVariables(labelsMap)
	p.AssemblyLines = MapString(p.AssemblyLines, removeLabelFromLine)
}

func removeLabelFromLine(str string) string {
	line := strings.Split(str, ")")
	if len(line) == 1 {
		return str
	}
	return line[len(line)-1]
}

func (p *ParseResult) unionLabelWithLine() {
	var result []string
	lines := p.AssemblyLines
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if isLabel(line) {
			nextLine := lines[i+1]
			lines[i+1] = line + nextLine
			continue
		}
		result = append(result, line)
	}
	p.AssemblyLines = result
}

func (p *ParseResult) saveLabelIndexes() {
	for index, line := range p.AssemblyLines {
		labels := getLabelsFromLine(line)
		indexStr := strconv.Itoa(index)
		for _, label := range labels {
			p.AddReservedSymbolWithValue(label, indexStr)
		}
	}
}

func getLabelsFromLine(line string) []string {
	var result []string

	splitLine := strings.Split(line, ")")
	for _, str := range splitLine {
		if strings.HasPrefix(str, "(") {
			result = append(result, str[1:])
		}
	}
	return result
}

func (p *ParseResult) saveVariables(labelsMap map[string]string) {
	for _, line := range p.AssemblyLines {
		possibleVariable := line[1:]
		if line[0] == '@' && !contains(p.reservedSymbols, possibleVariable) && !contains(labelsMap, possibleVariable) && isVariable(possibleVariable) {
			p.AddReservedSymbol(possibleVariable)
		}
	}
}

func (p *ParseResult) changeVariableValuesIntoAssemblyLines() {
	symbols := p.GetReservedSymbols()
	p.AssemblyLines = MapString(p.AssemblyLines, getStringReplaceValuesFunc(symbols))
}

func getStringReplaceValuesFunc(symbols map[string]string) func(str string) string {
	return func(str string) string {
		var result = str
		for key, value := range symbols {
			result = strings.ReplaceAll(result, " "+key+" ", " "+value+" ")
			if strings.HasPrefix(result, key+" ") {
				result = strings.Replace(result, key, value, 1)
			}
			if strings.HasPrefix(result, "@"+key+" ") || result == "@"+key {
				result = strings.Replace(result, "@"+key, "@"+value, 1)
			}
			if strings.HasSuffix(result, " "+key) {
				result = result[0:len(result)-len(key)] + value
			}
		}
		return result
	}
}

func isVariable(variable string) bool {
	_, err := strconv.Atoi(variable)
	return err != nil
}

func contains(mp map[string]string, str string) bool {
	_, ok := mp[str]
	return ok
}

func parseLabelName(str string) string {
	strLen := len(str)
	return str[1 : strLen-1]
}

func mapFrom(labels []string) map[string]string {
	var labelsMap = make(map[string]string)
	for _, label := range labels {
		labelsMap[label] = ""
	}
	return labelsMap
}

func isLabel(line string) bool {
	return strings.HasPrefix(line, "(") && strings.HasSuffix(line, ")")
}

func isNotLabel(line string) bool {
	return !isLabel(line)
}

func removeCommentLine(str string) string {
	index := strings.Index(str, "//")
	if index == -1 {
		return str
	}
	return str[0:index]
}

func isNotEmptyLine(line string) bool {
	return len(strings.TrimSpace(line)) != 0
}

func isNotComment(line string) bool {
	trimmed := strings.TrimSpace(line)
	isComment := strings.HasPrefix(trimmed, "//")
	return !isComment
}
