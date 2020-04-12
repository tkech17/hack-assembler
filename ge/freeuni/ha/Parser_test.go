package main

import (
	"testing"
)

func Test_should_returnIsEmptyString_when_givenEmptyString(t *testing.T) {
	var s = ""
	var isNotEmptyString = isNotEmptyLine(s)
	AssertFalse(t, isNotEmptyString, "string ["+s+"] is not empty line")
}

func Test_should_returnIsNotEmptyString_when_givenNonEmptyString(t *testing.T) {
	var s = "abc"
	var isNotEmptyString = isNotEmptyLine(s)
	AssertTrue(t, isNotEmptyString, "string ["+s+"] is empty line")
}

func TestIsNotComment(t *testing.T) {
	AssertTrue(t, isNotComment("blaa"), "[blaa] is not comment")
	AssertTrue(t, isNotComment(""), "[$emptyLine] is not comment")
	AssertTrue(t, isNotComment("   blaa "), "[   blaa ] is not comment")
	AssertFalse(t, isNotComment("//"), "[//] is comment")
	AssertFalse(t, isNotComment("    //"), "[    //] is comment")
	AssertFalse(t, isNotComment("  //  sdfdsfsd "), "[  //  sdfdsfsd ] is comment")
}

func TestRemoveCommentLine(t *testing.T) {
	AssertEqualsString(t, "bla", removeCommentLine("bla"))
	AssertEqualsString(t, "bla", removeCommentLine("bla//"))
	AssertEqualsString(t, "bla", removeCommentLine("bla//asdsadsad"))
}

func TestIsLabel(t *testing.T) {
	AssertTrue(t, isLabel("(bla)"), "[(bla)] Is Label")
	AssertFalse(t, isLabel("bla"), "[bla] Is Not Label")
}

func TestIsNotLabel(t *testing.T) {
	AssertFalse(t, isNotLabel("(bla)"), "[(bla)] Is Label")
	AssertTrue(t, isNotLabel("bla"), "[bla] Is Not Label")
}

func TestParseLabelName(t *testing.T) {
	AssertEqualsString(t, "LABEL", parseLabelName("(LABEL)"))
}

func TestSaveLabelIndexes(t *testing.T) {
	result := ParseResult{*GetSymbolTable(), []string{"@LABEL"}}
	labels := map[string]bool{"LABEL": true}
	result.saveLabelIndexes(labels)

	symbol, exists := result.GetReservedSymbol("LABEL")

	AssertTrue(t, exists, "LABEL should be in Symbols")
	AssertEqualsString(t, "0", symbol)
}

func TestSaveVariables(t *testing.T) {
	result := ParseResult{*GetSymbolTable(), []string{"@LABEL", "D=M", "@12", "@temp", "@temp2"}}
	labels := map[string]bool{"LABEL": true}
	result.saveVariables(labels)

	symbol, exists := result.GetReservedSymbol("temp")
	AssertTrue(t, exists, "temp should be in Symbols")
	AssertEqualsString(t, "16", symbol)

	symbol, exists = result.GetReservedSymbol("temp2")

	AssertTrue(t, exists, "temp2 should be in Symbols")
	AssertEqualsString(t, "17", symbol)

}

func TestChangeVariableValuesIntoAssemblyLines(t *testing.T) {
	result := ParseResult{*GetSymbolTable(), []string{"(LABEL)", "@LABEL", "D=M", "@12", "@temp", "@temp2"}}
	result.removeLabelsAndSaveVariables()

	result.changeVariableValuesIntoAssemblyLines()

	lines := result.AssemblyLines
	AssertEqualsString(t, "@0", lines[0])
	AssertEqualsString(t, "D=M", lines[1])
	AssertEqualsString(t, "@12", lines[2])
	AssertEqualsString(t, "@16", lines[3])
	AssertEqualsString(t, "@17", lines[4])
}

func TestRemoveLabelsAndSaveVariables(t *testing.T) {
	result := ParseResult{*GetSymbolTable(), []string{"(LABEL)", "@LABEL", "@12", "D=1", "@b!&a"}}
	result.removeLabelsAndSaveVariables()

	AssertEqualsInt(t, 4, len(result.AssemblyLines))
	AssertEqualsString(t, "@LABEL", result.AssemblyLines[0])

	symbol, exists := result.GetReservedSymbol("LABEL")

	AssertTrue(t, exists, "LABEL should be in Symbols")
	AssertEqualsString(t, "0", symbol)

	symbol, exists = result.GetReservedSymbol("b!&a")

	AssertTrue(t, exists, "b!&a should be in Symbols")
	AssertEqualsString(t, "16", symbol)
}

func TestContains(t *testing.T) {
	labels := map[string]bool{"LABEL": true}
	containsLabel := contains(labels, "LABEL")
	containsBla := contains(labels, "BLA")

	AssertFalse(t, containsBla, "Should Not Contain [BLA]")
	AssertTrue(t, containsLabel, "Should Contain [LABEL]")
}

func TestIsVariable(t *testing.T) {
	AssertTrue(t, isVariable("BLA"), "[BLA] is Variable")
	AssertFalse(t, isVariable("123312"), "[123312] is not Variable")
	AssertTrue(t, isVariable("BLA33"), "[BLA33] is Variable")
}

func TestGetStringReplaceValuesFunc(t *testing.T) {
	var str = "Toko Amiko BLA BL"
	var values = map[string]string{"Toko": "...", "BLA": "BLU", "BL":"LOL"}
	stringReplaceFunc := getStringReplaceValuesFunc(values)

	actual := stringReplaceFunc(str)

	AssertEqualsString(t, "... Amiko BLU LOL", actual)

	//-------------------------------------------------------------//

	str = "@Toko Amiko BLA BL"
	values = map[string]string{"Toko": "...", "BLA": "BLU", "BL":"LOL"}
	stringReplaceFunc = getStringReplaceValuesFunc(values)

	actual = stringReplaceFunc(str)

	AssertEqualsString(t, "@... Amiko BLU LOL", actual)

	//-------------------------------------------------------------//

	str = "@Toko"
	values = map[string]string{"Toko": "...", "BLA": "BLU", "BL":"LOL"}
	stringReplaceFunc = getStringReplaceValuesFunc(values)

	actual = stringReplaceFunc(str)

	AssertEqualsString(t, "@...", actual)
}

func TestParse(t *testing.T) {
	var str = "" +
		"// This file is part of www.nand2tetris.org\n" +
		"// and the book \"The Elements of Computing Systems\"\n" +
		"// by Nisan and Schocken, MIT Press.\n" +
		"       // File name: projects/06/add/Add.asm\n" +
		"  \n" +
		"  // Computes R0 = 2 + 3  (R0 refers to RAM[0])\n" +
		"\n" +
		"@2 //ბლა\n" +
		"D=A\n" +
		"  @3  \n" +
		"    D=D+A\n" +
		"@temp\n" +
		"M=D\n"

	var expectedLines = []string{
		"@2",
		"D=A",
		"@3",
		"D=D+A",
		"@16",
		"M=D",
	}

	result := getParser()
	result.Parse(str)

	AssertEqualsInt(t, len(result.AssemblyLines), len(expectedLines))
	for i := range expectedLines {
		expected := expectedLines[i]
		actual := result.AssemblyLines[i]
		AssertEqualsString(t, expected, actual)
	}
}
