package griblib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

const (
	//ExportNone - do not export anything
	ExportNone = 0
	// PrintMessageDisciplines - only print disciplines for the sections
	PrintMessageDisciplines = 1
	// PrintMessageCategories - only print categories
	PrintMessageCategories = 2
	// ExportJSONToConsole - export json to console
	ExportJSONToConsole = 3
	// ExportToPNG - export data as a png
	ExportToPNG = 4
)

// Export exports messages to the supported formats
func Export(messages []Message, options Options) {
	switch options.ExportType {
	case ExportNone:
	case PrintMessageDisciplines:
		printDisciplines(messages)
	case PrintMessageCategories:
		printCategories(messages)
	case ExportJSONToConsole:
		exportJSONConsole(messages)
	case ExportToPNG:
		ExportMessagesAsPngs(messages)
	default:
		fmt.Printf("Error: Export type %d not supported. \n", options.ExportType)
	}
}

func printDisciplines(messages []Message) {
	for _, message := range messages {
		fmt.Println(DisciplineDescription(message.Section0.Discipline))
	}
}

func printCategories(messages []Message) {
	for _, m := range messages {
		category := m.Section4.ProductDefinitionTemplate.ParameterCategory
		discipline := m.Section0.Discipline
		fmt.Println(ReadProductDisciplineParameters(discipline, category))
	}
}

func exportJSONConsole(messages []Message) {
	fmt.Println("[")
	for _, message := range messages {
		export(&message)
		fmt.Println(",")
	}
	fmt.Println("]")
}

func export(m *Message) {

	// json print
	js, _ := json.Marshal(m)
	var out bytes.Buffer
	json.Compact(&out, js)
	out.WriteTo(os.Stdout)
	fmt.Println("")
}
