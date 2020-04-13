package main

type HackAssembler struct {
}

type TranslatorI interface {
	Translate(parser *ParseResult) string
}

var translator TranslatorI = GetTranslator()

func GetHackAssembler() *HackAssembler {
	return &HackAssembler{}
}

func (a *HackAssembler) Assemble(assembly string) string {
	var parser = getParser()
	parser.Parse(assembly)
	result := translator.Translate(parser)
	return result
}