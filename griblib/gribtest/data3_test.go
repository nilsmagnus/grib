package gribtest

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_read_integrationtest_file(t *testing.T) {
	messages := openGrib(t, "../integrationtestdata/gfs.t18z.pgrb2.1p00.f003")

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
	messages := openGrib(t, "../integrationtestdata/gfs.t06z.pgrb2.1p00.f000")

	if len(messages) != 354 {
		t.Errorf("should have exactly 354 message in testfile, was %d", len(messages))
	}

}

func Test_read3_integrationtest_file_hour0(t *testing.T) {
	messages := openGrib(t, "../integrationtestdata/template5_3.grib2")

	fixtures := openCsv(t, "../integrationtestdata/template_ugrd.csv")

	assert.Len(t, messages, 2, "should have exactly 2 messages in testfile")

	assert.Equal(t, uint16(3), messages[0].Section5.DataTemplateNumber, "Data template number should be 3")

	assert.Len(t, messages[0].Data(), len(fixtures))

	assert.InEpsilonSlice(t, fixtures, messages[0].Data(), 1e-5)
}
