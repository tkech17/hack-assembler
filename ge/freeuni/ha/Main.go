package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	defaultDirectory = "ge/freeuni/ha/resources"
	defaultFile      = "tmp.txt"
)

type Assembler interface {
	Assemble(assembly string) string
}

var assembler Assembler = GetHackAssembler()

func main() {
	var content = getContentFromFile()
	var parsed = assembler.Assemble(content)
	fmt.Println(parsed)
}

func getContentFromFile() string {
	var filePath = getFilePath()
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("File reading error", err)
	}
	return string(data)
}

func getFilePath() string {
	arguments := os.Args[1:]
	var folder = getFolder(arguments)
	var fileName = getFileName(arguments)
	return folder + "/" + fileName
}

func getFolder(arguments []string) string {
	for _, argument := range arguments {
		args := strings.Split(argument, "=")
		if "--folder" == args[0] {
			return args[1]
		}
	}
	return defaultDirectory
}

func getFileName(arguments []string) string {
	for _, argument := range arguments {
		args := strings.Split(argument, "=")
		if "--file" == args[0] {
			return args[1]
		}
	}
	return defaultFile
}
