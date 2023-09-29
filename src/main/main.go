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
		mappingFilePath = "./mappings.json"
	} else {
		mappingFilePath = fmt.Sprintf("%s/Documents/bqt/mappings.json", homeDir)
	}
	mappingFile, err := os.ReadFile(mappingFilePath)
	if err != nil {
		fmt.Printf("Cannot find 'mappings.json' file in path %s, %s\n", mappingFilePath, err.Error())
		os.Exit(0)
	}
	m := map[string]string{}

	// read json file into a map[string]string
	json.Unmarshal([]byte(mappingFile), &m)

	for k, v := range m {
		m[k] = strings.Replace(v, "${env}", env, 1)
	}

	// grab the current airflow date (today -1)
	dt := time.Now().AddDate(0, 0, -1)
	ds := fmt.Sprintf("%d-%02d-%02d", dt.Year(), dt.Month(), dt.Day())
	ds_nodash := fmt.Sprintf("%02d%02d%02d", dt.Year(), dt.Month(), dt.Day())
	ts := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d+00:00",
		dt.Year(), dt.Month(), dt.Day(),
		dt.Hour(), dt.Minute(), dt.Second())
	yesterday_ds := fmt.Sprint(dt.AddDate(0, 0, -1))
	tomorrow_ds := fmt.Sprint(dt.AddDate(0, 0, 1))

	m["ds_nodash"] = ds_nodash
	m["ds"] = ds
	m["ts"] = ts
	m["yesterday_ds"] = yesterday_ds
	m["tomorrow_ds"] = tomorrow_ds

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

func TemplateSQLFile(fileName string, isTerraform bool, mapping map[string]string) string {
	sqlFile := ReadSQL(fileName)
	sqlFilePointer := &sqlFile

	var template string
	for k, v := range mapping {
		if isTerraform {
			template = fmt.Sprintf("${%s}", k)
		} else {
			template = fmt.Sprintf("{{ %s }}", k)
		}
		*sqlFilePointer = strings.ReplaceAll(*sqlFilePointer, template, v)
	}

	return sqlFile
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("First argument must be a sql file")
		os.Exit(0)
	}

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

	// read in sql file
	fileName := args[0]
	templatedSQL := TemplateSQLFile(fileName, isTerraform, m)

	// Send the templated string to the clipboard (doesn't work on linux)
	ExportToClipboard(templatedSQL)
	curr_clipboard := clipboard.Read(clipboard.FmtText)

	// don't print the output if quiet flag is provided
	if quiet {
		os.Exit(0)
	} else {
		fmt.Println(string(curr_clipboard))
	}

}
