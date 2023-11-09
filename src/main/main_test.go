package main

import (
	"testing"

	"golang.design/x/clipboard"
)

func TestExportToClipboard(t *testing.T) {

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
