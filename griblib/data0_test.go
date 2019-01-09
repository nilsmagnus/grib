package griblib

import (
	"os"
	"testing"
)

func Test_read0_integrationtest_file_hour0(t *testing.T) {
	testFile, fileOpenErr := os.Open("integrationtestdata/template5_0.grib2")

	if fileOpenErr != nil {
		t.Fatalf("Grib file for integration tests not found %s", fileOpenErr.Error())
	}
	messages, messageParseErr := ReadMessages(testFile)

	if messageParseErr != nil {
		t.Fatal("Error reading messages: ", messageParseErr.Error())
	}

	if len(messages) != 2 {
		t.Errorf("should have exactly 2 message in testfile, was %d", len(messages))
	}

	if messages[0].Section5.DataTemplateNumber != 0 {
		t.Errorf("Data template number should be 0 (found %d)", messages[0].Section5.DataTemplateNumber)
	}

}
