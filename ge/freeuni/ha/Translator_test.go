package main

import "testing"

func TestGetTranslator(t *testing.T) {
	translator := GetTranslator()
	AssertNotNull(t, translator)
}

func TestGet16IntBinaryString(t *testing.T) {
	AssertEqualsString(t, "0000000000000000", get16IntBinaryString("0"))
	AssertEqualsString(t, "0000000000000001", get16IntBinaryString("1"))
	AssertEqualsString(t, "0000000000000010", get16IntBinaryString("2"))
	AssertEqualsString(t, "0000000000000011", get16IntBinaryString("3"))
}

func TestIsAType(t *testing.T) {
	AssertTrue(t, isAType("@12"), "[@12] is A type")
	AssertFalse(t, isAType("D=1"), "[D=1] is not A type")
}

func TestGetAssemblyLine(t *testing.T) {
	table := GetSymbolTable()
	AssertEqualsString(t, "0000000000000000", getAssemblyLine("@0", nil))
	AssertEqualsString(t, "0000000000000001", getAssemblyLine("@1", nil))
	AssertEqualsString(t, "0000000000000010", getAssemblyLine("@2", nil))
	AssertEqualsString(t, "0000000000000011", getAssemblyLine("@3", nil))
	AssertEqualsString(t, "1110110000010000", getCInstructionCode("D=A", table))
}
func TestGetJumpExpression(t *testing.T) {
	AssertEqualsString(t, "JGT", getJumpExpression("D=2;JGT"))
	AssertEqualsString(t, "", getJumpExpression("D=A"))
}

func TestGetJumpCode(t *testing.T) {
	table := GetSymbolTable()
	AssertEqualsString(t, "001", getJumpCode("D=2;JGT", table))
	AssertEqualsString(t, "000", getJumpCode("D=A", table))
}

func TestGetDestinationCode(t *testing.T) {
	table := GetSymbolTable()
	AssertEqualsString(t, "010", getDestinationCode("D=2;JGT", table))
	AssertEqualsString(t, "111", getDestinationCode("AMD=A", table))
	AssertEqualsString(t, "000", getDestinationCode("temp; JMP", table))
}

func TestGetDestinationExpression(t *testing.T) {
	AssertEqualsString(t, "DA", getDestinationExpression("DA=2;JGT"))
	AssertEqualsString(t, "", getDestinationExpression("D; JME"))
}

func TestGetComputationExpression(t *testing.T) {
	AssertEqualsString(t, "1", getComputationExpression("DA=1;JGT"))
	AssertEqualsString(t, "M+1", getComputationExpression("DA=M+1"))
	AssertEqualsString(t, "-1", getComputationExpression("D=-1;JME"))
	AssertEqualsString(t, "0", getComputationExpression("0;JMP"))
}

func TestGetComputationCode(t *testing.T) {
	table := GetSymbolTable()
	AssertEqualsString(t, "0111111", getComputationCode("DA=1;JGT", table))
	AssertEqualsString(t, "1110111", getComputationCode("DA=M+1", table))
	AssertEqualsString(t, "0111010", getComputationCode("D=-1;JME", table))
	AssertEqualsString(t, "0101010", getComputationCode("0;JME", table))
}

func TestGetCInstructionCode(t *testing.T) {
	table := GetSymbolTable()
	AssertEqualsString(t, "1110110000010000", getCInstructionCode("D=A", table))
}
