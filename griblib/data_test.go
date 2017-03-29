package griblib

import (
	"math"
	"os"
	"testing"
)

func Test_read_integrationtest_file(t *testing.T) {
	testFile, fileOpenErr := os.Open("integrationtestdata/gfs.t00z.pgrb2.2p50.f012")

	if fileOpenErr != nil {
		t.Fatal("Grib file for integration tests not found")
	}
	messages, messageParseErr := ReadMessages(testFile, Options{
		MaximumNumberOfMessages: math.MaxInt32,
		Discipline:              -1,
	})

	if messageParseErr != nil {
		t.Fatal("Error reading messages: ", messageParseErr.Error())
	}

	if len(messages) != 77 {
		t.Error("should have exactly 77 message in testfile")
	}

}
