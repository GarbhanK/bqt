package templater

import (
	"fmt"
	"testing"
)

func TestReadSQL(t *testing.T) {

	result := ReadSQL("test.sql")

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
	result := ReadSQL("terraform_test.sql")

	sql := "select * from `${params.project}.transactions.coffee` c\n"
	sql += "where date(insertionTimestamp) >= '${ds_nodash}'\n"
	sql += "left join `${params.web_project}.unified_segment.tracks` t\n"
	sql += "on c.userId = t.userId\n"
	sql += "group by insertionTimestamp desc"
	expected := sql

	if len(result) > 0 {
		t.Logf("ReadSQL('terraform.sql') PASSED. Is not an empty string\n")
	} else {
		t.Errorf("ReadSQL('terraform_test.sql') FAILED. Got an empty string\n")
	}

	if result != expected {
		t.Errorf("ReadSQL('test.sql') FAILED. Expected %s, got %s\n", expected, result)
	} else {
		t.Logf("ReadSQL('test.sql') PASSED. Expected %s, got %s\n", expected, result)
	}
}
