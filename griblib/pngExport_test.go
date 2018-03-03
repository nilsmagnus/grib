package griblib

import (
	"os"
	"testing"
)

func Test_message_to_png(t *testing.T) {

	testInput, err := os.Open("integrationtestdata/gfs.t18z.pgrb2.1p00.f003")

	if err != nil {
		t.Fatal(err)
	}

	messages, messageErr := ReadMessages(testInput)

	if messageErr != nil {
		t.Fatal(messageErr)
	}

	image := imageFromMessage(messages[0])

	if image == nil {
		t.Error("Image is nill")
	}

}
