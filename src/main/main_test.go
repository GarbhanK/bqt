package main

import (
	"reflect"
	"testing"
)

// go test -v tests/main_test.go

func TestReadSQL(t *testing.T) {

	result := ReadSQL("test.sql")
	var expected string

	expected = "select * from from `{{ params.project }}.transactions.coffee`" 
	+ "c where date(insertionTimestamp) >= '{{ ds_nodash }}'"
	+ "left join `{{ params.web_project }}.unified_segment.tracks` t"
	+ "on c.userId = t.userId group by insertionTimestamp desc"

	if result != expected {
		t.Errorf("ReadSQL('test.sql') FAILED. Expected %s, got %s\n", expected, result)
	} else {
		t.Logf("ReadSQL('test.sql') PASSED. Expected %s, got %s\n", expected, result)
	}
}

func TestCreateMapping(t *testing.T) {

	result_live := CreateMapping("live")
	expected_live := make(map[string]string)

	result_dev := CreateMapping("dev")
	expected_dev := make(map[string]string)

	result_staging := CreateMapping("staging")
	expected_staging := make(map[string]string)

	// could just need fmt.Sprint(map1) == fmt.Spring(map2)

	eq_live := reflect.DeepEqual(result_live, expected_live)
	if eq_live {
		t.Logf("CreateMapping('live') PASSED. Expected %s, got %s\n", expected_live, result_live)
	} else {
		t.Errorf("CreateMapping('live') FAILED. Expected %s, got %s\n", expected_live, result_live)
	}

	eq_dev := reflect.DeepEqual(result_dev, expected_dev)
	if eq_dev {
		t.Logf("CreateMapping('dev') PASSED. Expected %s, got %s\n", expected_dev, result_dev)
	} else {
		t.Errorf("CreateMapping('dev') FAILED. Expected %s, got %s\n", expected_dev, result_dev)
	
	eq_live := reflect.DeepEqual(result_live, expected_live)
	if eq_live {
		t.Logf("CreateMapping('live') PASSED. Expected %s, got %s\n", expected_live, result_live)
	} else {
		t.Errorf("CreateMapping('live') FAILED. Expected %s, got %s\n", expected_live, result_live)
}
