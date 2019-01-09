package griblib

import (
	"math"
	"os"
	"testing"
)

func Test_read2_integrationtest_file_hour0(t *testing.T) {

	fixtures := []float64{
		1.18876,
		1.16876,
		1.13876,
		1.11876,
		1.09876,
		1.06876,
		1.04876,
		1.02876,
		0.998765,
		0.968765,
	}

	testFile, fileOpenErr := os.Open("integrationtestdata/template5_2.grib2")

	if fileOpenErr != nil {
		t.Fatalf("Grib file for integration tests not found %s", fileOpenErr.Error())
	}
	messages, messageParseErr := ReadMessages(testFile)

	if messageParseErr != nil {
		t.Fatal("Error reading messages: ", messageParseErr.Error())
	}

	if len(messages) != 2 {
		t.Errorf("should have exactly 2 messages in testfile, was %d", len(messages))
	}

	if messages[0].Section5.DataTemplateNumber != 2 {
		t.Errorf("Data template number should be 2 (found %d)", messages[0].Section5.DataTemplateNumber)
	}

	for index, data := range fixtures {
		if math.Ceil(messages[0].Section7.Data[index]*100000+.5) != data*100000 {
			t.Errorf("Expected value %f at index %d, found %f", data, index, messages[0].Section7.Data[index])
		}
	}

}
