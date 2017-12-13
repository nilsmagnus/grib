package griblib

// Product0 http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_temp4-0.shtml
// Analysis or forecast at a horizontal level or in a horizontal layer at a point in time
type Product0 struct {
	ParameterCategory uint8   `json:"parameterCategory"`
	ParameterNumber   uint8   `json:"parameterNumber"`
	ProcessType       uint8   `json:"processType"`
	BackgroundProcess uint8   `json:"backgroundProcess"`
	AnalysisProcess   uint8   `json:"analysisProcess"`
	Hours             uint16  `json:"hours"`
	Minutes           uint8   `json:"minutes"`
	TimeUnitIndicator uint8   `json:"timeUnitIndicator"`
	ForecastTime      uint32  `json:"forecastTime"`
	FirstSurface      Surface `json:"firstSurface"`
	SecondSurface     Surface `json:"secondSurface"`
}

//Product1 http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_temp4-1.shtml
type Product1 struct {
	Product0
	EnsembleForecastType    uint8 `json:"ensembleForecastType"`
	PertubationNumber       uint8 `json:"pertubationNumber"`
	ForecastInEnsembleCount uint8 `json:"forecastInEnsembleCount"`
}

//Product2 http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_temp4-2.shtml
type Product2 struct {
	Product0
	DerivedForecast         uint8 `json:"derivedForecast"`
	ForecastInEnsembleCount uint8 `json:"forecastInEnsembleCount"`
}

//Product5 http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_temp4-5.shtml
type Product5 struct {
	Product0
	ForecastProbabilityNumber  uint8  `json:"forecastProbabilityNumber"`
	TotalForecastProbabilities uint8  `json:"totalForecastProbabilities"`
	ProbabilityType            uint8  `json:"probabilityType"`
	ScaleFactorLowerLimit      uint8  `json:"scaleFactorLowerLimit"`
	ScaleValueLowerLimit       uint32 `json:"scaleValueLowerLimit"`
	ScaleFactorUpperLimit      uint8  `json:"scaleFactorUpperLimit"`
	ScaleValueUpperLimit       uint32 `json:"scaleValueUpperLimit"`
}

//Product6 http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_temp4-6.shtml
type Product6 struct {
	Product0
	PercentileValue uint8 `json:"percentileValue"` // 0-100
}

//Product7 http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_temp4-7.shtml
type Product7 struct {
	Product0
}

//Product8 http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_temp4-8.shtml
type Product8 struct {
	Product0
	Time                              Time                     `json:"time"`
	NumberOfIntervalTimeRanges        uint8                    `json:"numberOfIntervalTimeRanges"` // length of last datatype
	TotalMissingDataValuesCount       uint32                   `json:"totalMissingDataValuesCount"`
	TimeRangeSpecification1           TimeRangeSpecification   `json:"timeRangeSpecification1"`
	TimeRangeSpecification2           TimeRangeSpecification   `json:"timeRangeSpecification2"`           // 59-70
	AdditionalTimeRangeSpecifications []TimeRangeSpecification `json:"additionalTimeRangeSpecifications"` // 71-n
}

//TimeRangeSpecification describes timerange for products
type TimeRangeSpecification struct {
	StatisticalFieldCalculationProcess                     uint8  `json:"statisticalFieldCalculationProcess"`                     // 47
	IncrementBetweenSuccessiveFieldsType                   uint8  `json:"incrementBetweenSuccessiveFieldsType"`                   // 48
	IncrementBetweenSuccessiveFieldsRangeTimeUnitIndicator uint8  `json:"incrementBetweenSuccessiveFieldsRangeTimeUnitIndicator"` // 49
	StatististicalProcessTimeLength                        uint32 `json:"statististicalProcessTimeLength"`                        // 50-53
	IncrementBetweenSuccessiveFieldsTimeUnitIndicator      uint8  `json:"incrementBetweenSuccessiveFieldsTimeUnitIndicator"`      // 54
	TimeIncrementBetweenSuccessiveField                    uint32 `json:"timeIncrementBetweenSuccessiveField"`                    // 55-58
}

//Surface describes a surface for a product, see http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_table4-5.shtml
type Surface struct {
	Type  uint8  `json:"type"` // type 220: Planetary Boundary Layer
	Scale uint8  `json:"scale"`
	Value uint32 `json:"value"` // e.g. meters above sea-level
}
