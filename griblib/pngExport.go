package griblib

import (
	"image"
	"log"

	"fmt"
	"image/color"
	"image/png"
	"math"
	"os"
	"reflect"
)

func ExportMessagesAsPngs(messages []*Message) {
	for i, message := range messages {
		dataImage, err := imageFromMessage(message)
		if err != nil {
			log.Fatalf("Message could not be converted to image: %v\n", err)
		} else {
			err2 := writeImageToFilename(dataImage, imageFileName(i, message))
			if err2 != nil {
				log.Fatalf("Image could not be written to file: %v\n", err2)
			}
		}
	}
}

func ExportMessageAsPng(message *Message, filename string) error {
	dataImage, err := imageFromMessage(message)
	if err != nil {
		return err
	}
	return writeImageToFilename(dataImage, filename)
}

func imageFileName(messageNumber int, message *Message) string {
	dataname := ReadProductDisciplineParameters(message.Section0.Discipline, message.Section4.ProductDefinitionTemplate.ParameterCategory)
	return fmt.Sprintf("%s - discipline%d category%d messageIndex%d.png",
		dataname,
		message.Section0.Discipline,
		message.Section4.ProductDefinitionTemplate.ParameterCategory,
		messageNumber)
}

func writeImageToFilename(img image.Image, name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	if err2 := png.Encode(f, img); err2 != nil {
		err3 := f.Close()
		if err3 != nil {
			log.Fatalf("Error closing file: %v\n", err3)
		}
		return err2
	}

	if err2 := f.Close(); err2 != nil {
		return err2
	}
	return nil
}

func imageFromMessage(message *Message) (image.Image, error) {

	grid0, ok := message.Section3.Definition.(*Grid0)

	if !ok {
		err := fmt.Errorf("currently not supporting definition of type %s ", reflect.TypeOf(message.Section3.Definition))
		return nil, err
	}

	height := int(grid0.Nj)
	width := int(grid0.Ni)

	maxValue, minValue := MaxMin(message.Section7.Data)

	rgbaImage := image.NewNRGBA(image.Rect(0, 0, width, height))
	length := len(message.Section7.Data)
	if length == width*height {
		//		log.Printf("d=%d , w=%d, h=%d, wxh=%d\n", length, width, height, width*height)
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				value := message.Section7.Data[y*width+x]
				red := uint8(0)
				blue := uint8(254)
				rgbaImage.Set(x, y, color.NRGBA{
					R: red,
					G: 0,
					B: blue,
					A: RedValue(value, maxValue, minValue),
				})
			}
		}
	}
	return rgbaImage, nil
}

// RedValue returns a number between 0 and 255
func RedValue(value float64, maxValue float64, minValue float64) uint8 {
	//value  = value - 273
	if value > 0 {
		percentOfMaxValue := (math.Abs(value) + math.Abs(minValue)) / (math.Abs(maxValue) + math.Abs(minValue))
		return uint8(percentOfMaxValue * 255.0)
	}
	return 0
}

func MaxMin(float64s []float64) (float64, float64) {
	biggest, smallest := -9999999.0, 999999.0
	for _, v := range float64s {
		if v > biggest {
			biggest = v
		}
		if v < smallest {
			smallest = v
		}
	}
	return biggest, smallest
}
