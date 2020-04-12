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
	p.removeLabelsAndSaveVariables()
	p.changeVariableValuesIntoAssemblyLines()
}

func (p *ParseResult) removeLabelsAndSaveVariables() {
	labelLines := Filter(p.AssemblyLines, isLabel)
	labelNames := MapString(labelLines, parseLabelName)
	labelsMap := mapFrom(labelNames)
	p.AssemblyLines = Filter(p.AssemblyLines, isNotLabel)
	p.saveLabelIndexes(labelsMap)
	p.saveVariables(labelsMap)
}

func (p *ParseResult) saveLabelIndexes(labels map[string]bool) {
	for index, line := range p.AssemblyLines {
		possibleLabel := line[1:]
		if contains(labels, possibleLabel) {
			p.AddReservedSymbolWithValue(possibleLabel, strconv.Itoa(index))
		}
	}
}

func (p *ParseResult) saveVariables(labelsMap map[string]bool) {
	for _, line := range p.AssemblyLines {
		possibleVariable := line[1:]
		if line[0] == '@' && !contains(labelsMap, possibleVariable) && isVariable(possibleVariable) {
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
			if strings.HasPrefix(result, "@"+key+" ") || result == "@"+key{
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

func contains(mp map[string]bool, str string) bool {
	_, ok := mp[str]
	return ok
}

func parseLabelName(str string) string {
	strLen := len(str)
	return str[1 : strLen-1]
}

func mapFrom(labels []string) map[string]bool {
	var labelsMap = make(map[string]bool)
	for _, label := range labels {
		labelsMap[label] = true
	}
	return labelsMap
}

func isLabel(line string) bool {
	return strings.HasPrefix(line, "(")
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
