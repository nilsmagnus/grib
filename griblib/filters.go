package griblib

import (
	"fmt"
	"reflect"
)

type GeoFilter struct {
	MinLat  int32 `json:"minLat"`
	MaxLat  int32 `json:"maxLat"`
	MinLong int32 `json:"minLong"`
	MaxLong int32 `json:"maxLong"`
}

const (
	LatitudeNorth  = 90000000
	LatitudeSouth  = -90000000
	LongitudeStart = 0
	LongitudeEnd   = 360000000
)

func Filter(messages []Message, options Options) (filtered []Message) {

	for _, message := range messages {
		discipline := satisfiesDiscipline(options.Discipline, message)
		category := satisfiesCategory(options.Category, message)
		if !isEmpty(options.GeoFilter) {
			if data, err := filterValuesFromGeoFilter(message, options.GeoFilter); err == nil {
				message.Section7.Data = *data
			} else {
				fmt.Println(err.Error())
			}
		}
		if discipline && category {
			filtered = append(filtered, message)
		}

	}

	return filtered
}

func isEmpty(geoFilter GeoFilter) bool {
	return geoFilter == GeoFilter{MinLong: LongitudeStart, MaxLong: LongitudeEnd, MinLat: LatitudeNorth, MaxLat: LatitudeSouth}
}

func filterValuesFromGeoFilter(message Message, filter GeoFilter) (*[]int64, error) {
	grid, ok := message.Section3.Definition.(*Grid0)
	if ok {
		startNi, stopNi, startNj, stopNj := startStopIndexes(filter, *grid)

		data := make([]int64, (stopNi-startNi)*(stopNj-startNj))

		filteredIndex := 0
		for i := startNj; i < stopNj; i++ {
			for j := startNi; j < stopNi; j++ {
				data[filteredIndex] = message.Section7.Data[i*grid.Nj+j]
				filteredIndex++
			}
		}
		return &data, nil
	} else {
		return &message.Section7.Data, fmt.Errorf("grid not of wanted type (wanted Grid0), was %s", reflect.TypeOf(message.Section3.Definition))
	}
}

func startStopIndexes(filter GeoFilter, grid Grid0) (uint32, uint32, uint32, uint32) {

	// ni is number of points west-east
	startNi := uint32(filter.MinLong/grid.Di) + 1
	stopNi := uint32(filter.MaxLong/grid.Di) + 1

	// nj is number of points north-south
	startNj := uint32((LatitudeNorth - filter.MaxLat) / grid.Dj)
	stopNj := uint32((LatitudeNorth - filter.MinLat) / grid.Dj)

	return startNi, stopNi, startNj, stopNj
}

func satisfiesDiscipline(discipline int, message Message) bool {
	if discipline == -1 {
		return true
	}
	if discipline == int(message.Section0.Discipline) {
		return true
	}
	return false
}

func satisfiesCategory(product int, message Message) bool {
	if product == -1 {
		return true
	}
	if product == int(message.Section4.ProductDefinitionTemplate.ParameterCategory) {
		return true
	}
	return false
}
