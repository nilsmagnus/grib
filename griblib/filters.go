package griblib

// Filter filters messages based on flags in options.
//
// Currently only supports filtering on discipline and category
//
func Filter(messages []Message, options Options) (filtered []Message) {

	for _, message := range messages {
		if satisfiesDiscipline(options.Discipline, message) && satisfiesCategory(options.Category, message) {
			filtered = append(filtered, message)
		}
	}

	return filtered
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
