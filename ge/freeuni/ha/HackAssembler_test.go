package main

import "testing"

func TestGetHackAssembler(t *testing.T) {
	assembler := GetHackAssembler()
	AssertNotNull(t, assembler)
}