package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"golang.design/x/clipboard"
)

func readSQL(fileName string) string {
	f, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fileString := string(f)
	return fileString
}

func createMapping(env string) map[string]string {
	mappingFile, err := os.ReadFile("mappings.json")
	if err != nil {
		fmt.Printf("Cannot find 'mappings.json' file, %s", err.Error())
		os.Exit(0)
	}
	m := map[string]string{}

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

func exportToClipboard(templatedStr string) {
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
		fmt.Println("Please an input SQL file")
		os.Exit(0)
	}

	// read in sql file
	fileName := args[0]
	sqlFile := readSQL(fileName)

	var isTerraform bool
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
		}
	}

	// template/value mapping from 'mapping.json'
	m := createMapping(env)

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
	exportToClipboard(formattedString)
	curr_clipboard := clipboard.Read(clipboard.FmtText)
	fmt.Print(string(curr_clipboard))

}
