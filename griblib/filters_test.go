package griblib

import "testing"

func Test_filter_on_discipline(t *testing.T) {

	messages := []Message{
		{Section0: Section0{Discipline: 1}},
		{Section0: Section0{Discipline: 2}},
	}

	unfiltered := Filter(messages, Options{Discipline: -1})

	if len(unfiltered) != len(messages) {
		t.Error("should not filter when option is -1")
	}

	filtered := Filter(messages, Options{Discipline: 2})

	if len(filtered) != len(messages)-1 {
		t.Error("should have filtered when option is different from message")
	}
}

func Test_filter_on_category(t *testing.T) {

	messages := []Message{
		//message.Section4.ProductDefinitionTemplate.ParameterCategory
		{Section4: Section4{ProductDefinitionTemplate: Product0{ParameterCategory: 1}}},
		{Section4: Section4{ProductDefinitionTemplate: Product0{ParameterCategory: 2}}},
	}

	unfiltered := Filter(messages, Options{Category: -1})

	if len(unfiltered) != len(messages) {
		t.Error("should not filter when option is -1")
	}

	filtered := Filter(messages, Options{Category: 2})

	if len(filtered) != len(messages)-1 {
		t.Error("should have filtered when option is different from message")
	}
}
