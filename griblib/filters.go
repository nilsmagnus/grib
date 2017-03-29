package griblib

import "fmt"

// Filter filters messages based on flags in options.
//
// Currently only supports filtering on discipline and category
//

type GeoFilter struct {
	MinLat  int32
	MaxLat  int32
	MinLong int32
	MaxLong int32
}

func isEmpty(geoFilter GeoFilter) bool {
	return geoFilter == GeoFilter{MinLat:0, MaxLat:360, MinLong:-90, MaxLong:90}
}

func Filter(messages []Message, options Options) (filtered []Message) {

	for _, message := range messages {
		discipline := satisfiesDiscipline(options.Discipline, message)
		category := satisfiesCategory(options.Category, message)
		if discipline && category {
			filtered = append(filtered, message)
		}
		if !isEmpty(options.GeoFilter) {
			filterValuesFromGeoFilter(&message, options.GeoFilter)
		}

	}

	return filtered
}
func filterValuesFromGeoFilter(message *Message, filter GeoFilter) {
	grid, ok := message.Section3.Definition.(Grid0)
	if ok {
		fmt.Println(grid)

		startNi, stopNi, startNj, stopNj := startStopIndexes(filter, grid)

		filteredValues := make([]int64, (stopNi - startNi) * (stopNj - startNj))

		filteredIndex := 0
		for i := startNj; i < stopNj; i++ {
			for j := startNi; j < stopNi; j++ {
				filteredValues[filteredIndex] = message.Section7.Data[i + j * grid.Nj]
				filteredIndex++
			}
		}

		message.Section7.Data = filteredValues
	}
}

func startStopIndexes(filter GeoFilter, grid Grid0) (uint32, uint32, uint32, uint32) {
	startNi := uint32(filter.MinLat / grid.Di)
	stopNi := uint32(filter.MaxLat / grid.Di)
	startNj := uint32(filter.MinLong / grid.Dj)
	stopNj := uint32(filter.MaxLong / grid.Dj)
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
