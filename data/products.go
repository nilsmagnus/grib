package data

import "fmt"

// http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_temp4-0.shtml
// Analysis or forecast at a horizontal level or in a horizontal layer at a point in time
type Product0 struct {
	ParameterCategory uint8
	ParameterNumber   uint8
	ProcessType       uint8
	BackgroundProcess uint8
	AnalysisProcess   uint8
	Hours             uint16
	Minutes           uint8
	TimeUnitIndicator uint8
	ForecastTime      uint32

	FirstSurface  Surface
	SecondSurface Surface
}

type Surface struct {
	Type  uint8
	Scale uint8
	Value uint32
}

func (p Product0) String() string {
	return fmt.Sprint(
        ReadProductDisciplineParameters(0, int(p.ParameterCategory)),
        " - ",
        ReadProductDisciplineCategoryParameters(0, int(p.ParameterCategory), int(p.ParameterNumber)),
    ) //, " ", p.Hours, p.Minutes, p.TimeUnitIndicator, p.ForecastTime)
}
