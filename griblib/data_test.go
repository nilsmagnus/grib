package griblib

import (
	"fmt"
	"os"
	"testing"
)

func Test_read_integrationtest_file(t *testing.T) {
	testFile, fileOpenErr := os.Open("integrationtestdata/gfs.t18z.pgrb2.1p00.f003")

	if fileOpenErr != nil {
		t.Fatal("Grib file for integration tests not found")
	}
	messages, messageParseErr := ReadMessages(testFile)

	if messageParseErr != nil {
		t.Fatal("Error reading messages: ", messageParseErr.Error())
	}

	if len(messages) != 78 {
		t.Errorf("should have exactly 78 message in testfile, was %d", len(messages))
	}

	for _, m := range messages {
		// isobaric temperatures at level 100Pa
		if m.Section4.ProductDefinitionTemplate.ParameterCategory == 0 &&
			m.Section4.ProductDefinitionTemplate.FirstSurface.Type == 100 &&
			m.Section4.ProductDefinitionTemplate.FirstSurface.Value == 100 {
			var max float64 = 00
			var min float64 = 1000
			for _, kelvin := range m.Section7.Data {
				if kelvin < 227 || kelvin > 281.2 {
					t.Errorf("Got kelvin out of range: %f\n", kelvin)
				}
				if kelvin > max {
					max = kelvin
				}
				if kelvin < min {
					min = kelvin
				}
			}
			fmt.Printf("max: %f min: %f\n", max, min)

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
