package main

import (
	"fmt"
	"os"

	"github.com/garbhank/bqt/src/mapper"
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

type options struct {
	isTerraform bool
	quiet bool
	isTest bool
	env string
}

func newOptions(env string) *options {
	
	e := options{}
	e.env = "dev"	// defaults to dev for safety
	return &e
}


func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("First argument must be a sql file")
		os.Exit(0)
	}

	opts := newOptions()

	for i := range args {
		switch args[i] {
		case "tf":
			opts.isTerraform = true
		case "live":
			opts.env = "live"
		case "dev":
			opts.env = "dev"
		case "staging":
			opts.env = "staging"
		case "quiet":
			opts.quiet = true
		case "testMap":
			opts.isTest = true
		}
	}

	// template/value mapping from 'mapping.json'
	m := mapper.CreateMapping(opts.env, opts.isTest)
	m = mapper.AddAirflowTemplateVars(m)

	// read in sql file
	fileName := args[0]
	templatedSQL := templater.TemplateSQLFile(fileName, opts.isTerraform, m)

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
