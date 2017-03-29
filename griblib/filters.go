package griblib

// Filter filters messages based on flags in options.
//
// Currently only supports filtering on discipline and category
//

type GeoFilter struct {
	MinLat  float32
	MaxLat  float32
	MinLong float32
	MaxLong float32
}

func isEmpty(geoFilter GeoFilter) bool {
	return geoFilter == GeoFilter{}
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
	/*	grid, ok := message.Section3.Definition.(Grid0)
		if ok {
			fmt.Printf("should filter grid %v on filter %v", grid, filter)
		}*/
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
