package griblib

import (
	"image"

	"fmt"
	"image/color"
	"log"
	"math"
	"reflect"
	"sort"
)

func ExportMessagesAsPngs(messages []Message) {
	for i, message := range messages {
		dataImage := imageFromMessage(message)
		writeImageToFilename(dataImage, imageFileName(i, message))
	}
}

func imageFileName(messageNumber int, message Message) string {

}

func writeImageToFilename(message image.Image, name string) {

}
func imageFromMessage(message Message) image.Image {

	if grid0, ok := message.Section3.Definition.(*Grid0); !ok {
		err := fmt.Errorf("Currently not supporting definition of type %s ", reflect.TypeOf(message.Section3.Definition))
		log.Fatal(err)
		return nil
	} else {
		height := int(grid0.Di)
		width := int(grid0.Dj)

		maxValue, minValue := maxMin(message.Section7.Data)

		rgbaImage := image.NewNRGBA(image.Rect(0, 0, width, height))
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				rgbaImage.Set(x, y, color.NRGBA{
					R: redValue(message.Section7.Data[y*height+x], maxValue, minValue),
					G: 0,
					B: blueValue(message.Section7.Data[y*height+x], maxValue, minValue),
					A: 255,
				})
			}
		}
		return rgbaImage
	}
}

// returns a number between 0 and 255
func blueValue(value float64, maxValue float64, minValue float64) uint8 {
	if value > 0 {
		return 0
	}
	percentOfMaxValue := (math.Abs(value) + math.Abs(minValue)) / (math.Abs(maxValue) + math.Abs(minValue))
	return uint8(percentOfMaxValue * 255.0)
}

// returns a number between 0 and 255
func redValue(value float64, maxValue float64, minValue float64) uint8 {
	if value < 0 {
		return 0
	}
	percentOfMaxValue := (math.Abs(value) + math.Abs(minValue)) / (math.Abs(maxValue) + math.Abs(minValue))
	return uint8(percentOfMaxValue * 255.0)
}

func maxMin(float64s []float64) (float64, float64) {
	tmp := make([]float64, len(float64s))
	copy(tmp, float64s)
	sort.Slice(float64s, func(i, j int) bool {
		return float64s[i] < float64s[j]
	})
	sort.Float64s(float64s)
	return float64s[0], float64s[len(float64s)-1]
}
