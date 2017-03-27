package main

import (
	"fmt"
	"github.com/nilsmagnus/grib/data"
	"os"
	"encoding/json"
	"bytes"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	category := 0 // temperature
	product := 6 // temperature

	f, err := os.Open("gfs.t00z.pgrb2.0p25.f000")
	check(err)

	for {
		//sections,
		message, err := data.ReadMessage(f)
		if err != nil && err.Error() == "EOF" {
			fmt.Println("END")
			break
		}
		check(err)

		if message.Section4.ProductDefinitionTemplate.ParameterCategory == uint8(category) && message.Section4.ProductDefinitionTemplate.ParameterNumber == uint8(product) {
			fmt.Println("Found!")

			export(&message)
			break;
		}
	}

	check(f.Close())
}

func export(m *data.Message) {
	templateNumber := int(m.Section4.ProductDefinitionTemplateNumber)
	template := m.Section4.ProductDefinitionTemplate
	category := int(template.ParameterCategory)
	number := int(template.ParameterNumber)

	d := make(map[string]interface{})

	d["type"] = data.ReadDataType(int(m.Section1.Type));
	d["template"] = data.ReadProductDefinitionTemplateNumber(templateNumber);
	d["category"] = data.ReadProductDisciplineParameters(templateNumber, category);
	d["parameter"] = data.ReadProductDisciplineCategoryParameters(templateNumber, category, number);
	d["grid"] = data.ReadGridDefinitionTemplateNumber(int(m.Section3.TemplateNumber));
	d["surface1"] = data.ReadSurfaceTypesUnits(int(m.Section4.ProductDefinitionTemplate.FirstSurface.Type));
	d["surface1value"] = m.Section4.ProductDefinitionTemplate.FirstSurface.Value;
	d["surface1scale"] = m.Section4.ProductDefinitionTemplate.FirstSurface.Scale;
	d["surface2"] = data.ReadSurfaceTypesUnits(int(m.Section4.ProductDefinitionTemplate.SecondSurface.Type));
	d["surface2value"] = m.Section4.ProductDefinitionTemplate.SecondSurface.Value;
	d["data"] = m.Section7.Data;

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