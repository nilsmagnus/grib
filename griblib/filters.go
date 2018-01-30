package griblib

import (
	"fmt"
	"reflect"
)

// GeoFilter is used to filter values. Only values inside the filter is returned when using this filter
type GeoFilter struct {
	MinLat  int32 `json:"minLat"`
	MaxLat  int32 `json:"maxLat"`
	MinLong int32 `json:"minLong"`
	MaxLong int32 `json:"maxLong"`
}

const (
	// LatitudeNorth is the north-most latitude value
	LatitudeNorth = 90000000
	// LatitudeSouth is the south-most latitude value
	LatitudeSouth = -90000000
	// LongitudeStart is the minimum value for longitude
	LongitudeStart = 0
	// LongitudeEnd is the maximum value for longitude
	LongitudeEnd = 360000000
)

// Filter messages with values from options
func Filter(messages []Message, options Options) (filtered []Message) {

	for _, message := range messages {
		discipline := satisfiesDiscipline(options.Discipline, message)
		category := satisfiesCategory(options.Category, message)
		surface := satisfiesSurface(options.Surface, message)
		if !surface || !discipline || !category {
			continue
		}
		if !isEmpty(options.GeoFilter) {
			if data, err := filterValuesFromGeoFilter(message, options.GeoFilter); err == nil {
				message.Section7.Data = *data
				if grid0, ok := message.Section3.Definition.(*Grid0); ok {
					updatedGrid := filteredGrid(grid0, options.GeoFilter)
					message.Section3.Definition = updatedGrid
					message.Section3.DataPointCount = uint32(len(*data))
				}

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
func satisfiesSurface(s Surface, message Message) bool {
	return s == Surface{} ||
		(message.Section4.ProductDefinitionTemplate.FirstSurface.Type == s.Type &&
			message.Section4.ProductDefinitionTemplate.FirstSurface.Value == s.Value)
}

func filteredGrid(grid0 *Grid0, filter GeoFilter) *Grid0 {
	grid0.La1 = filter.MinLat
	grid0.La2 = filter.MaxLat
	grid0.Lo1 = filter.MinLong
	grid0.Lo2 = filter.MaxLong
	startnj, stopnj, startni, stopni := startStopIndexes(filter, *grid0)
	grid0.Ni = stopni - startni
	grid0.Nj = stopnj - startnj
	return grid0
}

func isEmpty(geoFilter GeoFilter) bool {
	return geoFilter == GeoFilter{
		MinLong: LongitudeStart,
		MaxLong: LongitudeEnd,
		MinLat:  LatitudeNorth,
		MaxLat:  LatitudeSouth,
	} || geoFilter == GeoFilter{}
}

func filterValuesFromGeoFilter(message Message, filter GeoFilter) (*[]float64, error) {
	grid0, ok := message.Section3.Definition.(*Grid0)
	if ok {
		startNi, stopNi, startNj, stopNj := startStopIndexes(filter, *grid0)

		data := make([]float64, (stopNi-startNi)*(stopNj-startNj))

		filteredIndex := 0
		for j := startNj; j < stopNj; j++ {
			for i := startNi; i < stopNi; i++ {
				data[filteredIndex] = message.Section7.Data[j*grid0.Nj+i]
				filteredIndex++
			}
		}
		return &data, nil
	}
	return &message.Section7.Data, fmt.Errorf("grid not of wanted type (wanted Grid0), was %v", reflect.TypeOf(message.Section3.Definition))
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
