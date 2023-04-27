package gribtest

import (
	"bufio"
	"os"
	"strconv"
	"testing"

	"github.com/nilsmagnus/grib/griblib"
)

func openGrib(t *testing.T, filename string) []*griblib.Message {
	t.Helper()

	testFile, fileOpenErr := os.Open(filename)
	if fileOpenErr != nil {
		t.Fatal("Grib file for integration tests not found")
	}

	defer func() {
		_ = testFile.Close()
	}()

	messages, messageParseErr := griblib.ReadMessages(testFile)
	if messageParseErr != nil {
		t.Fatal("Error reading messages: ", messageParseErr.Error())
	}

	return messages
}

func openCsv(t *testing.T, filename string) []float64 {
	t.Helper()

	resultFile, csvFileOpenErr := os.Open(filename)
	if csvFileOpenErr != nil {
		t.Fatalf("CSV file for integration tests not found %s", csvFileOpenErr.Error())
	}
	defer resultFile.Close()

	fixtures := []float64{}
	scanner := bufio.NewScanner(resultFile)
	for scanner.Scan() {
		f, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			t.Fatalf("Could not parse CSV file %s", err.Error())
		}
		fixtures = append(fixtures, f)
	}

	return fixtures
}
