package gribtest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_constant(t *testing.T) {
	messages := openGrib(t, "../integrationtestdata/constant.grb")

	fixtures := openCsv(t, "../integrationtestdata/constant.csv")

	assert.Len(t, messages, 1, "should have exactly 1 message in testfile")

	assert.Equal(t, uint16(3), messages[0].Section5.DataTemplateNumber, "Data template number should be 3")

	assert.Len(t, messages[0].Data(), len(fixtures))

	assert.InEpsilonSlice(t, fixtures, messages[0].Data(), 1e-5)
}
