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
// TODO: Increase the text coverage. Test for negative cases as well.

func commonUtility(input string, wait string) string {
	shortenedUrl := service.CreateShortenedUrl(input)

	if wait == "yup" {
		// approx 3 min, as entries will be removed if
		// there are in the db for more than 3 min
		// will also be removed from after after 3 min
		time.Sleep(302 * time.Second)
		service.RemovedExpiredEntries()
	}

	origUrl := service.UrlRedirection(shortenedUrl)
	return origUrl
}

func TestUrlcreation(t *testing.T) {

	// 1. test for a empty string
	testUrl := ""
	outputUrl := commonUtility(testUrl, "")

	if testUrl != outputUrl {
		t.Errorf("Test UrlCreation failed. Expected %s, got %s", testUrl, outputUrl)
	} else {
		t.Logf("Success, Expected %s, got %s", testUrl, outputUrl)
	}

	// 2. test for a real string
	testUrl = "https://rahulverma.me"
	outputUrl = commonUtility(testUrl, "")
	if testUrl != outputUrl {
		t.Errorf("Test UrlCreation failed. Expected %s, got %s", testUrl, outputUrl)
	} else {
		t.Logf("Success, Expected %s, got %s", testUrl, outputUrl)
	}

}

// check whether db cleaning service is working correctly or not.
// func TestDbpurging(t *testing.T) {
// 	testUrl := "https://google.com"
// 	outputUrl := commonUtility(testUrl, "yup")

// 	if outputUrl != "" {
// 		t.Errorf("Test DbPurging failed. Expected %s, got %s", "", outputUrl)
// 	} else {
// 		t.Logf("Success, Expected %s, got %s", "", outputUrl)
// 	}
// }
