package main

type HackAssembler struct {

}

type Parser interface {
	Parse(content string)
}

var parser Parser = getParser()

func GetHackAssembler() *HackAssembler {
	return &HackAssembler{}
}

func (a *HackAssembler) Assemble(assembly string) string  {
	parser.Parse(assembly)
	return ""
}