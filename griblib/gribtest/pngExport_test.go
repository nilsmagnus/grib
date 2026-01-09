package gribtest

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/nilsmagnus/grib/griblib"
)

func beforeTests(t *testing.T) func(t *testing.T) {
	err := os.MkdirAll("testoutput", 0750)
	if err != nil {
		t.Errorf("Could not create testoutput directory: %v", err)
		t.Fail()
	}

	return func(t *testing.T) {
		err2 := os.RemoveAll("testoutput")
		if err2 != nil {
			t.Errorf("Could not remove testoutput directory: %v", err2)
			t.Fail()
		}
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
			log.Printf("wrote image to %s \n", dataname)
			if errf != nil {
				t.Error(errf)
			}
		}
	}

}

func Test_maxmin(t *testing.T) {
	biggest, smallest := griblib.MaxMin([]float64{0, -154, 54, 64, -10})
	if biggest != 64.0 {
		t.Errorf("Expected biggest to be 64, was %f", biggest)
	}
	if smallest != -154.0 {
		t.Errorf("Expected smallest to be -154, was %f", smallest)
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
