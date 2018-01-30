package griblib

import (
	"testing"
	"os"
	"fmt"
)

func Test_calculcate_startStopIndexes(t *testing.T) {
	filter := GeoFilter{MinLong: 4400000, MaxLong: 32000000, MinLat: 57000000, MaxLat: 71000000}
	grid := Grid0{Di: 2500000, Dj: 2500000, Lo1: 0, Lo2: 357500000, La1: 90000000, La2: -2057483648, Ni: 144, Nj: 73}
	startNi, stopNi, startNj, stopNj := startStopIndexes(filter, grid)

	if startNi != 2 {
		t.Errorf("startNi should have been 2, was %d", startNi)
	}
	if stopNi != 13 {
		t.Errorf("stopNi should have been 13, was %d", stopNi)
	}
	if startNj != 7 {
		t.Errorf("startNj should have been 7, was %d", startNj)
	}
	if stopNj != 13 {
		t.Errorf("stopNj should have been 13, was %d", stopNj)
	}

}

func Test_filter_values_on_geofilter(t *testing.T) {
	filter := GeoFilter{MinLong: 4400000, MaxLong: 32000000, MinLat: 57000000, MaxLat: 71000000}
	grid := Grid0{Di: 2500000, Dj: 2500000, Lo1: 0, Lo2: 357500000, La1: 90000000, La2: -2057483648, Ni: 144, Nj: 73}

	// create monotonically increasing values in test-map
	testData := make([]float64, grid.Ni*grid.Nj)
	for k := range testData {
		testData[k] = float64(k)
	}

	message := Message{
		Section7: Section7{Data: testData},
		Section3: Section3{Definition: &grid},
	}

	filteredValues, err := filterValuesFromGeoFilter(message, filter)

	if err != nil {
		t.Fatal("did not filter shit")
	}

	if len(*filteredValues) != 6*11 {
		t.Errorf("Length of result is just wrong, should have been %d, was %d", 66, len(*filteredValues))
	}

}

func Test_filter_on_discipline(t *testing.T) {

	messages := []Message{
		{Section0: Section0{Discipline: 1}},
		{Section0: Section0{Discipline: 2}},
	}

	unfiltered := Filter(messages, Options{Discipline: -1})

	if len(unfiltered) != len(messages) {
		t.Error("should not filter when option is -1")
	}

	filtered := Filter(messages, Options{Discipline: 2})

	if len(filtered) != len(messages)-1 {
		t.Error("should have filtered when option is different from message")
	}
}

func Test_filter_on_category(t *testing.T) {

	messages := []Message{
		//message.Section4.ProductDefinitionTemplate.ParameterCategory
		{Section4: Section4{ProductDefinitionTemplate: Product0{ParameterCategory: 1}}},
		{Section4: Section4{ProductDefinitionTemplate: Product0{ParameterCategory: 2}}},
	}

	unfiltered := Filter(messages, Options{Category: -1})

	if len(unfiltered) != len(messages) {
		t.Error("should not filter when option is -1")
	}

	filtered := Filter(messages, Options{Category: 2})

	if len(filtered) != len(messages)-1 {
		t.Error("should have filtered when option is different from message")
	}
}

func Test_temperature_layers(t *testing.T) {
	file, err := os.Open("integrationtestdata/gfs.t00z.pgrb2.2p50.f006")
	if err != nil {
		t.Fatal("Could not open testfile")
	}
	messages, msgErr := ReadMessages(file)

	if msgErr != nil {
		t.Fatal("Error reading messages from testfile")
	}
	if len(messages) != 77 {
		t.Errorf("expected 77 messages, got %d\n", len(messages))
	}

	filtered := Filter(messages, Options{Discipline: 0, Category: 0, Surface: Surface{Value: 200, Type: 100}})

	if len(filtered) != 1 {
		t.Errorf("expected 1 messages, got %d\n", len(filtered))
	}

	fmt.Println("layers for ")
	for _, f := range filtered {
		fmt.Printf("c: %d\ts1: %d\t\tt1: %d\tv1: %d (meter over sea-level?)\n",
			f.Section4.ProductDefinitionTemplate.ParameterCategory,
			f.Section4.ProductDefinitionTemplate.FirstSurface.Scale,
			f.Section4.ProductDefinitionTemplate.FirstSurface.Type,
			f.Section4.ProductDefinitionTemplate.FirstSurface.Value,
		)
	}

}
