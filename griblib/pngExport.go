package griblib

import (
	"image"

	"fmt"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"reflect"
)

func ExportMessagesAsPngs(messages []*Message) {
	for i, message := range messages {
		dataImage := imageFromMessage(message)
		writeImageToFilename(dataImage, imageFileName(i, message))
	}
}

func imageFileName(messageNumber int, message *Message) string {
	dataname := ReadProductDisciplineParameters(message.Section0.Discipline, message.Section4.ProductDefinitionTemplate.ParameterCategory)
	return fmt.Sprintf("%s - discipline%d category%d messageIndex%d.png",
		dataname,
		message.Section0.Discipline,
		message.Section4.ProductDefinitionTemplate.ParameterCategory,
		messageNumber)
}

func writeImageToFilename(img image.Image, name string) {
	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func imageFromMessage(message *Message) image.Image {

	grid0, ok := message.Section3.Definition.(*Grid0)

	if !ok {
		err := fmt.Errorf("Currently not supporting definition of type %s ", reflect.TypeOf(message.Section3.Definition))
		log.Fatal(err)
		return nil
	}

	height := int(grid0.Ni)
	width := int(grid0.Nj)

	maxValue, minValue := maxMin(message.Section7.Data)

	rgbaImage := image.NewNRGBA(image.Rect(0, 0, width, height))
	length := len(message.Section7.Data)
	if length == width*height {
		//		log.Printf("d=%d , w=%d, h=%d, wxh=%d\n", length, width, height, width*height)
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				value := message.Section7.Data[y*width+x]
				red := redValue(value, maxValue, minValue)
				blue := blueValue(value, maxValue, minValue)
				rgbaImage.Set(x, y, color.NRGBA{
					R: red,
					G: 0,
					B: blue,
					A: 255,
				})
			}
		}
	}
	return rgbaImage
}

// returns a number between 0 and 255
func blueValue(value float64, maxValue float64, minValue float64) uint8 {
	value  = value - 273
	if value < 0 {
		percentOfMaxValue := (math.Abs(value) + math.Abs(minValue)) / (math.Abs(maxValue) + math.Abs(minValue))
		return uint8(percentOfMaxValue * 255.0)
	}
	return 0
}

// returns a number between 0 and 255
func redValue(value float64, maxValue float64, minValue float64) uint8 {
	value  = value - 273
	if value > 0 {
		percentOfMaxValue := (math.Abs(value) + math.Abs(minValue)) / (math.Abs(maxValue) + math.Abs(minValue))
		return uint8(percentOfMaxValue * 255.0)
	}
	return 0
}

func maxMin(float64s []float64) (float64, float64) {
	max, min := -9999999.0, 999999.0
	for _, v := range float64s {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max, min
}
