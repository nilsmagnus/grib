package gribtest

import (
	"math"
	"os"
	"testing"

	"github.com/nilsmagnus/grib/griblib"
)

func Test_read2_integrationtest_file_hour0(t *testing.T) {

	testFile, gribFileOpenErr := os.Open("../integrationtestdata/template5_2.grib2")
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

	if messages[0].Section5.DataTemplateNumber != 2 {
		t.Errorf("Data template number should be 2 (found %d)", messages[0].Section5.DataTemplateNumber)
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
