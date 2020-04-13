package main

import (
	"strconv"
	"testing"
)

type SymbolTableI interface {
	GetDestinationCode(dest string) (string, bool)
	GetJumperCode(dest string) (string, bool)
	GetExpressionCode(dest string) (string, bool)
	GetReservedSymbol(dest string) (string, bool)
	AddReservedSymbolWithValue(key string, value string)
	AddReservedSymbol(key string)
	GetReservedSymbols() map[string]string
}

func TestGetSymbolTable(t *testing.T) {
	var table SymbolTableI = GetSymbolTable()
	AssertNotNull(t, table)
}

func TestSymbolTable_GetDestinationCode(t *testing.T) {
	var table SymbolTableI = GetSymbolTable()
	_, exists := table.GetDestinationCode("blaaaa")
	AssertFalse(t, exists, "destinations should not contain [blaaaa]")

	code, exists := table.GetDestinationCode("AMD")
	AssertTrue(t, exists, "destinations should contain [AMD]")
	AssertEqualsString(t, code, "111")
}

func TestSymbolTable_GetJumperCode(t *testing.T) {
	var table SymbolTableI = GetSymbolTable()
	_, exists := table.GetJumperCode("blaaaa")
	AssertFalse(t, exists, "JumpersTable should not contain [blaaaa]")

	code, exists := table.GetJumperCode("JMP")
	AssertTrue(t, exists, "JumpersTable should contain [JMP]")
	AssertEqualsString(t, code, "111")
}

func TestSymbolTable_GetExpressionCode(t *testing.T) {
	var table SymbolTableI = GetSymbolTable()
	_, exists := table.GetExpressionCode("blaaaa")
	AssertFalse(t, exists, "ExpressionsTable should not contain [blaaaa]")

	code, exists := table.GetExpressionCode("D&M")
	AssertTrue(t, exists, "ExpressionsTable should contain [D&M]")
	AssertEqualsString(t, code, "1000000")
}

func TestSymbolTable_GetReservedSymbol(t *testing.T) {
	var table SymbolTableI = GetSymbolTable()
	_, exists := table.GetReservedSymbol("blaaaa")
	AssertFalse(t, exists, "ReservedSymbol should not contain [blaaaa]")

	code, exists := table.GetReservedSymbol("SCREEN")
	AssertTrue(t, exists, "ExpressionsTable should contain [SCREEN]")
	AssertEqualsString(t, code, "16384")

	code, exists = table.GetReservedSymbol("R10")
	AssertTrue(t, exists, "ExpressionsTable should contain [R10]")
	AssertEqualsString(t, code, "10")
}

func TestSymbolTable_AddReservedSymbolWithValue(t *testing.T) {
	var table SymbolTableI = GetSymbolTable()
	symbol := "blaaaa"
	_, exists := table.GetReservedSymbol(symbol)
	AssertFalse(t, exists, "ReservedSymbol should not contain [blaaaa]")

	table.AddReservedSymbolWithValue(symbol, symbol)

	code, exists := table.GetReservedSymbol(symbol)
	AssertTrue(t, exists, "ExpressionsTable should contain [SCREEN]")
	AssertEqualsString(t, code, symbol)
}

func TestSymbolTable_AddReservedSymbol(t *testing.T) {
	var table SymbolTableI = GetSymbolTable()

	symbol := "BLA"
	table.AddReservedSymbol(symbol)

	code, exists := table.GetReservedSymbol(symbol)
	AssertTrue(t, exists, "ExpressionsTable should contain [BLA]")
	AssertEqualsString(t, code, strconv.Itoa(16))

	symbol2 := "BLU"
	table.AddReservedSymbol(symbol2)

	code, exists = table.GetReservedSymbol(symbol2)
	AssertTrue(t, exists, "ExpressionsTable should contain [BLU]")
	AssertEqualsString(t, code, strconv.Itoa(17))
}

func TestSymbolTable_GetReservedSymbolKeys(t *testing.T) {
	var table SymbolTableI = GetSymbolTable()

	symbol := "BLA"
	table.AddReservedSymbol(symbol)

	reservedSymbols := table.GetReservedSymbols()
	AssertEqualsInt(t, 24, len(reservedSymbols))
}