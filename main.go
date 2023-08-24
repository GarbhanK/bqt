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

func createMapping() map[string]string {
	mappingFile, err := os.ReadFile("mappings.json")
	if err != nil {
		panic(err)
	}
	m := map[string]string{}

	json.Unmarshal([]byte(mappingFile), &m)
	fmt.Println(m)

	// make ds_nodash current_date
	dt := time.Now()
	ds_nodash := fmt.Sprintf("%d-%d-%d", dt.Year(), dt.Month(), dt.Day())
	m["ds_nodash"] = ds_nodash

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
	fmt.Println("Hello bqt!")

	// read in sql file
	sqlFile := readSQL("wow.sql")

	// template/value mapping from 'mapping.json'
	m := createMapping()

	var formatString string
	tempFile := strings.Clone(sqlFile)

	for k, v := range m {
		// TODO: add branch for terraform/airflow
		template := fmt.Sprintf("{{ %s }}", k)
		formatString = strings.ReplaceAll(tempFile, template, v)
		tempFile = strings.Clone(formatString)
	}
	fmt.Println(formatString)

	exportToClipboard(formatString)

	curr_clipboard := clipboard.Read(clipboard.FmtText)
	fmt.Println(string(curr_clipboard))

	fmt.Println("byebye")
}
