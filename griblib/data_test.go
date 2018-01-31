package griblib

import (
	"os"
	"testing"
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
		// isobaric temperatures at level 100Pa
		if m.Section4.ProductDefinitionTemplate.ParameterCategory == 0 &&
			m.Section4.ProductDefinitionTemplate.FirstSurface.Type == 100 &&
			m.Section4.ProductDefinitionTemplate.FirstSurface.Value == 100 {
			for _, kelvin := range m.Section7.Data {
				if kelvin < 237.5 || kelvin > 269.9 {
					t.Errorf("Got kelvin out of range: %f\n", kelvin)
				}
			}

		}
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
