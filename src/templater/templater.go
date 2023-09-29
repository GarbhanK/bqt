package templater

import (
	"fmt"
	"os"
	"strings"
)

func ReadSQL(fileName string) string {
	f, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fileString := string(f)
	return fileString
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
