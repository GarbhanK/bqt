package main

import (
	"fmt"
	"os"

	"github.com/garbhank/bqt/src/parser"
	"github.com/garbhank/bqt/src/templater"
	"golang.design/x/clipboard"
)

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

	var isTerraform bool
	var quiet bool
	var isTest bool

	// defaults to dev for safety
	env := "dev"
	for i := range args {
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
	m := parser.CreateMapping(env, isTest)
	m = parser.AddAirflowTemplateVars(m)

	// read in sql file
	fileName := args[0]
	templatedSQL := templater.TemplateSQLFile(fileName, isTerraform, m)

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
