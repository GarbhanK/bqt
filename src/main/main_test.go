package main

import (
	"fmt"
	"testing"

	"golang.design/x/clipboard"
)

func TestExportToClipboard(t *testing.T) {
	err := clipboard.Init()
	if err != nil {
		fmt.Printf("Unable to init clipboard, %s", err.Error())
	}

	var input_string string = "select * from `table` where query_type = 'test'"
	ExportToClipboard(input_string)

	output := string(clipboard.Read(clipboard.FmtText)) 

	if (len(output) <= 0) {
		t.Error("Clipboard contents is either non-existent or empty\n")
	} 
	
	if input_string == output {
		t.Log("ExportToClipboard PASSED. Input string is equal to string read from clipboard\n")
	} else {
		t.Error("ExportToClipboard FAILED. Mismatching input string and clipboard output\n")
	}
}
