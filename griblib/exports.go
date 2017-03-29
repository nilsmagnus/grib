package griblib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

const (
	ExportNone          = 0
	ExportJsonToConsole = 1
)

// Export exports messages to the supported formats
func Export(messages []Message, options Options) {
	switch options.ExportType {
	case ExportNone:
	case ExportJsonToConsole:
		exportJSONConsole(messages)
	}
}

func exportJSONConsole(messages []Message) {
	for _, message := range messages {
		export(&message)
	}
}

func export(m *Message) {
	templateNumber := int(m.Section4.ProductDefinitionTemplateNumber)
	template := m.Section4.ProductDefinitionTemplate
	category := int(template.ParameterCategory)
	number := int(template.ParameterNumber)

	d := make(map[string]interface{})

	d["type"] = ReadDataType(int(m.Section1.Type))
	d["template"] = ReadProductDefinitionTemplateNumber(templateNumber)
	d["category"] = ReadProductDisciplineParameters(templateNumber, category)
	d["parameter"] = ReadProductDisciplineCategoryParameters(templateNumber, category, number)
	d["grid"] = ReadGridDefinitionTemplateNumber(int(m.Section3.TemplateNumber))
	d["surface1"] = ReadSurfaceTypesUnits(int(m.Section4.ProductDefinitionTemplate.FirstSurface.Type))
	d["surface1value"] = m.Section4.ProductDefinitionTemplate.FirstSurface.Value
	d["surface1scale"] = m.Section4.ProductDefinitionTemplate.FirstSurface.Scale
	d["surface2"] = ReadSurfaceTypesUnits(int(m.Section4.ProductDefinitionTemplate.SecondSurface.Type))
	d["surface2value"] = m.Section4.ProductDefinitionTemplate.SecondSurface.Value
	d["data"] = m.Section7.Data

	for k, v := range m.Section3.Definition.Export() {
		d[k] = v
	}

	// json print
	js, _ := json.Marshal(d)
	var out bytes.Buffer
	json.Indent(&out, js, "", "\t")
	out.WriteTo(os.Stdout)
	fmt.Println("")
}
