package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	defaultSourceDirectory      = "resources"
	defaultFile                 = "Max.asm"
	defaultDestinationDirectory = "results"
)

type Assembler interface {
	Assemble(assembly string) string
}

var assembler Assembler = GetHackAssembler()

func main() {
	var content = getContentFromFile()
	var parsed = assembler.Assemble(content)
	saveIntoDestinationFile(parsed)
	fmt.Println(parsed)
}

func saveIntoDestinationFile(content string) {
	destFile := gedDestinationFileName()
	if strings.Contains(destFile, "/") {
		err := os.MkdirAll(destFile[0:strings.LastIndex(destFile, "/")], os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(destFile)
	wd, _ := os.Getwd()
	file, err := os.Create(wd + "/" + destFile)
	defer closeFile(file)
	if err == nil {
		_, writeErr := file.WriteString(content)
		if writeErr != nil {
			log.Fatal(writeErr)
		}
	} else {
		log.Fatal(err)
	}
}

//noinspection GoUnhandledErrorResult
func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func gedDestinationFileName() string {
	arguments := os.Args[1:]
	var destFolder = getDestinationFolder(arguments)
	var festFile = getDestinationFileName(arguments)
	return destFolder + "/" + festFile
}

func getDestinationFileName(arguments []string) string {
	var filename string = defaultFile
	for _, argument := range arguments {
		args := strings.Split(argument, "=")
		if "--source-file" == args[0] {
			filename = args[1]
		}
	}
	filename = filename[0:len(filename)-4] + ".hack"
	return filename
}

func getDestinationFolder(arguments []string) string {
	for _, argument := range arguments {
		args := strings.Split(argument, "=")
		if "--dest-folder" == args[0] {
			return args[1]
		}
	}
	return defaultDestinationDirectory
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
	var folder = getSourceFolder(arguments)
	var fileName = getSourceFileName(arguments)
	return folder + "/" + fileName
}

func getSourceFolder(arguments []string) string {
	for _, argument := range arguments {
		args := strings.Split(argument, "=")
		if "--source-folder" == args[0] {
			return args[1]
		}
	}
	return defaultSourceDirectory
}

func getSourceFileName(arguments []string) string {
	for _, argument := range arguments {
		args := strings.Split(argument, "=")
		if "--source-file" == args[0] {
			return args[1]
		}
	}
	return defaultFile
}
