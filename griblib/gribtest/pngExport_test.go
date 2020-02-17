package gribtest

import (
	"fmt"
	"os"
	"testing"

	"github.com/nilsmagnus/grib/griblib"
)

func beforeTests(t *testing.T) func(t *testing.T) {
	os.MkdirAll("testoutput", os.ModePerm)

	return func(t *testing.T) {
		os.RemoveAll("testoutput")
	}
}

func Test_message_to_png(t *testing.T) {
	defer beforeTests(t)(t)

	testInput, err := os.Open("../integrationtestdata/gfs.t00z.pgrb2.2p50.f012")

	if err != nil {
		t.Fatal(err)
	}

	messages, messageErr := griblib.ReadMessages(testInput)

	if messageErr != nil {
		t.Fatal(messageErr)
	}

	for _, message := range messages {
		if message.Section0.Discipline == 0 &&
			message.Section4.ProductDefinitionTemplate.ParameterCategory == 1 &&
			message.Section4.ProductDefinitionTemplate.FirstSurface.Type == 1 {
			dataname := griblib.ReadProductDisciplineParameters(message.Section0.Discipline, message.Section4.ProductDefinitionTemplate.ParameterCategory)

			errf := griblib.ExportMessageAsPng(message, fmt.Sprintf("%s.png", dataname))
			fmt.Printf("wrote image to %s \n", dataname)
			if errf != nil {
				t.Error(errf)
			}
		}
	}

}

func Test_maxmin(t *testing.T) {
	max, min := griblib.MaxMin([]float64{0, -154, 54, 64, -10})
	if max != 64.0 {
		t.Errorf("Expected max to be 64, was %f", max)
	}
	if min != -154.0 {
		t.Errorf("Expected min to be -154, was %f", min)
	}
}

func Test_redvalue(t *testing.T) {

	if red := griblib.RedValue(0, 100, -100); red != 0 {
		t.Errorf("expected blue to be 125 , but was %d", red)
	}
	if red := griblib.RedValue(50, 100, -100); red != 191 {
		t.Errorf("expected blue to be 191 , but was %d", red)
	}
	if red := griblib.RedValue(50, 100, 0); red != 127 {
		t.Errorf("expected blue to be 127 , but was %d", red)
	}
}
