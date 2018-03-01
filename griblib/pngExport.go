package griblib

import "image"

func ExportMessagesAsPngs(messages []Message, options Options) {

	for i, message := range messages {
		dataImage := imageFromMessage(message)
	}
}
func imageFromMessage(message Message) image.Image {

}
