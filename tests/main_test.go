package tests

import (
	service "goRubu/services"
	"testing"
	"time"
)

// NOTE - you need to save UT's in separate files ending with _test.
// Go provides "go test" command to execute these files and run tests.

// main test function. Checks both url creation and url redirection
// should start with "Test" and latter name should start with capital letter

func commonUtility(input string, wait string) string {
	shortened_url := service.CreateShortenedUrl(input)

	if wait == "yup" {
		// approx 3 min, as entries will be removed if
		// there are in the db for more than 3 min
		// will also be removed from after after 3 min
		time.Sleep(182 * time.Second)
		service.RemovedExpiredEntries()
	}

	orig_url := service.UrlRedirection(shortened_url)
	return orig_url
}

func TestUrlcreation(t *testing.T) {

	// 1. test for a empty string
	test_url := ""
	output_url := commonUtility(test_url, "")

	if test_url != output_url {
		t.Errorf("Test UrlCreation failed. Expected %s, got %s", test_url, output_url)
	} else {
		t.Logf("Success, Expected %s, got %s", test_url, output_url)
	}

	// 2. test for a real string
	test_url = "https://rahulverma.me"
	output_url = commonUtility(test_url, "")
	if test_url != output_url {
		t.Errorf("Test UrlCreation failed. Expected %s, got %s", test_url, output_url)
	} else {
		t.Logf("Success, Expected %s, got %s", test_url, output_url)
	}

}

// check whether db cleaning service is working correctly or not.
func TestDbpurging(t *testing.T) {
	test_url := "https://google.com"
	output_url := commonUtility(test_url, "yup")

	if output_url != "" {
		t.Errorf("Test DbPurging failed. Expected %s, got %s", "", output_url)
	} else {
		t.Logf("Success, Expected %s, got %s", "", output_url)
	}
}
