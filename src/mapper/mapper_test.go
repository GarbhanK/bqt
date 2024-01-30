package mapper

import (
	"fmt"
	"testing"
)

func TestCreateMappingLive(t *testing.T) {

	expected_live := map[string]string{
		"params.project":     "gk-data-eu-live",
		"params.web_project": "my-webproject",
		"environment":        "live",
	}

	result_live := CreateMapping("live", true)

	eq_live := fmt.Sprint(result_live) == fmt.Sprint(expected_live)
	if eq_live {
		t.Logf("CreateMapping('live') PASSED. Expected %s\n, got %s\n", expected_live, result_live)
	} else {
		t.Errorf("CreateMapping('live') FAILED. Expected %s\n, got %s\n", expected_live, result_live)
	}

}

func TestCreateMappingDev(t *testing.T) {

	expected_dev := map[string]string{
		"params.project":     "gk-data-eu-dev",
		"params.web_project": "my-webproject",
		"environment":        "dev",
	}

	result_dev := CreateMapping("dev", true)

	eq_dev := fmt.Sprint(result_dev) == fmt.Sprint(expected_dev)
	if eq_dev {
		t.Logf("CreateMapping('dev') PASSED. Expected %s\n, got %s\n", expected_dev, result_dev)
	} else {
		t.Errorf("CreateMapping('dev') FAILED. Expected %s\n, got %s\n", expected_dev, result_dev)
	}
}

// func TestAddTemplateAirflowVars(t *testing.T) {

// 	sample_mapping := map[string]string{
// 		"params.project":     "gk-africa-data-eu-dev",
// 		"params.web_project": "testscore-web",
// 		"environment":        "dev",
// 	}
	
// 	dt := time.Now().AddDate(0, 0, -1)
// 	var current_timestamp string = fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d+00:00",
// 		dt.Year(), dt.Month(), dt.Day(),
// 		dt.Hour(), dt.Minute(), dt.Second())

// 	expected_output := map[string]string{
// 		"params.project": "gk-africa-data-eu-dev",
// 		"params.web_project": "testscore-web",
// 		"environment": "dev",
// 		"ds": "2023-11-08",
// 		"ds_nodash": "20231108",
// 		"ts": current_timestamp,
// 		"yesterday_ds": "2023-11-07",
// 		"yesterday_ds_nodash": "20231107",
// 		"tomorrow_ds": "2023-11-09",
// 		"tomorrow_ds_nodash": "20231109",
// 	}	

// 	func_result := AddAirflowTemplateVars(sample_mapping)

// 	if len(func_result) != len(expected_output) {
// 		t.Error("AddAirflowTemplateVars FAILED. Unequal amount of map entries\n")
//    }

//    // check if keys and matching values are present in both maps
//    for key, value := range expected_output {
//       if res_value, ok := func_result[key]; !ok || res_value != value {
// 		t.Errorf("AddAirflowTemplateVars FAILED. Mismatched values for %s and %s\n", value, res_value)
//       }
//    }

// 	t.Log("AddAirflowTemplateVars PASSED.\n")
// }
