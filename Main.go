package main

import (
	"fmt"
	"os"
	"strings"
)

const defaultDirectory = "files"
const defaultFile = "tmp.txt"

func main() {
	var filePath = getFilePath()
	fmt.Println(filePath)
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
