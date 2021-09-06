package gribtest

import (
	"log"
	"math"
	"os"
	"testing"

	"github.com/nilsmagnus/grib/griblib"
)

func Test_read_integrationtest_file(t *testing.T) {
	testFile, fileOpenErr := os.Open("../integrationtestdata/gfs.t18z.pgrb2.1p00.f003")

	if fileOpenErr != nil {
		t.Fatal("Grib file for integration tests not found")
	}
	messages, messageParseErr := griblib.ReadMessages(testFile)

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
				if kelvin < 197 || kelvin > 350 {
					t.Errorf("Got kelvin out of range: %f\n", kelvin)
				}
				if kelvin > max {
					max = kelvin
				}
				if kelvin < min {
					min = kelvin
				}
			}
			log.Printf("category number %v,", m.Section4.ProductDefinitionTemplate.ParameterCategory)
			log.Printf("parameter number %v,", m.Section4.ProductDefinitionTemplate.ParameterNumber)
			log.Printf("surface type %v, surface value %v max: %f min: %f\n", surface.Type, surface.Value, max, min)

		}
	}

}

func Test_read_integrationtest_file_hour0(t *testing.T) {
	testFile, fileOpenErr := os.Open("../integrationtestdata/gfs.t06z.pgrb2.1p00.f000")

	if fileOpenErr != nil {
		t.Fatal("Grib file for integration tests not found")
	}
	messages, messageParseErr := griblib.ReadMessages(testFile)

	if messageParseErr != nil {
		t.Fatal("Error reading messages: ", messageParseErr.Error())
	}

	if len(messages) != 354 {
		t.Errorf("should have exactly 354 message in testfile, was %d", len(messages))
	}

}

func Test_read3_integrationtest_file_hour0(t *testing.T) {

	testFile, gribFileOpenErr := os.Open("../integrationtestdata/template5_3.grib2")
	if gribFileOpenErr != nil {
		t.Fatalf("Grib file for integration tests not found %s", gribFileOpenErr.Error())
	}
	defer testFile.Close()

	resultFile, csvFileOpenErr := os.Open("../integrationtestdata/template_ugrd.csv")
	if gribFileOpenErr != nil {
		t.Fatalf("CSV file for integration tests not found %s", csvFileOpenErr.Error())
	}
	defer resultFile.Close()

	fixtures, errFixtures := readCsvAsSlice(resultFile)
	if errFixtures != nil {
		t.Fatalf("Could not parse CSV file %s", errFixtures.Error())
	}

	messages, messageParseErr := griblib.ReadMessages(testFile)

	if messageParseErr != nil {
		t.Fatal("Error reading messages: ", messageParseErr.Error())
	}

	if len(messages) != 2 {
		t.Errorf("should have exactly 2 messages in testfile, was %d", len(messages))
	}

	if messages[0].Section5.DataTemplateNumber != 3 {
		t.Errorf("Data template number should be 3 (found %d)", messages[0].Section5.DataTemplateNumber)
	}

	if len(fixtures) != len(messages[0].Data()) {
		t.Errorf("should have exactly 2 message in testfile, was %d", len(fixtures))
	}

	for index, data := range fixtures {
		if math.Ceil(messages[0].Section7.Data[index]*10000+.5) != math.Ceil(data*10000+.5) {
			t.Errorf("Expected value %f at index %d, found %f", data, index, messages[0].Section7.Data[index])
		}
	}

}
