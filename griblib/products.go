package griblib

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
	FirstSurface      Surface
	SecondSurface     Surface
}

// http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_temp4-1.shtml
type Product1 struct {
	Product0
	EnsembleForecastType    uint8
	PertubationNumber       uint8
	ForecastInEnsembleCount uint8
}

// http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_temp4-2.shtml
type Product2 struct {
	Product0
	DerivedForecast         uint8
	ForecastInEnsembleCount uint8
}

// http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_temp4-5.shtml
type Product5 struct {
	Product0
	ForecastProbabilityNumber  uint8
	TotalForecastProbabilities uint8
	ProbabilityType            uint8
	ScaleFactorLowerLimit      uint8
	ScaleValueLowerLimit       uint32
	ScaleFactorUpperLimit      uint8
	ScaleValueUpperLimit       uint32
}

// http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_temp4-6.shtml
type Product6 struct {
	Product0
	PercentileValue uint8 // 0-100
}

// http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_temp4-7.shtml
type Product7 struct {
	Product0
}

// http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_temp4-8.shtml
type Product8 struct {
	Product0
	Time                              Time
	NumberOfIntervalTimeRanges        uint8 // length of last datatype
	TotalMissingDataValuesCount       uint32
	TimeRangeSpecification1           TimeRangeSpecification
	TimeRangeSpecification2           TimeRangeSpecification   // 59-70
	AdditionalTimeRangeSpecifications []TimeRangeSpecification // 71-n
}

type TimeRangeSpecification struct {
	StatisticalFieldCalculationProcess                     uint8  // 47
	IncrementBetweenSuccessiveFieldsType                   uint8  // 48
	IncrementBetweenSuccessiveFieldsRangeTimeUnitIndicator uint8  // 49
	StatististicalProcessTimeLength                        uint32 // 50-53
	IncrementBetweenSuccessiveFieldsTimeUnitIndicator      uint8  // 54
	TimeIncrementBetweenSuccessiveField                    uint32 // 55-58
}

type Surface struct {
	Type  uint8
	Scale uint8
	Value uint32
}
