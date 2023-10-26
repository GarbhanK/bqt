package mapper

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func CreateMapping(env string, isTest bool) map[string]string {

	var mappingFilePath string

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	if isTest {
		mappingFilePath = "./test_mappings.json"
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

	return m
}

func AddAirflowTemplateVars(m map[string]string) map[string]string {
	// grab the current airflow date (today -1)
	dt := time.Now().AddDate(0, 0, -1)

	// create airflow template variables ref: https://airflow.apache.org/docs/apache-airflow/stable/templates-ref.html
	ds := fmt.Sprintf("%d-%02d-%02d", dt.Year(), dt.Month(), dt.Day())
	ds_nodash := fmt.Sprintf("%02d%02d%02d", dt.Year(), dt.Month(), dt.Day())
	ts := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d+00:00",
		dt.Year(), dt.Month(), dt.Day(),
		dt.Hour(), dt.Minute(), dt.Second())
	yesterday_ds := fmt.Sprint(dt.AddDate(0, 0, -1))
	tomorrow_ds := fmt.Sprint(dt.AddDate(0, 0, 1))

	m["ds"] = ds
	m["ds_nodash"] = ds_nodash
	m["ts"] = ts
	m["yesterday_ds"] = yesterday_ds
	m["tomorrow_ds"] = tomorrow_ds

	return m
}
