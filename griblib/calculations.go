package griblib

import (
	"fmt"
	"reflect"
)

// AverageValueBasic takes a GeoFilter, Grid0 and data to calculate the average value within that area.
// See GeoFilter for how to define an area
func AverageValueBasic(filter GeoFilter, grid0 *Grid0, data []float64) (float64, error) {
	startNi, stopNi, startNj, stopNj := StartStopIndexes(filter, *grid0)

	numberOfDataPoints := (stopNi - startNi) * (stopNj - startNj)
	value := 0.0

	for j := startNj; j < stopNj; j++ {
		for i := startNi; i < stopNi; i++ {
			value += data[j*grid0.Nj+i]
		}
	}
	return value / float64(numberOfDataPoints), nil
}

func AverageValue(filter GeoFilter, message *Message) (float64, error) {
	grid0, ok := message.Section3.Definition.(*Grid0)
	data := message.Section7.Data
	if ok {
		return AverageValueBasic(filter, grid0, data)
	}
	return -1, fmt.Errorf("grid not of wanted type (wanted Grid0), was %v", reflect.TypeOf(message.Section3.Definition))
}
