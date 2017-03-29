package griblib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

const (
	ExportNone              = 0
	PrintMessageDisciplines = 1
	PrintMessageCategories  = 2
	ExportJsonToConsole     = 3
)

// Export exports messages to the supported formats
func Export(messages []Message, options Options) {
	switch options.ExportType {
	case ExportNone:
	case PrintMessageDisciplines:
		printDisciplines(messages)
	case PrintMessageCategories:
		printCategories(messages)
	case ExportJsonToConsole:
		exportJSONConsole(messages, options)
	}
}

func printDisciplines(messages []Message) {
	for _, message := range messages {
		fmt.Println(ReadDiscipline(message.Section0.Discipline))
	}
}

func printCategories(messages []Message) {
	for _, m := range messages {
		category := m.Section4.ProductDefinitionTemplate.ParameterCategory
		templateNumber := m.Section4.ProductDefinitionTemplateNumber
		fmt.Println(ReadProductDisciplineParameters(templateNumber, category))
	}
}

func exportJSONConsole(messages []Message, options Options) {
	fmt.Println("[")
	for _, message := range messages {
		export(&message, options)
		fmt.Println(",")
	}
	fmt.Println("]")
}

func export(m *Message, options Options) {
	templateNumber := m.Section4.ProductDefinitionTemplateNumber
	template := m.Section4.ProductDefinitionTemplate
	category := template.ParameterCategory
	number := template.ParameterNumber

	d := make(map[string]interface{})

	d["type"] = ReadDataType(m.Section1.Type)
	d["template"] = ReadProductDefinitionTemplateNumber(templateNumber)
	d["category"] = ReadProductDisciplineParameters(templateNumber, category)
	d["parameter"] = ReadProductDisciplineCategoryParameters(templateNumber, category, number)
	d["grid"] = ReadGridDefinitionTemplateNumber(int(m.Section3.TemplateNumber))
	d["surface1"] = ReadSurfaceTypesUnits(int(m.Section4.ProductDefinitionTemplate.FirstSurface.Type))
	d["surface1value"] = m.Section4.ProductDefinitionTemplate.FirstSurface.Value
	d["surface1scale"] = m.Section4.ProductDefinitionTemplate.FirstSurface.Scale
	d["surface2"] = ReadSurfaceTypesUnits(int(m.Section4.ProductDefinitionTemplate.SecondSurface.Type))
	d["surface2value"] = m.Section4.ProductDefinitionTemplate.SecondSurface.Value
	if options.DataExport {
		d["data"] = m.Section7.Data
	}

	grid, _ := m.Section3.Definition.(Grid)
	for k, v := range grid.Export() {
		d[k] = v
	}

	// json print
	js, _ := json.Marshal(d)
	var out bytes.Buffer
	json.Compact(&out, js)
	out.WriteTo(os.Stdout)
	fmt.Println("")
}
