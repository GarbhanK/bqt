package main

import (
	"fmt"
	"os"
	"strings"

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
	dat, err := os.ReadFile("wow.sql")
	check(err)
	fmt.Print(string(dat))

	// read from json config to map template value to string

	// text replace the template vars "{{ x }}"
	strings.ReplaceAll(string(dat), "{{ params.project }}", "gk-africa-data-eu-dev")

	// return output to clipboard
	err = clipboard.Init()
	if err != nil {
		fmt.Printf("Unable to init clipboard, %s", err.Error())
	}

	clipboard.Write(clipboard.FmtText, []byte("text data"))

	curr_clipboard := clipboard.Read(clipboard.FmtText)
	fmt.Println(string(curr_clipboard))

}
