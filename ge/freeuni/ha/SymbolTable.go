package main

import (
	"strconv"
)

type SymbolTable struct {
	destinations    map[string]string
	jumpers         map[string]string
	expressions     map[string]string
	reservedSymbols map[string]string
	nextRegister    int
}

func GetSymbolTable() *SymbolTable {
	return &SymbolTable{
		getDestinationsMap(),
		getJumpersMap(),
		getExpressionsMap(),
		getReservedSymbolsMap(),
		16,
	}
}

func (s *SymbolTable) GetDestinationCode(dest string) (string, bool) {
	val, exists := s.destinations[dest]
	return val, exists
}

func (s *SymbolTable) GetJumperCode(dest string) (string, bool) {
	val, exists := s.jumpers[dest]
	return val, exists
}

func (s *SymbolTable) GetExpressionCode(dest string) (string, bool) {
	val, exists := s.expressions[dest]
	return val, exists
}

func (s *SymbolTable) GetReservedSymbol(dest string) (string, bool) {
	val, exists := s.reservedSymbols[dest]
	return val, exists
}

func (s *SymbolTable) AddReservedSymbolWithValue(key string, value string) {
	s.reservedSymbols[key] = value
}

func (s *SymbolTable) AddReservedSymbol(key string) {
	s.AddReservedSymbolWithValue(key, strconv.Itoa(s.nextRegister))
	s.nextRegister++
}

func (s *SymbolTable) GetReservedSymbols() map[string]string {
	return s.reservedSymbols
}


func getExpressionsMap() map[string]string {
	return map[string]string{
		"0":   "0101010",
		"1":   "0111111",
		"-1":  "0111010",
		"D":   "0001100",
		"A":   "0110000",
		"!D":  "0001101",
		"!A":  "0110001",
		"-D":  "0001111",
		"-A":  "0110011",
		"D+1": "0011111",
		"A+1": "0110111",
		"D-1": "0001110",
		"A-1": "0110010",
		"D+A": "0000010",
		"D-A": "0010011",
		"A-D": "0000111",
		"D&A": "0000000",
		"D|A": "0010101",
		"M":   "1110000",
		"!M":  "1110001",
		"-M":  "1110011",
		"M+1": "1110111",
		"M-1": "1110010",
		"D+M": "1000010",
		"D-M": "1010011",
		"M-D": "1000111",
		"D&M": "1000000",
		"D|M": "1010101",
	}
}

func getJumpersMap() map[string]string {
	return map[string]string{
		"":    "000",
		"JGT": "001",
		"JEQ": "010",
		"JGE": "011",
		"JLT": "100",
		"JNE": "101",
		"JLE": "110",
		"JMP": "111",
	}
}

func getDestinationsMap() map[string]string {
	return map[string]string{
		"":    "000",
		"M":   "001",
		"D":   "010",
		"MD":  "011",
		"A":   "100",
		"AM":  "101",
		"AD":  "110",
		"AMD": "111",
	}
}

func getReservedSymbolsMap() map[string]string {
	var reservedSymbols = map[string]string{
		"SCREEN": "16384",
		"KBD":    "24576",
		"SP":     "0",
		"LCL":    "1",
		"ARG":    "2",
		"THIS":   "3",
		"THAT":   "4",
	}
	for i := 1; i <= 15; i++ {
		iAsString := strconv.Itoa(i)
		reservedSymbols["R"+iAsString] = iAsString
	}
	return reservedSymbols
}
