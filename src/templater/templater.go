package templater

import (
	"fmt"
	"os"
	"strings"

	// "github.com/xwb1989/sqlparser"
	"github.com/fatih/color"
)

func ReadSQL(fileName string) string {
	f, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fileString := string(f)
	return fileString
}

// This mostly works but I can't translate BigQuery SQL to MySQL (date function)
// func validateSQL(sqlFile string) error {

// 	formattedSQL := strings.ReplaceAll(sqlFile, "\n", " ")
// 	formattedSQL = strings.ReplaceAll(formattedSQL, "`", "'")

// 	fmt.Printf(formattedSQL)

// 	stmt, err := sqlparser.Parse(formattedSQL)
// 	if err != nil {
// 		panic(err)
// 	}

// 	switch stmt := stmt.(type) {
// 	case *sqlparser.Select:
// 		_ = stmt
// 		return nil
// 	case *sqlparser.DBDDL:
// 		fmt.Printf("DBDDL type")
// 		return errors.New("Warning - your query is a CREATE, ALTER, DROP, RENAME or TRUNCATE statement")
// 	}

// 	return errors.New("Cannot validate query type...")
// }

func ValidateSQL(sqlFile string) {
	formattedSQL := strings.ReplaceAll(sqlFile, "\n", " ")
	queryWords := strings.Split(formattedSQL, " ")

	var statementType string = queryWords[0]
	if strings.ToLower(statementType) == "create" {
		color.Yellow("WARNING - copied query is a CREATE statement!\n")
	}
}

func TemplateSQLFile(fileName string, isTerraform bool, mapping map[string]string) string {
	sqlFile := ReadSQL(fileName)
	sqlFilePtr := &sqlFile

	var template string
	for k, v := range mapping {
		if isTerraform {
			template = fmt.Sprintf("${%s}", k)
		} else {
			template = fmt.Sprintf("{{ %s }}", k)
		}
		*sqlFilePtr = strings.ReplaceAll(*sqlFilePtr, template, v)
	}

	// err := validateSQL(sqlFile)
	// if err != nil {
	// 	fmt.Printf("Warning... your query is a CREATE statement...\n")
	// 	panic(err)
	// }

	return sqlFile
}
