package gribtest

import (
	"log"
	"os"
	"testing"

	"github.com/nilsmagnus/grib/griblib"
)

func Test_calculcate_startStopIndexes(t *testing.T) {
	filter := griblib.GeoFilter{MinLong: 4_400_000, MaxLong: 32_000_000, MinLat: 71_000_000, MaxLat: 57_000_000}
	grid := griblib.Grid0{Di: 2_500_000, Dj: 2_500_000, Lo1: 0, Lo2: 357_500_000, La1: 90_000_000, La2: -2057483648, Ni: 144, Nj: 73}
	startNi, stopNi, startNj, stopNj := griblib.StartStopIndexes(filter, grid)

	if startNi != 1 {
		t.Errorf("startNi should have been 2, was %d", startNi)
	}
	if stopNi != 12 {
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
	filter := griblib.GeoFilter{MinLong: 4400000, MaxLong: 32000000, MinLat: 71000000, MaxLat: 57000000}
	grid := griblib.Grid0{Di: 2500000, Dj: 2500000, Lo1: 0, Lo2: 357500000, La1: 90000000, La2: -2057483648, Ni: 144, Nj: 73}

	// create monotonically increasing values in test-map
	testData := make([]float64, grid.Ni*grid.Nj)
	for k := range testData {
		testData[k] = float64(k)
	}

	message := griblib.Message{
		Section7: griblib.Section7{Data: testData},
		Section3: griblib.Section3{Definition: &grid},
	}

	filteredValues, err := griblib.FilterValuesFromGeoFilter(&message, filter)

	if err != nil {
		t.Fatal("did not filter shit")
	}

	if len(*filteredValues) != 6*11 {
		t.Errorf("Length of result is just wrong, should have been %d, was %d", 66, len(*filteredValues))
	}

}

func Test_filter_on_discipline(t *testing.T) {

	messages := []*griblib.Message{
		{Section0: griblib.Section0{Discipline: 1}},
		{Section0: griblib.Section0{Discipline: 2}},
	}

	unfiltered := griblib.Filter(messages, griblib.Options{Discipline: -1})

	if len(unfiltered) != len(messages) {
		t.Error("should not filter when option is -1")
	}

	filtered := griblib.Filter(messages, griblib.Options{Discipline: 2})

	if len(filtered) != len(messages)-1 {
		t.Error("should have filtered when option is different from message")
	}
}

func Test_filter_on_category(t *testing.T) {

	messages := []*griblib.Message{
		//message.Section4.ProductDefinitionTemplate.ParameterCategory
		{Section4: griblib.Section4{ProductDefinitionTemplate: griblib.Product0{ParameterCategory: 1}}},
		{Section4: griblib.Section4{ProductDefinitionTemplate: griblib.Product0{ParameterCategory: 2}}},
	}

	unfiltered := griblib.Filter(messages, griblib.Options{Category: -1})

	if len(unfiltered) != len(messages) {
		t.Error("should not filter when option is -1")
	}

	filtered := griblib.Filter(messages, griblib.Options{Category: 2})

	if len(filtered) != len(messages)-1 {
		t.Error("should have filtered when option is different from message")
	}
}

func Test_temperature_layers(t *testing.T) {
	file, err := os.Open("../integrationtestdata/gfs.t00z.pgrb2.2p50.f006")
	if err != nil {
		t.Fatal("Could not open testfile")
	}
	messages, msgErr := griblib.ReadMessages(file)

	if msgErr != nil {
		t.Fatal("Error reading messages from testfile")
	}
	if len(messages) != 415 {
		t.Errorf("expected 415 messages, got %d\n", len(messages))
	}

	filtered := griblib.Filter(messages, griblib.Options{Discipline: 0, Category: 0, Surface: griblib.Surface{Value: 200, Type: 100}})

	if len(filtered) != 1 {
		t.Errorf("expected 1 messages, got %d\n", len(filtered))
	}

	log.Println("layers for ")
	for _, f := range filtered {
		log.Printf("c: %d\ts1: %d\t\tt1: %d\tv1: %d (meter over sea-level?)\n",
			f.Section4.ProductDefinitionTemplate.ParameterCategory,
			f.Section4.ProductDefinitionTemplate.FirstSurface.Scale,
			f.Section4.ProductDefinitionTemplate.FirstSurface.Type,
			f.Section4.ProductDefinitionTemplate.FirstSurface.Value,
		)
	}

}
