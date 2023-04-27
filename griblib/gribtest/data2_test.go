package gribtest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_read2_integrationtest_file_hour0(t *testing.T) {
	messages := openGrib(t, "../integrationtestdata/template5_2.grib2")

	fixtures := openCsv(t, "../integrationtestdata/template_ugrd.csv")

	assert.Len(t, messages, 2, "should have exactly 2 messages in testfile")

	assert.Equal(t, uint16(2), messages[0].Section5.DataTemplateNumber, "Data template number should be 2")

	assert.Len(t, messages[0].Data(), len(fixtures))

	assert.InEpsilonSlice(t, fixtures, messages[0].Data(), 1e-5)
}
