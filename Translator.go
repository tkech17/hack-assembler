package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Translator struct {
}

func GetTranslator() *Translator {
	return &Translator{}
}

func (t *Translator) Translate(parser *ParseResult) string {
	var result []string
	for _, line := range parser.AssemblyLines {
		var code = getAssemblyLine(line, &parser.SymbolTable)
		result = append(result, code)
	}
	return strings.Join(result, "\n")
}

func getAssemblyLine(line string, symbolTable *SymbolTable) string {
	if isAType(line) {
		return get16IntBinaryString(line[1:])
	} else {
		return getCInstructionCode(line, symbolTable)
	}
}

func getCInstructionCode(line string, table *SymbolTable) string {
	result := "111"
	result += getComputationCode(line, table)
	result += getDestinationCode(line, table)
	result += getJumpCode(line, table)
	return result
}

func getComputationCode(line string, table *SymbolTable) string {
	computationExpression := getComputationExpression(line)
	code, exists := table.GetExpressionCode(computationExpression)
	if !exists {
		panic("Computation expression not found for line [" + line + "]")
	}
	return code
}

func getComputationExpression(line string) string {
	expressions := strings.Split(line, ";")
	computationPart := expressions[0]
	computations := strings.Split(computationPart, "=")
	if len(computations) == 1 {
		return computations[0]
	}
	return computations[1]
}

func getDestinationCode(line string, table *SymbolTable) string {
	destinationExpression := getDestinationExpression(line)
	code, _ := table.GetDestinationCode(destinationExpression)
	return code
}

func getDestinationExpression(line string) string {
	expressions := strings.Split(line, "=")
	if len(expressions) == 1 {
		return ""
	}
	return expressions[0]
}

func getJumpCode(line string, table *SymbolTable) string {
	jumpExpr := getJumpExpression(line)
	code, _ := table.GetJumperCode(jumpExpr)
	return code
}

func getJumpExpression(line string) string {
	expressions := strings.Split(line, ";")
	if len(expressions) == 1 {
		return ""
	}
	return expressions[1]
}

func get16IntBinaryString(s string) string {
	integer, _ := strconv.ParseInt(s, 10, 16)
	integer2base := fmt.Sprintf("%16b", integer)
	integer2base = strings.ReplaceAll(integer2base, " ", "0")
	return integer2base
}

func isAType(line string) bool {
	return line[0] == '@'
}
