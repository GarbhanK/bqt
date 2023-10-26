package main

import (
	"fmt"
	"testing"

	"github.com/garbhank/bqt/src/mapper"
	"github.com/garbhank/bqt/src/templater"
)

func TestReadSQL(t *testing.T) {

	result := templater.ReadSQL("test.sql")

	sql := "select * from `{{ params.project }}.transactions.coffee` c\n"
	sql += "where date(insertionTimestamp) >= '{{ ds_nodash }}'\n"
	sql += "left join `{{ params.web_project }}.unified_segment.tracks` t\n"
	sql += "on c.userId = t.userId\n"
	sql += "group by insertionTimestamp desc"
	expected := sql

	if len(result) > 0 {
		t.Logf("ReadSQL('test.sql') PASSED. Is not an empty string\n")
	} else {
		t.Errorf("ReadSQL('test.sql') FAILED. Got an empty string\n")
	}

	if result != expected {
		t.Errorf("ReadSQL('test.sql') FAILED. Expected %s, got %s\n", expected, result)
	} else {
		t.Logf("ReadSQL('test.sql') PASSED. Expected %s, got %s\n", expected, result)
	}
}

func TestReadSQLTerraform(t *testing.T) {
	result := templater.ReadSQL("terraform_template.sql")

	sql := "select * from `${params.project}.transactions.coffee` c\n"
	sql += "where date(insertionTimestamp) >= '${ds_nodash}'\n"
	sql += "left join `${params.web_project}.unified_segment.tracks` t\n"
	sql += "on c.userId = t.userId\n"
	sql += "group by insertionTimestamp desc"
	expected := sql

	if len(result) > 0 {
		t.Logf("ReadSQL('test.sql') PASSED. Is not an empty string\n")
	} else {
		t.Errorf("ReadSQL('test.sql') FAILED. Got an empty string\n")
	}

	if result != expected {
		t.Errorf("ReadSQL('test.sql') FAILED. Expected %s, got %s\n", expected, result)
	} else {
		t.Logf("ReadSQL('test.sql') PASSED. Expected %s, got %s\n", expected, result)
	}
}

func TestCreateMapping(t *testing.T) {

	expected_live := map[string]string{
		"params.project":     "gk-africa-data-eu-live",
		"params.web_project": "livescore-web",
		"environment":        "live",
	}

	result_live := mapper.CreateMapping("live", true)

	expected_dev := map[string]string{
		"params.project":     "gk-africa-data-eu-dev",
		"params.web_project": "livescore-web",
		"environment":        "dev",
	}

	result_dev := mapper.CreateMapping("dev", true)

	eq_live := fmt.Sprint(result_live) == fmt.Sprint(expected_live)
	if eq_live {
		t.Logf("CreateMapping('live') PASSED. Expected %s\n, got %s\n", expected_live, result_live)
	} else {
		t.Errorf("CreateMapping('live') FAILED. Expected %s\n, got %s\n", expected_live, result_live)
	}

	eq_dev := fmt.Sprint(result_dev) == fmt.Sprint(expected_dev)
	if eq_dev {
		t.Logf("CreateMapping('dev') PASSED. Expected %s\n, got %s\n", expected_dev, result_dev)
	} else {
		t.Errorf("CreateMapping('dev') FAILED. Expected %s\n, got %s\n", expected_dev, result_dev)
	}
}
