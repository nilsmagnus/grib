package griblib

import (
	"os"
	"testing"
	"fmt"
)

func Test_read_integrationtest_file(t *testing.T) {
	testFile, fileOpenErr := os.Open("integrationtestdata/gfs.t00z.pgrb2.2p50.f012")

	if fileOpenErr != nil {
		t.Fatal("Grib file for integration tests not found")
	}
	messages, messageParseErr := ReadMessages(testFile)

	if messageParseErr != nil {
		t.Fatal("Error reading messages: ", messageParseErr.Error())
	}

	if len(messages) != 77 {
		t.Error("should have exactly 77 message in testfile")
	}

	for _, m := range messages {
//		if m.Section4.ProductDefinitionTemplate.ParameterCategory == 0 {
			fmt.Printf("category %d   \tparameter number %d ",
				m.Section4.ProductDefinitionTemplate.ParameterCategory,
				m.Section4.ProductDefinitionTemplate.ParameterNumber)
			fmt.Printf("surface: type %d\t value %d\n",
				m.Section4.ProductDefinitionTemplate.FirstSurface.Type,
				m.Section4.ProductDefinitionTemplate.FirstSurface.Value)
//		}
	}

}

func Test_read_integrationtest_file_hour0(t *testing.T) {
	testFile, fileOpenErr := os.Open("integrationtestdata/gfs.t06z.pgrb2.1p00.f000")

	if fileOpenErr != nil {
		t.Fatal("Grib file for integration tests not found")
	}
	messages, messageParseErr := ReadMessages(testFile)

	if messageParseErr != nil {
		t.Fatal("Error reading messages: ", messageParseErr.Error())
	}

	if len(messages) != 354 {
		t.Errorf("should have exactly 354 message in testfile, was %d", len(messages))
	}

}
