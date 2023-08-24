package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"golang.design/x/clipboard"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Hello bqt!")

	// read in sql file
	f, err := os.ReadFile("wow.sql")
	check(err)
	// fmt.Print(string(f))
	sqlFile := string(f)

	// read from json config to map template value to string
	mappingFile, err := os.ReadFile("mappings.json")
	if err != nil {
		panic(err)
	}
	m := map[string]string{}

	json.Unmarshal([]byte(mappingFile), &m)
	fmt.Println(m)

	// make ds_noj=dash current_date
	dt := time.Now()
	ds_nodash := fmt.Sprintf("%d-%d-%d", dt.Year(), dt.Month(), dt.Day())
	m["ds_nodash"] = ds_nodash

	var formatString string
	tempFile := strings.Clone(sqlFile)
	for k, v := range m {
		template := fmt.Sprintf("{{ %s }}", k)
		formatString = strings.ReplaceAll(tempFile, template, v)
		tempFile = strings.Clone(formatString)
	}
	fmt.Println(formatString)

	// text replace the template vars "{{ x }}"
	// fmtdString := strings.ReplaceAll(string(sqlFile), "{{ params.project }}", "gk-africa-data-eu-dev")
	// fmt.Print(fmtdString)

	// return output to clipboard
	err = clipboard.Init()
	if err != nil {
		fmt.Printf("Unable to init clipboard, %s", err.Error())
	}

	clipboard.Write(clipboard.FmtText, []byte("text data"))

	curr_clipboard := clipboard.Read(clipboard.FmtText)
	fmt.Println(string(curr_clipboard))

}
