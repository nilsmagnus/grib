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

	if len(messages) != 366 {
		t.Errorf("should have exactly 366 message in testfile, was %d", len(messages))
	}

	if messages[0].Section5.DataTemplateNumber != 3 {
		t.Errorf("Data template number should be 3 (found %d)", messages[0].Section5.DataTemplateNumber)
	}

	for _, m := range messages {
		surface := m.Section4.ProductDefinitionTemplate.FirstSurface
		if surface.Type == 1 && // ground surface
			m.Section0.Discipline == 0 && // meterology
			m.Section4.ProductDefinitionTemplate.ParameterCategory == 0 { // temperature
			var max float64 = 00
			var min float64 = 1000
			for _, kelvin := range m.Section7.Data {
				if kelvin < 200 || kelvin > 350 {
					t.Errorf("Got kelvin out of range: %f\n", kelvin)
				}
				if kelvin > max {
					max = kelvin
				}
				if kelvin < min {
					min = kelvin
				}
			}
			fmt.Printf("category number %v,", m.Section4.ProductDefinitionTemplate.ParameterCategory)
			fmt.Printf("parameter number %v,", m.Section4.ProductDefinitionTemplate.ParameterNumber)
			fmt.Printf("surface type %v, surface value %v max: %f min: %f\n", surface.Type, surface.Value, max, min)

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
