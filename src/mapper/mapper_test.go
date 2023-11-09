package mapper

import (
	"fmt"
	"testing"
)

func TestCreateMappingLive(t *testing.T) {

	expected_live := map[string]string{
		"params.project":     "gk-africa-data-eu-live",
		"params.web_project": "testscore-web",
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
		"params.project":     "gk-africa-data-eu-dev",
		"params.web_project": "testscore-web",
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
