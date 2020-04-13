package main

import (
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

const testFilesDir = "resources/"

type AssemblerTest interface {
	Assemble(assembly string) string
}

func TestAllFiles(t *testing.T) {
	files := getFilesKeyValues()
	for key, value := range files {
		contentB, _ := ioutil.ReadFile(testFilesDir + key)
		hackB, _ := ioutil.ReadFile(testFilesDir + value)
		expected := strings.TrimSpace(string(hackB))
		var assembler AssemblerTest = GetHackAssembler()
		actual := assembler.Assemble(string(contentB))
		AssertEqualsString(t, expected, actual)
	}
}

func getFilesKeyValues() map[string]string {
	files, err := ioutil.ReadDir(testFilesDir)
	if err != nil {
		log.Fatal(err)
	}

	var result []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".asm") {
			result = append(result, file.Name())
		}
	}
	var mp = make(map[string]string)
	for _, file := range result {
		value := file[0:len(file)-4] + ".hack"
		mp[file] = value
	}
	return mp
}
