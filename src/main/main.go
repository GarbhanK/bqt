package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"golang.design/x/clipboard"
)

func ReadSQL(fileName string) string {
	f, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fileString := string(f)
	return fileString
}

func CreateMapping(env string, isTest bool) map[string]string {

	var mappingFilePath string

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	if isTest {
		mappingFilePath = fmt.Sprintf("mappings.json")
	} else {
		mappingFilePath = fmt.Sprintf("%s/Documents/bqt/mappings.json", homeDir)
	}
	mappingFile, err := os.ReadFile(mappingFilePath)
	if err != nil {
		fmt.Printf("Cannot find 'mappings.json' file, %s\n", err.Error())
		os.Exit(0)
	}
	m := map[string]string{}

	// read json file into a map[string]string
	json.Unmarshal([]byte(mappingFile), &m)

	for k, v := range m {
		m[k] = strings.Replace(v, "${env}", env, 1)
	}

	// make ds_nodash current_date
	dt := time.Now()
	ds_nodash := fmt.Sprintf("%d-%d-%d", dt.Year(), dt.Month(), dt.Day())
	ds := fmt.Sprintf("%d%d%d", dt.Year(), dt.Month(), dt.Day())
	m["ds_nodash"] = ds_nodash
	m["ds"] = ds

	return m
}

func ExportToClipboard(templatedStr string) {
	err := clipboard.Init()
	if err != nil {
		fmt.Printf("Unable to init clipboard, %s", err.Error())
	}

	byteSql := []byte(templatedStr)
	clipboard.Write(clipboard.FmtText, byteSql)
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("First argument must be a sql file")
		os.Exit(0)
	}

	// read in sql file
	fileName := args[0]
	sqlFile := ReadSQL(fileName)

	var isTerraform bool
	var quiet bool
	var isTest bool

	// defaults to dev for safety
	env := "dev"
	for i, _ := range args {
		switch args[i] {
		case "tf":
			isTerraform = true
		case "live":
			env = "live"
		case "dev":
			env = "dev"
		case "staging":
			env = "staging"
		case "quiet":
			quiet = true
		case "testMap":
			isTest = true
		}
	}

	// template/value mapping from 'mapping.json'
	m := CreateMapping(env, isTest)

	var formattedString string
	tempFile := strings.Clone(sqlFile)

	var template string
	for k, v := range m {
		if isTerraform {
			template = fmt.Sprintf("${%s}", k)
		} else {
			template = fmt.Sprintf("{{ %s }}", k)
		}
		formattedString = strings.ReplaceAll(tempFile, template, v)
		tempFile = strings.Clone(formattedString)
	}

	// Send the templated string to the clipboard (doesn't work on linux)
	ExportToClipboard(formattedString)
	curr_clipboard := clipboard.Read(clipboard.FmtText)

	// don't print the output if quiet flag is provided
	if quiet {
		os.Exit(0)
	} else {
		fmt.Println(string(curr_clipboard))
	}

}
