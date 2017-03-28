package grib

import "fmt"

// translated from https://github.com/ArtemisiaSolutions/grib2js/blob/master/lib/grib/Grib2Spec.js

// Discipline (Code table 0.0)
func ReadDiscipline(value int) string {
	switch value {
	case 0:
		return "Meteorological"
	case 1:
		return "Hydrological"
	case 2:
		return "Land Surface"
	case 3:
		return "Space"
	case 4:
		return "Space Weather Products"
	case 10:
		return "Oceanographic"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown: ", value)
	}
}

// Master tables version number (Code table 1.0)
func ReadMasterTableVersionNumber(value int) string {
	switch value {
	case 0:
		return "Experimental"
	case 1:
		return "Initial"
	case 2:
		return "Version Implemented on 4 November 2003"
	case 3:
		return "Version Implemented on 2 November 2005"
	case 4:
		return "Version Implemented on 7 November 2007"
	case 5:
		return "Version Implemented on 4 November 2009"
	case 6:
		return "Version Implemented on 15 September 2010"
	case 7:
		return "Version Implemented on 4 May 2011"
	case 8:
		return "Version Implemented on 8 November 2011"
	case 9:
		return "Pre-operational to be implemented by next amendment"
	case 255:
		return "Master tables not used"
	default:
		return "Future"
	}
}

// Local tables version number (Code table 1.1)
func ReadLocalTableVersionNumber(value int) string {
	switch value {
	case 0:
		return "Local tables not used. Only table entries and templates from the current master table are valid."
	case 255:
		return "Missing"
	default:
		return "Number of local tables version used"
	}
}

// Significance of reference time (Code table 1.2)
func ReadReferenceTimeSignificance(value int) string {
	switch value {
	case 0:
		return "Analysis"
	case 1:
		return "Start of forecast"
	case 2:
		return "Verifying time of forecast"
	case 3:
		return "Observation time"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown: ", value)
	}
}

// Production status of data(Code table 1.3)
func ReadProductionStatus(value int) string {
	switch value {
	case 0:
		return "Operational"
	case 1:
		return "Operational test"
	case 2:
		return "Research"
	case 3:
		return "Re-analysis"
	case 4:
		return "THORPEX Interactive Grand Global Ensemble (TIGGE)"
	case 5:
		return "THORPEX Interactive Grand Global Ensemble (TIGGE) test"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown: ", value)
	}
}

// Type of data (Code table 1.4)
func ReadDataType(value int) string {
	switch value {
	case 0:
		return "Analysis"
	case 1:
		return "Forecast"
	case 2:
		return "Analysis and forecast"
	case 3:
		return "Control forecast"
	case 4:
		return "Perturbed forecast"
	case 5:
		return "Control and perturbed forecast"
	case 6:
		return "Processed satellite observations"
	case 7:
		return "Processed radar observations"
	case 8:
		return "Event Probability"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown: ", value)
	}
}

// Source of Grid Definition (Code table 3.0)
func ReadGridDefinitionSource(value int) string {
	switch value {
	case 0:
		return "Specified in Code table 3.1"
	case 1:
		return "Predetermined grid definition"
	case 255:
		return "A grid definition does not apply to this product"
	default:
		return fmt.Sprint("Unknown: ", value)
	}
}

// Grid Definition Template Number (Code table 3.1)
func ReadGridDefinitionTemplateNumber(value int) string {
	switch value {
	case 0:
		return "Latitude/longitude (See template 3.0)"
	case 1:
		return "Rotated latitude/longitude (See template 3.1)"
	case 2:
		return "Stretched latitude/longitude (See template 3.2)"
	case 3:
		return "Stretched and rotated latitude/longitude (See template 3.3)"
	case 10:
		return "Mercator (See template 3.10)"
	case 20:
		return "Polar stereographic (See template 3.20)"
	case 30:
		return "Lambert Conformal (See template 3.30)"
	case 31:
		return "Albers equal-area (See template 3.31)"
	case 40:
		return "Gaussian latitude/longitude (See template 3.40)"
	case 41:
		return "Rotated Gaussian latitude/longitude (See template 3.41)"
	case 42:
		return "Stretched Gaussian latitude/longitude (See template 3.42)"
	case 43:
		return "Stretched and rotated Gaussian latitude/longitude (See template 3.43)"
	case 44:
		return "Latitude/Longitude With Data-Sampling From A Higher Resolution Latitude/Longitude Source-Grid (See template 3.44)"
	case 50:
		return "Spherical harmonic coefficients (See template 3.50)"
	case 51:
		return "Rotated spherical harmonic coefficients (See template 3.51)"
	case 52:
		return "Stretched spherical harmonic coefficients (See template 3.52)"
	case 53:
		return "Stretched and rotated spherical harmonic coefficients (See template 3.53)"
	case 90:
		return "Space view perspective orthographic (See template 3.90)"
	case 100:
		return "Triangular grid based on an icosahedron (See template 3.100)"
	case 101:
		return "General Unstructured Grid (See template 3.101)"
	case 110:
		return "Equatorial azimuthal equidistant projection (See template 3.110)"
	case 120:
		return "Azimuth-range projection (See template 3.120)"
	case 130:
		return "Irregular Latitude/Longitude (See template 3.130)"
	case 204:
		return "Curvilinear Orthogonal Grids (See template 3.204)"
	case 1000:
		return "Cross-section grid, with points equally spaced on the horizontal (See template 3.1000)"
	case 1100:
		return "Hovmöller diagram grid, with points equally spaced on the horizontal (See template 3.1100)"
	case 1200:
		return "Time section grid (See template 3.1200)"
	case 32768:
		return "Rotated Latitude/Longitude (Arakawa Staggered E-Grid) (See template 3.32768)"
	case 32769:
		return "Rotated Latitude/Longitude (Arakawa Non-E Staggered Grid) (See template 3.32769)"
	case 65535:
		return "Missing"
	default:
		return fmt.Sprint("Unknown: ", value)
	}
}

// Shape of the earth (Code table 3.2)
func ReadEarthShape(value int) string {
	switch value {
	case 0:
		return "Earth assumed spherical with radius = 6,367,470.0m"
	case 1:
		return "Earth assumed spherical with radius specified by data producer"
	case 2:
		return "Earth assumed oblate spheroid with size as determined by IAU in 1965"
	case 3:
		return "Earth assumed oblate spheroid with major and minor axes defined by data producer"
	case 4:
		return "Earth assumed oblate spheroid as defined in IAG-GRS80 model"
	case 5:
		return "Earth assumed represented by WSG84"
	case 6:
		return "Earth assumed spherical with radius of 6,371,229.0m"
	case 7:
		return "Earth assumed oblate spheroid with major and minor axes specified (in m) by data producer"
	case 8:
		return "Earth model assumed spherical with radius 6,371,200 m, but the horizontal datum of the resulting Latitude/Longitude field is the WGS84 reference frame"
	case 9:
		return "Earth represented by the OSGB 1936 Datum, using the Airy_1830 Spheroid, the Greenwich meridian as 0 Longitude, the Newlyn datum as mean sea level, 0 height."
	case 10:
		return "Earth model assumed WGS84 with Corrected Geomagnetic Coordinates (Latitude and Longitude) defined by (Gustafsson et al., 1992)"
	case 11:
		return "Sun model assumed spherical with radius=695,990.000 m (Allen, C.W., 1976 Astrophysical Quantities (3rd Ed.; London: Athlone)). Stonyhurst latitude and longitude system with origin at the intersection of the solar central meridian as seen from Earth and the solar equator. See Thompson, W, Coordinate systems for solar image data, A&A 449, 791-803 (2006)"
	case 12:
		return "Sun model assumed spherical with radius=695,990.000 m (Allen, C.W., 1976 Astrophysical Quantities (3rd Ed.; London: Athlone)). Carrington latitude and longitude system that rotate with a side real period of 25.38 days. See Thompson, W, Coordinate systems for solar image data, A&A 449, 791-803 (2006)"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown: ", value)
	}
}

// Resolution and Component Flags (Flag table 3.3)
/*
func ReadResolutionComponentFlags(value int) (string) {
var flags = {}
flags.iDirectionIncrements = "Not given"
flags.jDirectionIncrements = "Not given"
flags.uAndvComponents = "Relative to easterly and northerly directions"
if((value & 4) === 4) (string) {
flags.iDirectionIncrements = "Given"
}
if((value & 8) === 8) (string) {
flags.jDirectionIncrements = "Given"
}
if((value & 16) === 16) (string) {
flags.uAndvComponents = "Relative to the defined grid"
}
return flags
}*/

// Scanning Mode (Flag table 3.4)
/*
func ReadScanMode(number, value int) (string) {
var flags = {}
flags.iDirection = "+i(+x)"
flags.jDirection = "-j(-y)"
flags.adjacentConsecutivePoints = "i(x)"
flags.rowsDirection = "All rows scan in the same direction"
if((value & 1) === 1) (string) {
flags.iDirection = "-i(-x)"
}
if((value & 2) === 2) (string) {
flags.jDirection = "+j(+y)"
}
if((value & 4) === 4) (string) {
flags.adjacentConsecutivePoints = "j(y)"
}
if((value & 8) === 8) (string) {
flags.rowsDirection = "Adjacent rows scan in the opposite direction"
}
return flags
}*/

/*
// Projection Centre (Flag table 3.5)
func ReadProjectionCentre(value int) (string) {
var flags = {}
flags.projectionPlane = "North Pole is on the projection plane"
flags.projectionType = "Only one projection center is used"
if((value & 1) === 1) (string) {
flags.projectionPlane = "South Pole is on the projection plane"
}
if((value & 2) === 2) (string) {
flags.projectionType = "Projection is bi-polar and symmetric"
}
return flags
}
*/

// Spectral data representation type (Code table 3.6)
func ReadSpectralDataRepresentationType(value int) string {
	switch value {
	case 1:
		return "Legendre functions"
	default:
		return fmt.Sprint("Unknown: ", value)
	}
}

// Spectral data representation mode (Code table 3.7)
func ReadSpectralDataRepresentationMode(value int) string {
	switch value {
	case 1:
		return "Complex numbers are stored as pairs of real numbers"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown: ", value)
	}
}

// Grid point position (code table 3.8)
func ReadGridPointPosition(value int) string {
	switch value {
	case 0:
		return "Grid points at triangle vertices"
	case 1:
		return "Grid points at centres of triangles"
	case 2:
		return "Grid points at midpoints of triangle sides"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown: ", value)
	}
}

/*
// Numbering order of diamonds as seen from the corresponding pole (Flag table 3.9)
func ReadDiamondsNumber(number, value int) (string) {
var flags = {}
if((value & 1) === 1) (string) {
flags.orientation = "Anti-clockwise (i.e., counter-clockwise) orientation"
} else {
flags.orientation = "Clockwise orientation"
}
return flags
}
*/

/*
// Scanning mode for one diamond (Flag table 3.10)
func ReadDiamondScanningMode(number, value int) (string) {
var flags = {}
flags.iDirection = "Points scan in +i direction, i.e. from pole to equator"
flags.jDirection = "Points scan in +j direction, i.e. from west to east"
flags.adjacentConsecutivePoints = "Adjacent points in i direction are consecutive"
if((value & 1) === 1) (string) {
flags.iDirection = "Points scan in -i direction, i.e. from equator to pole"
}
if((value & 2) === 2) (string) {
flags.jDirection = "Points scan in -j direction, i.e. from east to west"
}
if((value & 4) === 4) (string) {
flags.adjacentConsecutivePoints = "Adjacent points in j direction are consecutive"
}
return flags
}
*/

// Interpretation of list of numbers defining number of points (Code table 3.11)
func ReadListInterpretation(value int) string {
	switch value {
	case 0:
		return "There is no appended list"
	case 1:
		return "Numbers define number of points corresponding to full coordinate circles (i.e. parallels), coordinate values on each circle are multiple of the circle mesh, and extreme coordinate values given in grid definition (i.e. extreme longitudes) may not be reached in all rows"
	case 2:
		return "Numbers define number of points corresponding to coordinate lines delimited by extreme coordinate values given in grid definition (i.e. extreme longitudes) which are present in each row"
	case 3:
		return "Numbers define the actual latitudes for each row in the grid. The list of numbers are integer values of the valid latitudes in microdegrees (scale by 106) or in unit equal to the ratio of the basic angle and the subdivisions number for each row, in the same order as specified in the \"scanning mode flag\" (bit no. 2)"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown: ", value)
	}
}

// Physical meaning of vertical coordinate (Code table 3.15)
func ReadVerticalCoordinatePhysicalMeaning(value int) string {
	switch value {
	case 20:
		return "Temperature (K)"
	case 100:
		return "Pressure (Pa)"
	case 101:
		return "Pressure deviation from mean sea level (Pa)"
	case 102:
		return "Altitude above mean sea level (m)"
	case 103:
		return "Height above ground (m)"
	case 104:
		return "Sigma coordinate"
	case 105:
		return "Hybrid coordinate"
	case 106:
		return "Depth below land surface (m)"
	case 107:
		return "Potential temperature theta (K)"
	case 108:
		return "Pressure deviation from ground level (Pa)"
	case 109:
		return "Potential vorticity (K m-2 kg-1 s-1)"
	case 110:
		return "Geometrical height (m)"
	case 111:
		return "Eta coordinate"
	case 112:
		return "Geopotential height (gpm)"
	case 113:
		return "Logarithmic hybrid coordinate"
	case 160:
		return "Depth below sea level (m)"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown: ", value)
	}
}

// Type of horizontal line (Code table 3.20)
func ReadHorizontalLineType(value int) string {
	switch value {
	case 0:
		return "There is no appended list"
	case 1:
		return "Numbers define number of points corresponding to full coordinate circles (i.e. parallels).  Coordinate values on each circle are multiple of the circle mesh, and extreme coordinate values given in grid definition may not be reached in all rows."
	case 2:
		return "Numbers define number of points corresponding to coordinate lines delimited by extreme coordinate values given in grid definition which are present in each row."
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown: ", value)
	}
}

// Vertical dimension coordinate values definition (Code table 3.21)
func ReadVerticalDimensionCoordinateValuesDefinition(value int) string {
	switch value {
	case 0:
		return "Explicit coordinate values set"
	case 1:
		return "Linear coordinates"
	case 11:
		return "Geometric coordinates"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown: ", value)
	}
}

// Product Definition Template Number (Code table 4.0)
func ReadProductDefinitionTemplateNumber(value int) string {
	switch value {
	case 0:
		return "Analysis or forecast at a horizontal level or in a horizontal layer at a point in time (template 4.0)"
	case 1:
		return "Individual ensemble forecast, control and perturbed, at a horizontal level or in a horizontal layer at a point in time (template 4.1)"
	case 2:
		return "Derived forecast based on all ensemble members at a horizontal level or in a horizontal layer at a point in time (template 4.2)"
	case 3:
		return "Derived forecasts based on a cluster of ensemble members over a rectangular area at a horizontal level or in a horizontal layer at a point in time (template 4.3)"
	case 4:
		return "Derived forecasts based on a cluster of ensemble members over a circular area at a horizontal level or in a horizontal layer at a point in time (template 4.4)"
	case 5:
		return "Probability forecasts at a horizontal level or in a horizontal layer at a point in time (template 4.5)"
	case 6:
		return "Percentile forecasts at a horizontal level or in a horizontal layer at a point in time (template 4.6)"
	case 7:
		return "Analysis or forecast error at a horizontal level or in a horizontal layer at a point in time (template 4.7)"
	case 8:
		return "Average, accumulation, extreme values or other statistically processed values at a horizontal level or in a horizontal layer in a continuous or non-continuous time interval (template 4.8)"
	case 9:
		return "Probability forecasts at a horizontal level or in a horizontal layer in a continuous or non-continuous time interval (template 4.9)"
	case 10:
		return "Percentile forecasts at a horizontal level or in a horizontal layer in a continuous or non-continuous time interval (template 4.10)"
	case 11:
		return "Individual ensemble forecast, control and perturbed, at a horizontal level or in a horizontal layer, in a continuous or non-continuous interval (template 4.11)"
	case 12:
		return "Derived forecasts based in all ensemble members at a horizontal level or in a horizontal layer, in a continuous or non-continuous interval (template 4.12)"
	case 13:
		return "Derived forecasts based on a cluster of ensemble members over a rectangular area, at a horizontal level or in a horizontal layer, in a continuous or non-continuous interval (template 4.13)"
	case 14:
		return "Derived forecasts based on a cluster of ensemble members over a circular area, at a horizontal level or in a horizontal layer, in a continuous or non-continuous interval (template 4.14)"
	case 15:
		return "Average, accumulation, extreme values or other statistically-processed values over a spatial area at a horizontal level or in a horizontal layer at a point in time. (template 4.15)"
	case 20:
		return "Radar product (template 4.20)"
	case 30:
		return "Satellite product (template 4.30) (deprecated)"
	case 31:
		return "Satellite product (template 4.31)"
	case 32:
		return "Analysis or forecast at a horizontal level or in a horizontal layer at a point in time for simulate (synthetic) satellite data (see Template 4.32)"
	case 40:
		return "Analysis or forecast at a horizontal level or in a horizontal layer at a point in time for atmospheric chemical constituents.  (see Template 4.40)"
	case 41:
		return "Individual ensemble forecast, control and perturbed, at a horizontal level or in a horizontal layer at a point in time for atmospheric chemical constituents.  (see Template 4.41)"
	case 42:
		return "Average, accumulation, and/or extreme values or other statistically processed values at a horizontal level or in a horizontal layer in a continuous or non-continuous time interval for atmospheric chemical constituents.  (see Template 4.42)"
	case 43:
		return "Individual ensemble forecast, control and perturbed, at a horizontal level or in a horizontal layer, in a continuous or non-continuous time interval for atmospheric chemical constituents.  (see Template 4.43)"
	case 44:
		return "Analysis or forecast at a horizontal level or in a horizontal layer at a point in time for aerosol.  (see Template 4.44)"
	case 45:
		return "Individual ensemble forecast, control and perturbed, at a horizontal level or in a horizontal layer, in a continuous or non-continuous time interval for aerosol.  (see Template 4.45)"
	case 46:
		return "Average, accumulation, and/or extreme values or other statistically processed values at a horizontal level or in a horizontal layer in a continuous or non-continuous time interval for aerosol.  (see Template 4.46)"
	case 47:
		return "Individual ensemble forecast, control and perturbed, at a horizontal level or in a horizontal layer, in a continuous or non-continuous time interval for aerosol.  (see Template 4.47)"
	case 48:
		return "Analysis or forecast at a horizontal level or in a horizontal layer at a point in time for aerosol.  (see Template 4.48)"
	case 50:
		return "Analysis or forecast of a multi component parameter or matrix element at error at a point in time.  (see Template 4.50)  "
	case 51:
		return "Categorical forecast at a horizontal level or in a horizontal layer at a point in time.  (see Template 4.51)"
	case 52:
		return "Analysis or forecast of Wave Parameters at the Sea Surface at a point in time.  (see Template 4.52)"
	case 91:
		return "Categorical forecast at a horizontal level or in a horizontal layer in a continuous or non-continuous time interval.  (see Template 4.91)"
	case 254:
		return "CCITT IA5 character string  (see Template 4.254)"
	case 1000:
		return "Cross section of analysis and forecast at a point in time  (see Template 4.1000)"
	case 1001:
		return "Cross section of averaged or otherwise statistically processed analysis or forecast over a range of time  (see Template 4.1001)"
	case 1002:
		return "Cross section of analysis and forecast, averaged or or otherwise statistically processed  (see Template 4.1002)"
	case 1100:
		return "Hovmöller-type grid with no averaging or other statistical processing  (see Template 4.1100)"
	case 1101:
		return "Hovmöller-type grid with averaging or other statistical processing  (see Template 4.1101)"
	case 65535:
		return "Missing"
	default:
		return fmt.Sprint("Unknown: ", value)
	}
}

// Category of parameters by product discipline (Code table 4.1)
func ReadProductDisciplineParameters(discipline, category int) string {
	switch discipline {
	//Meteorological Products
	case 0:
		switch category {
		case 0:
			return "Temperature (see Table 4.2-0-0)"
		case 1:
			return "Moisture (see Table 4.2-0-1)"
		case 2:
			return "Momentum (see Table 4.2-0-2)"
		case 3:
			return "Mass (see Table 4.2-0-3)"
		case 4:
			return "Short-wave Radiation (see Table 4.2-0-4)"
		case 5:
			return "Long-wave Radiation (see Table 4.2-0-5)"
		case 6:
			return "Cloud (see Table 4.2-0-6)"
		case 7:
			return "Thermodynamic Stability indices (see Table 4.2-0-7)"
		case 8:
			return "Kinematic stability indices"
		case 9:
			return "Temperature probabilities"
		case 10:
			return "Moisture probabilities"
		case 11:
			return "Momentum probabilities"
		case 12:
			return "Mass probabilities"
		case 13:
			return "Aerosols (see Table 4.2-0-13)"
		case 14:
			return "Trace gases(e.g Ozone, CO2) (see Table 4.2-0-14)"
		case 15:
			return "Radar (see Table 4.2-0-15)"
		case 16:
			return "Forecast Radar Imagery (see Table 4.2-0-16)"
		case 17:
			return "Electro-dynamics (see Table 4.2-0-17)"
		case 18:
			return "Nuclear/radiology (see Table 4.2-0-18)"
		case 19:
			return "Physical atmospheric properties (see Table 4.2-0-19)"
		case 20:
			return "Atmospheric chemical Constituents (see Table 4.2-0-20)"
		case 190:
			return "CCITT IA5 string (see Table 4.2-0-190)"
		case 191:
			return "Miscellaneous (see Table 4.2-0-191)"
		case 192:
			return "Covariance (see Table 4.2-0-192)"
		case 255:
			return "Missing"
		default:
			return fmt.Sprint("Unknown ", category)
		}
	//Hydrological Products
	case 1:
		switch category {
		case 0:
			return "Hydrology basic products (see Table 4.2-1-0)"
		case 1:
			return "Hydrology probabilities (see Table 4.2-1-1)"
		case 2:
			return "Inland water and sediment properties (see Table 4.2-1-2)"
		case 255:
			return "Missing"
		default:
			return fmt.Sprint("Unknown ", category)
		}
	//Land surface products
	case 2:
		switch category {
		case 0:
			return "Vegetation/Biomass (see Table 4.2-2-0)"
		case 1:
			return "Agri-/aquacultural Special Products (see Table 4.2-2-1)"
		case 2:
			return "Transportation-related Products (see Table 4.2-2-2)"
		case 3:
			return "Soil Products (see Table 4.2-2-3)"
		case 4:
			return "Fire Weather (see Table 4.2-2-4)"
		case 255:
			return "Missing"
		default:
			return fmt.Sprint("Unknown ", category)
		}
	//Space products
	case 3:
		switch category {
		case 0:
			return "Image format products (see Table 4.2-3-0)"
		case 1:
			return "Quantitative products (see Table 4.2-3-1)"
		case 192:
			return "Forecast Satellite Imagery (see Table 4.2-3-192)"
		case 255:
			return "Missing"
		default:
			return fmt.Sprint("Unknown ", category)
		}
	//Space weather products
	case 4:
		switch category {
		case 0:
			return "Temperature (see Table 4.2-4-0)"
		case 1:
			return "Momentum (see Table 4.2-4-1)"
		case 2:
			return "Charged Particle Mass and Number (see Table 4.2-4-2)"
		case 3:
			return "Electric and Magnetic Fields (see Table 4.2-4-3)"
		case 4:
			return "Energetic Particles (see Table 4.2-4-4)"
		case 5:
			return "Waves (see Table 4.2-4-5)"
		case 6:
			return "Solar Electromagnetic Emissions (see Table 4.2-4-6)"
		case 7:
			return "Terrestrial Electromagnetic Emissions (see Table 4.2-4-7)"
		case 8:
			return "Imagery (see Table 4.2-4-8)"
		case 9:
			return "Ion-Neutral Coupling (see Table 4.2-4-9)"
		case 255:
			return "Missing"
		default:
			return fmt.Sprint("Unknown ", category)
		}
	//Oceanographic products
	case 10:
		switch category {
		case 0:
			return "Waves (see Table 4.2-10-0)"
		case 1:
			return "Currents (see Table 4.2-10-1)"
		case 2:
			return "Ice (see Table 4.2-10-2)"
		case 3:
			return "Surface Properties (see Table 4.2-10-3)"
		case 4:
			return "Sub-surface Properties (see Table 4.2-10-4)"
		case 191:
			return "Miscellaneous (see Table 4.2-10-191)"
		case 255:
			return "Missing"
		default:
			return fmt.Sprint("Unknown ", category)
		}
	default:
		return fmt.Sprint("Unknown ", discipline)
	}
}

// Parameter number by product discipline and parameter category (code table 4.2)
func ReadProductDisciplineCategoryParameters(discipline, category, number int) string {
	switch discipline {
	//Product Discipline 0: Meteorological products
	case 0:
		switch category {
		//Parameter Category 0: Temperature
		case 0:
			switch number {
			case 0:
				return "Temperature (K)"
			case 1:
				return "Virtual temperature (K)"
			case 2:
				return "Potential temperature (K)"
			case 3:
				return "Pseudo-adiabatic potential temperature or equivalent potential temperature (K)"
			case 4:
				return "Maximum temperature (K)"
			case 5:
				return "Minimum temperature (K)"
			case 6:
				return "Dew point temperature (K)"
			case 7:
				return "Dew point depression(or deficit) (K)"
			case 8:
				return "Lapse rate (K m-1)"
			case 9:
				return "Temperature anomaly (K)"
			case 10:
				return "Latent heat net flux (W m-2)"
			case 11:
				return "Sensible heat net flux (W m-2)"
			case 12:
				return "Heat index (K)"
			case 13:
				return "Wind chill factor (K)"
			case 14:
				return "Minimum dew point depression (K)"
			case 15:
				return "Virtual potential temperature (K)"
			case 16:
				return "Snow Phase Change Heat Flux (W m-2)"
			case 17:
				return "Skin Temperature (K)"
			case 18:
				return "Snow Temperature (top of snow) (K)"
			case 19:
				return "Turbulent Transfer Coefficient for Heat (Numeric)"
			case 20:
				return "Turbulent Diffusion Coefficient for Heat (m2 s-1)"
			case 192:
				return "Snow Phase Change Heat Flux (W m-2)"
			case 193:
				return "Temperature Tendency by All Radiation (K s-1)"
			case 194:
				return "Relative Error Variance"
			case 195:
				return "Large Scale Condensate Heating Rate (K s-1)"
			case 196:
				return "Deep Convective Heating Rate (K s-1)"
			case 197:
				return "Total Downward Heat Flux at Surface (W m-2)"
			case 198:
				return "Temperature Tendency by All Physics (K s-1)"
			case 199:
				return "Temperature Tendency by Non-radiation Physics (K s-1)"
			case 200:
				return "Standard Dev. of IR Temp. over 1x1 deg. area (K)"
			case 201:
				return "Shallow Convective Heating Rate (K s-1)"
			case 202:
				return "Vertical Diffusion Heating rate (K s-1)"
			case 203:
				return "Potential Temperature at Top of Viscous Sublayer (K)"
			case 204:
				return "Tropical Cyclone Heat Potential (J m-2 K-1)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 1: Moisture
		case 1:
			switch number {
			case 0:
				return "Specific humidity (kg kg-1)"
			case 1:
				return "Relative humidity (%)"
			case 2:
				return "Humidity mixing ration (kg kg-1)"
			case 3:
				return "Precipitable water (kg m-2)"
			case 4:
				return "Vapor pressure (Pa)"
			case 5:
				return "Saturation deficit (Pa)"
			case 6:
				return "Evaporation (kg m-2)"
			case 7:
				return "Precipitation rate (kg m-2 s-1)"
			case 8:
				return "Total precipitation (kg m-2)"
			case 9:
				return "Large scale precipitation(non-convective) (kg m-2)"
			case 10:
				return "Convective precipitation (kg m-2)"
			case 11:
				return "Snow depth (m)"
			case 12:
				return "Snowfall rate water equivalent (kg m-2 s-1)"
			case 13:
				return "Water equivalent of accumulated snow depth (kg m-2)"
			case 14:
				return "Convective snow (kg m-2)"
			case 15:
				return "Large scale know (kg m-2)"
			case 16:
				return "Snow melt(kg m-2)"
			case 17:
				return "Snow age (day)"
			case 18:
				return "Absolute humidity (kg m-3)"
			case 19:
				return "Precipitation type (code table (4.201))"
			case 20:
				return "Integrated liquid water (kg m-2)"
			case 21:
				return "Condensate (kg kg-1)"
			case 22:
				return "Cloud mixing ratio (kg kg-1)"
			case 23:
				return "Ice water mixing ratio (kg kg-1)"
			case 24:
				return "Rain mixing ratio (kg kg-1)"
			case 25:
				return "Snow mixing ratio (kg kg-1)"
			case 26:
				return "Horizontal moisture convergence (kg kg-1 s-1)"
			case 27:
				return "Maximum relative humidity (%)"
			case 28:
				return "Maximum absolute humidity (kg m-3)"
			case 29:
				return "Total snowfall (m)"
			case 30:
				return "Precipitable water category (code table(4.202))"
			case 31:
				return "Hail (m)"
			case 32:
				return "Graupel(snow pellets) (kg kg-1)"
			case 33:
				return "Categorical Rain	(Code table 4.222)"
			case 34:
				return "Categorical Freezing Rain (Code table 4.222)"
			case 35:
				return "Categorical Ice Pellets	(Code table 4.222)"
			case 36:
				return "Categorical Snow	(Code table 4.222)"
			case 37:
				return "Convective Precipitation Rate (kg m-2 s-1)"
			case 38:
				return "Horizontal Moisture Divergence	(kg kg-1 s-1)"
			case 39:
				return "Percent frozen precipitation	(%)"
			case 40:
				return "Potential Evaporation (kg m-2)"
			case 41:
				return "Potential Evaporation Rate	(W m-2)"
			case 42:
				return "Snow Cover (%)"
			case 43:
				return "Rain Fraction of Total Cloud Water (Proportion)"
			case 44:
				return "Rime Factor	(Numeric)"
			case 45:
				return "Total Column Integrated Rain	(kg m-2)"
			case 46:
				return "Total Column Integrated Snow	(kg m-2)"
			case 47:
				return "Large Scale Water Precipitation (Non-Convective)	(kg m-2)"
			case 48:
				return "Convective Water Precipitation	(kg m-2)"
			case 49:
				return "Total Water Precipitation	(kg m-2)"
			case 50:
				return "Total Snow Precipitation	(kg m-2)"
			case 51:
				return "Total Column Water (Vertically integrated total water (vapour+cloud water/ice)	(kg m-2)"
			case 52:
				return "Total Precipitation Rate	(kg m-2 s-1)"
			case 53:
				return "Total Snowfall Rate Water Equivalent	(kg m-2 s-1)"
			case 54:
				return "Large Scale Precipitation Rate	(kg m-2 s-1)"
			case 55:
				return "Convective Snowfall Rate Water Equivalent	(kg m-2 s-1)"
			case 56:
				return "Large Scale Snowfall Rate Water Equivalent	(kg m-2 s-1)"
			case 57:
				return "Total Snowfall Rate	(m s-1)"
			case 58:
				return "Convective Snowfall Rate	(m s-1)"
			case 59:
				return "Large Scale Snowfall Rate (m s-1)"
			case 60:
				return "Snow Depth Water Equivalent	(kg m-2)"
			case 61:
				return "Snow Density	(kg m-3)"
			case 62:
				return "Snow Evaporation	(kg m-2)"
			case 64:
				return "Total Column Integrated Water Vapour	(kg m-2)"
			case 65:
				return "Rain Precipitation Rate	(kg m-2 s-1)"
			case 66:
				return "Snow Precipitation Rate	(kg m-2 s-1)"
			case 67:
				return "Freezing Rain Precipitation Rate	(kg m-2 s-1)"
			case 68:
				return "Ice Pellets Precipitation Rate (kg m-2 s-1)"
			case 69:
				return "Total Column Integrate Cloud Water (kg m-2)"
			case 70:
				return "Total Column Integrate Cloud Ice	(kg m-2)"
			case 71:
				return "Hail Mixing Ratio (kg kg-1)"
			case 72:
				return "Total Column Integrate Hail	(kg m-2)"
			case 73:
				return "Hail Prepitation Rate (kg m-2 s-1)"
			case 74:
				return "Total Column Integrate Graupel (kg m-2)"
			case 75:
				return "Graupel (Snow Pellets) Prepitation Rate (kg m-2 s-1)"
			case 76:
				return "Convective Rain Rate (kg m-2 s-1)"
			case 77:
				return "Large Scale Rain Rate (kg m-2 s-1)"
			case 78:
				return "Total Column Integrate Water (All components including precipitation) (kg m-2)"
			case 79:
				return "Evaporation Rate (kg m-2 s-1)"
			case 80:
				return "Total Condensate (kg kg-1)"
			case 81:
				return "Total Column-Integrate Condensate (kg m-2)"
			case 82:
				return "Cloud Ice Mixing Ratio (kg kg-1)"
			case 83:
				return "Specific Cloud Liquid Water Content	(kg kg-1)"
			case 84:
				return "Specific Cloud Ice Water Content	(kg kg-1)"
			case 85:
				return "Specific Rain Water Content	(kg kg-1)"
			case 86:
				return "Specific Snow Water Content	(kg kg-1)"
			case 90:
				return "Total Kinematic Moisture Flux (kg kg-1 m s-1)"
			case 91:
				return "U-component (zonal) Kinematic Moisture Flux	(kg kg-1 m s-1)"
			case 92:
				return "V-component (meridional) Kinematic Moisture Flux	(kg kg-1 m s-1)"
			case 192:
				return "Categorical Rain	(Code table 4.222)"
			case 193:
				return "Categorical Freezing Rain (Code table 4.222)"
			case 194:
				return "Categorical Ice Pellets	(Code table 4.222)"
			case 195:
				return "Categorical Snow	(Code table 4.222)"
			case 196:
				return "Convective Precipitation Rate (kg m-2 s-1)"
			case 197:
				return "Horizontal Moisture Divergence (kg kg-1 s-1)"
			case 198:
				return "Minimum Relative Humidity (%)"
			case 199:
				return "Potential Evaporation (kg m-2)"
			case 200:
				return "Potential Evaporation Rate (W m-2)"
			case 201:
				return "Snow Cover	(%)"
			case 202:
				return "Rain Fraction of Total Liquid Water (non-dim)"
			case 203:
				return "Rime Factor	(non-dim)"
			case 204:
				return "Total Column Integrated Rain	(kg m-2)"
			case 205:
				return "Total Column Integrated Snow	(kg m-2)"
			case 206:
				return "Total Icing Potential Diagnostic	(non-dim)"
			case 207:
				return "Number concentration for ice particles	(non-dim)"
			case 208:
				return "Snow temperature	(K)"
			case 209:
				return "Total column-integrated supercooled liquid water	(kg m-2)"
			case 210:
				return "Total column-integrated melting ice	(kg m-2)"
			case 211:
				return "Evaporation - Precipitation	(cm/day)"
			case 212:
				return "Sublimation (evaporation from snow)	(W m-2)"
			case 213:
				return "Deep Convective Moistening Rate	(kg kg-1 s-1)"
			case 214:
				return "Shallow Convective Moistening Rate (kg kg-1 s-1)"
			case 215:
				return "Vertical Diffusion Moistening Rate	(kg kg-1 s-1)"
			case 216:
				return "Condensation Pressure of Parcali Lifted From Indicate Surface (Pa)"
			case 217:
				return "Large scale moistening rate	(kg kg-1 s-1)"
			case 218:
				return "Specific humidity at top of viscous sublayer	(kg kg-1)"
			case 219:
				return "Maximum specific humidity at 2m	(kg kg-1)"
			case 220:
				return "Minimum specific humidity at 2m	(kg kg-1)"
			case 221:
				return "Liquid precipitation (rainfall)	(kg m-2)"
			case 222:
				return "Snow temperature, depth-avg	(K)"
			case 223:
				return "Total precipitation (nearest grid point)	(kg m-2)"
			case 224:
				return "Convective precipitation (nearest grid point)	(kg m-2)"
			case 225:
				return "Freezing Rain (kg m-2)"
			case 226:
				return "Predominant Weather (Numeric)"
			case 227:
				return "Frozen Rain (kg m-2)"
			case 241:
				return "Total Snow (kg m-2)"
			case 242:
				return "Relative Humidity with Respect to Precipitable Water (%)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 2: Momentum
		case 2:
			switch number {
			case 0:
				return "Wind direction(from which blowing) (deg true)"
			case 1:
				return "Wind speed (m s-1)"
			case 2:
				return "u-component of wind m s-1"
			case 3:
				return "v-component of wind m s-1"
			case 4:
				return "Stream function (m2 s-1)"
			case 5:
				return "Velocity potential (m2 s-1)"
			case 6:
				return "Montgomery streal function (m2 s-2)"
			case 7:
				return "Sigma coordinate vertical velocity (s-1)"
			case 8:
				return "Vertical velocity(pressure) (Pa s-1)"
			case 9:
				return "Vertical velocity(geometric) (m s-1)"
			case 10:
				return "Absolute vorticity (s-1)"
			case 11:
				return "Absolute divergence (s-1)"
			case 12:
				return "Relative vorticity (s-1)"
			case 13:
				return "Relative divergence (s-1)"
			case 14:
				return "Potential vorticity (K m2 kg-1 s-1)"
			case 15:
				return "Vertical u-component shear (s-1)"
			case 16:
				return "Vertical v-component shear (s-1)"
			case 17:
				return "Momentum flux, u-component (N m-2)"
			case 18:
				return "Momentum flux, v-component (N m-2)"
			case 19:
				return "Wind mixing energy (J)"
			case 20:
				return "Boundary layer dissipation (W m-2)"
			case 21:
				return "Maximum wind speed (m s-1)"
			case 22:
				return "Wind speed(gust) (m s-1)"
			case 23:
				return "u-component of wind(gust) (m s-1)"
			case 24:
				return "v-component of wind(gust) (m s-1)"
			case 25:
				return "Vertical Speed Shear	(s-1)"
			case 26:
				return "Horizontal Momentum Flux	(N m-2)"
			case 27:
				return "U-Component Storm Motion	(m s-1)"
			case 28:
				return "V-Component Storm Motion	(m s-1)"
			case 29:
				return "Drag Coefficient	(Numeric)"
			case 30:
				return "Frictional Velocity	(m s-1)"
			case 31:
				return "Turbulent Diffusion Coefficient for Momentum (m2 s-1)"
			case 32:
				return "Eta Coordinate Vertical Velocity	(s-1)"
			case 33:
				return "Wind Fetch (m)"
			case 192:
				return "Vertical Speed Shear	(s-1)"
			case 193:
				return "Horizontal Momentum Flux	(N m-2)"
			case 194:
				return "U-Component Storm Motion	(m s-1)"
			case 195:
				return "V-Component Storm Motion	(m s-1)"
			case 196:
				return "Drag Coefficient	(Numeric)"
			case 197:
				return "Frictional Velocity	(m s-1)"
			case 198:
				return "Latitude of U Wind Component of Velocity	(deg)"
			case 199:
				return "Longitude of U Wind Component of Velocity (deg)"
			case 200:
				return "Latitude of V Wind Component of Velocity	(deg)"
			case 201:
				return "Longitude of V Wind Component of Velocity (deg)"
			case 202:
				return "Longitude of Presure Point (deg)"
			case 203:
				return "Latitude of Presure Point (deg)"
			case 204:
				return "Vertical Eddy Diffusivity Heat exchange	(m2 s-1)"
			case 205:
				return "Covariance between Meridional and Zonal Components of the wind (m2 s-2)"
			case 206:
				return "Covariance between Temperature and Zonal Components of the wind (K*m s-1)"
			case 207:
				return "Covariance between Temperature and Meridional Components of the wind (K*m s-1)"
			case 208:
				return "Vertical Diffusion Zonal Acceleration (m s-2)"
			case 209:
				return "Vertical Diffusion Meridional Acceleration (m s-2)"
			case 210:
				return "Gravity wave drag zonal acceleration (m s-2)"
			case 211:
				return "Gravity wave drag meridional acceleration (m s-2)"
			case 212:
				return "Convective zonal momentum mixing acceleration (m s-2)"
			case 213:
				return "Convective meridional momentum mixing acceleration (m s-2)"
			case 214:
				return "Tendency of vertical velocity (m s-2)"
			case 215:
				return "Omega (Dp/Dt) divide by density	(K)"
			case 216:
				return "Convective Gravity wave drag zonal acceleration (m s-2)"
			case 217:
				return "Convective Gravity wave drag meridional acceleration (m s-2)"
			case 218:
				return "Velocity Point Model Surface"
			case 219:
				return "Potential Vorticity (Mass-Weighted)	(1/s/m)"
			case 220:
				return "Hourly Maximum of Upward Vertical Velocity in the lowest 400hPa (m s-1)"
			case 221:
				return "Hourly Maximum of Downward Vertical Velocity in the lowest 400hPa (m s-1)"
			case 222:
				return "U Component of Hourly Maximum 10m Wind Speed	(m s-1)"
			case 223:
				return "V Component of Hourly Maximum 10m Wind Speed	(m s-1)"
			case 224:
				return "Ventilation Rate	(m2 s-1)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 3: Mass
		case 3:
			switch number {
			case 0:
				return "Pressure (Pa)"
			case 1:
				return "Pressure reduced to MSL (Pa)"
			case 2:
				return "Pressure tendency (Pa s-1)"
			case 3:
				return "ICAO Standard Atmosphere Reference Height (m)"
			case 4:
				return "Geopotential (m2 s-2)"
			case 5:
				return "Geopotential height (gpm)"
			case 6:
				return "Geometric height (m)"
			case 7:
				return "Standard deviation of height (m)"
			case 8:
				return "Pressure anomaly (Pa)"
			case 9:
				return "Geopotential height anomaly (gpm)"
			case 10:
				return "Density (kg m-3)"
			case 11:
				return "Altimeter setting (Pa)"
			case 12:
				return "Thickness (m)"
			case 13:
				return "Pressure altitude (m)"
			case 14:
				return "Density altitude (m)"
			case 15:
				return "5-Wave Geopotential Height (gpm)"
			case 16:
				return "Zonal Flux of Gravity Wave Stress (N m-2)"
			case 17:
				return "Meridional Flux of Gravity Wave Stress (N m-2)"
			case 18:
				return "Planetary Boundary Layer Height	(m)"
			case 19:
				return "5-Wave Geopotential Height Anomaly (gpm)"
			case 20:
				return "Standard Deviation of Sub-Grid Scale Orography (m)"
			case 21:
				return "Angle of Sub-Grid Scale Orography (rad)"
			case 22:
				return "Slope of Sub-Grid Scale Orography (Numeric)"
			case 23:
				return "Gravity Wave Dissipation	(W m-2)"
			case 24:
				return "Anisotropy of Sub-Grid Scale Orography (Numeric)"
			case 25:
				return "Natural Logarithm of Pressure in Pa	(Numeric)"
			case 192:
				return "MSLP (Eta model reduction)	(Pa)"
			case 193:
				return "5-Wave Geopotential Height	(gpm)"
			case 194:
				return "Zonal Flux of Gravity Wave Stress (N m-2)"
			case 195:
				return "Meridional Flux of Gravity Wave Stress (N m-2)"
			case 196:
				return "Planetary Boundary Layer Height	(m)"
			case 197:
				return "5-Wave Geopotential Height Anomaly (gpm)"
			case 198:
				return "MSLP (MAPS System Reduction)	(Pa)"
			case 199:
				return "3-hr pressure tendency (Std. Atmos. Reduction)	(Pa s-1)"
			case 200:
				return "Pressure of level from which parcel was lifted	(Pa)"
			case 201:
				return "X-gradient of Log Pressure	(m-1)"
			case 202:
				return "Y-gradient of Log Pressure	(m-1)"
			case 203:
				return "X-gradient of Height	(m-1)"
			case 204:
				return "Y-gradient of Height	(m-1)"
			case 205:
				return "Layer Thickness	(m)"
			case 206:
				return "Natural Log of Surface Pressure	(ln (kPa))"
			case 207:
				return "Convective updraft mass flux (kg m-2 s-1)"
			case 208:
				return "Convective downdraft mass flux (kg m-2 s-1)"
			case 209:
				return "Convective detrainment mass flux	(kg m-2 s-1)"
			case 210:
				return "Mass Point Model Surface"
			case 211:
				return "Geopotential Height (nearest grid point)	(gpm)"
			case 212:
				return "Pressure (nearest grid point) (Pa)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 4: Short-wave Radiation
		case 4:
			switch number {
			case 0:
				return "Net long wave radiation flux(surface) (W m-2)"
			case 1:
				return "Net long wave radiation flux(top of atmosphere) (W m-2)"
			case 2:
				return "Short wave radiation flux (W m-2)"
			case 3:
				return "Global radiation flux (W m-2)"
			case 4:
				return "Brightness temperature (K)"
			case 5:
				return "Radiance(with respect to wave number) (W m-3 sr-1)"
			case 6:
				return "Radiance(with respect to wave length) (W m-3 sr-1)"
			case 7:
				return "Downward Short-Wave Radiation Flux (W m-2)"
			case 8:
				return "Upward Short-Wave Radiation Flux	(W m-2)"
			case 9:
				return "Net Short Wave Radiation Flux (W m-2)"
			case 10:
				return "Photosynthetically Active Radiation	(W m-2)"
			case 11:
				return "Net Short-Wave Radiation Flux, Clear Sky	(W m-2)"
			case 12:
				return "Downward UV Radiation (W m-2)"
			case 50:
				return "UV Index (Under Clear Sky) (Numeric)"
			case 51:
				return "UV Index (W m-2)"
			case 192:
				return "Downward Short-Wave Radiation Flux (W m-2)"
			case 193:
				return "Upward Short-Wave Radiation Flux	(W m-2)"
			case 194:
				return "UV-B Downward Solar Flux	(W m-2)"
			case 195:
				return "Clear sky UV-B Downward Solar Flux (W m-2)"
			case 196:
				return "Clear Sky Downward Solar Flux (W m-2)"
			case 197:
				return "Solar Radiative Heating Rate	(K s-1)"
			case 198:
				return "Clear Sky Upward Solar Flux	(W m-2)"
			case 199:
				return "Cloud Forcing Net Solar Flux	(W m-2)"
			case 200:
				return "Visible Beam Downward Solar Flux	(W m-2)"
			case 201:
				return "Visible Diffuse Downward Solar Flux	(W m-2)"
			case 202:
				return "Near IR Beam Downward Solar Flux	(W m-2)"
			case 203:
				return "Near IR Diffuse Downward Solar Flux	(W m-2)"
			case 204:
				return "Downward Total Radiation Flux (W m-2)"
			case 205:
				return "Upward Total Radiation Flux	(W m-2)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 5: Long-wave Radiation
		case 5:
			switch number {
			case 0:
				return "Net long wave radiation flux(surface) (W m-2)"
			case 1:
				return "Net long wave radiation flux(top of atmosphere) (W m-2)"
			case 2:
				return "Long wave radiation flux (W m-2)"
			case 3:
				return "Downward Long-Wave Rad. Flux (W m-2)"
			case 4:
				return "Upward Long-Wave Rad. Flux (W m-2)"
			case 5:
				return "Net Long-Wave Radiation Flux (W m-2)"
			case 6:
				return "Net Long-Wave Radiation Flux, Clear Sky (W m-2)"
			case 192:
				return "Downward Long-Wave Rad. Flux (W m-2)"
			case 193:
				return "Upward Long-Wave Rad. Flux (W m-2)"
			case 194:
				return "Long-Wave Radiative Heating Rate (K s-1)"
			case 195:
				return "Clear Sky Upward Long Wave Flux (W m-2)"
			case 196:
				return "Clear Sky Downward Long Wave Flux (W m-2)"
			case 197:
				return "Cloud Forcing Net Long Wave Flux (W m-2)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 6: Cloud
		case 6:
			switch number {
			case 0:
				return "Cloud ice (kg m-2)"
			case 1:
				return "Total cloud cover (%)"
			case 2:
				return "Convective cloud cover (%)"
			case 3:
				return "Low cloud cover (%)"
			case 4:
				return "Medium cloud cover (%)"
			case 5:
				return "High cloud cover (%)"
			case 6:
				return "Cloud water (kg m-2)"
			case 7:
				return "Cloud amount (%)"
			case 8:
				return "Cloud type (code table (4.203))"
			case 9:
				return "Thunderstorm maximum tops (m)"
			case 10:
				return "Thunderstorm coverage (code table (4.204))"
			case 11:
				return "Cloud base (m)"
			case 12:
				return "Cloud top (m)"
			case 13:
				return "Ceiling (m)"
			case 14:
				return "Non-Convective Cloud Cover (%)"
			case 15:
				return "Cloud Work Function (J kg-1)"
			case 16:
				return "Convective Cloud Efficiency	(Proportion)"
			case 17:
				return "Total Condensate	(kg kg-1)"
			case 18:
				return "Total Column-Integrated Cloud Water	(kg m-2)"
			case 19:
				return "Total Column-Integrated Cloud Ice (kg m-2)"
			case 20:
				return "Total Column-Integrated Condensate (kg m-2)"
			case 21:
				return "Ice fraction of total condensate	(Proportion)"
			case 22:
				return "Cloud Cover	(%)"
			case 23:
				return "Cloud Ice Mixing Ratio (kg kg-1)"
			case 24:
				return "Sunshine	(Numeric)"
			case 25:
				return "Horizontal Extent of Cumulonimbus (CB) (%)"
			case 26:
				return "Height of Convective Cloud Base (m)"
			case 27:
				return "Height of Convective Cloud Top (m)"
			case 28:
				return "Number Concentration of Cloud Droplets (kg-1)"
			case 29:
				return "Number Concentration of Cloud Ice (kg-1)"
			case 30:
				return "Number Density of Cloud Droplets (m-3)"
			case 31:
				return "Number Density of Cloud Ice (m-3)"
			case 32:
				return "Fraction of Cloud Cover (Numeric)"
			case 33:
				return "Sunshine Duration (s)"
			case 192:
				return "Non-Convective Cloud Cover (%)"
			case 193:
				return "Cloud Work Function (J kg-1)"
			case 194:
				return "Convective Cloud Efficiency (non-dim)"
			case 195:
				return "Total Condensate	(kg kg-1)"
			case 196:
				return "Total Column-Integrated Cloud Water (kg m-2)"
			case 197:
				return "Total Column-Integrated Cloud Ice (kg m-2)"
			case 198:
				return "Total Column-Integrated Condensate (kg m-2)"
			case 199:
				return "Ice fraction of total condensate	(non-dim)"
			case 200:
				return "Convective Cloud Mass Flux (Pa s-1)"
			case 201:
				return "Sunshine Duration (s)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 7: Thermodynamic Stability Indices
		case 7:
			switch number {
			case 0:
				return "Parcel lifted index (to 500 hPa) (K)"
			case 1:
				return "Best lifted index (to 500 hPa) (K)"
			case 2:
				return "K index K"
			case 3:
				return "KO index (K)"
			case 4:
				return "Total totals index (K)"
			case 5:
				return "Sweat index (numeric)"
			case 6:
				return "Convective available potential energy J kg-1"
			case 7:
				return "Convective inhibition (J kg-1)"
			case 8:
				return "Storm relative helicity (J kg-1)"
			case 9:
				return "Energy helicity index (numeric)"
			case 10:
				return "Surface Lifted Index (K)"
			case 11:
				return "Best (4 layer) Lifted Index (K)"
			case 12:
				return "Richardson Number (Numeric)"
			case 13:
				return "Showalter Index (K)"
			case 15:
				return "Updraft Helicity	(m2 s-2)"
			case 192:
				return "Surface Lifted Index (K)"
			case 193:
				return "Best (4 layer) Lifted Index (K)"
			case 194:
				return "Richardson Number (Numeric)"
			case 195:
				return "Convective Weather Detection Index"
			case 196:
				return "Ultra Violet Index (W m-2)"
			case 197:
				return "Updraft Helicity	(m2 s-2)"
			case 198:
				return "Leaf Area Index"
			case 199:
				return "Hourly Maximum of Updraft Helicity over Layer 2km to 5 km AGL (m2 s-2)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 13: Aerosols
		case 13:
			switch number {
			case 0:
				return "Aerosol type (code table (4.205))"
			case 192:
				return "Particulate matter (coarse)	µg m-3"
			case 193:
				return "Particulate matter (fine)	(µg m-3)"
			case 194:
				return "Particulate matter (fine)	(log10 (µg m-3))"
			case 195:
				return "Integrated column particulate matter (fine)	(log10 (µg m-3))"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 14: Trace Gases
		case 14:
			switch number {
			case 0:
				return "Total ozone (Dobson)"
			case 1:
				return "Ozone Mixing Ratio (kg kg-1)"
			case 2:
				return "Total Column Integrated Ozone (DU)"
			case 192:
				return "Ozone Mixing Ratio (kg kg-1)"
			case 193:
				return "Ozone Concentration (ppb)"
			case 194:
				return "Categorical Ozone Concentration	(Non-Dim)"
			case 195:
				return "Ozone Vertical Diffusion (kg kg-1 s-1)"
			case 196:
				return "Ozone Production (kg kg-1 s-1)"
			case 197:
				return "Ozone Tendency (kg kg-1 s-1)"
			case 198:
				return "Ozone Production from Temperature Term	(kg kg-1 s-1)"
			case 199:
				return "Ozone Production from Column Ozone Term	(kg kg-1 s-1)"
			case 200:
				return "Ozone Daily Max from 1-hour Average (ppbV)"
			case 201:
				return "Ozone Daily Max from 8-hour Average (ppbV)"
			case 202:
				return "PM 2.5 Daily Max from 1-hour Average (μg m-3)"
			case 203:
				return "PM 2.5 Daily Max from 24-hour Average (μg m-3)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 15: Radar
		case 15:
			switch number {
			case 0:
				return "Base spectrum width (m s-1)"
			case 1:
				return "Base reflectivity (dB)"
			case 2:
				return "Base radial velocity (m s-1)"
			case 3:
				return "Vertically-integrated liquid (kg m-1)"
			case 4:
				return "Layer-maximum base reflectivity (dB)"
			case 5:
				return "Precipitation (kg m-2)"
			case 6:
				return "Radar spectra (1)"
			case 7:
				return "Radar spectra (2)"
			case 8:
				return "Radar spectra (3)"
			case 9:
				return "Reflectivity of Cloud Droplets (dB)"
			case 10:
				return "Reflectivity of Cloud Ice (dB)"
			case 11:
				return "Reflectivity of Snow (dB)"
			case 12:
				return "Reflectivity of Rain (dB)"
			case 13:
				return "Reflectivity of Graupel (dB)"
			case 14:
				return "Reflectivity of Hail (dB)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 16: Forecast Radar Imagery
		case 16:
			switch number {
			case 0:
				return "Equivalent radar reflectivity factor for rain(m m6 m-3)"
			case 1:
				return "Equivalent radar reflectivity factor for snow (m m6 m-3)"
			case 2:
				return "Equivalent radar reflectivity factor for parameterized convection (m m6 m-3)"
			case 3:
				return "Echo Top (m)"
			case 4:
				return "Reflectivity	(dB)"
			case 5:
				return "Composite reflectivity (dB)"
			case 192:
				return "Equivalent radar reflectivity factor for rain(m m6 m-3)"
			case 193:
				return "Equivalent radar reflectivity factor for snow (m m6 m-3)"
			case 194:
				return "Equivalent radar reflectivity factor for parameterized convection (m m6 m-3)"
			case 195:
				return "Reflectivity (dB)"
			case 196:
				return "Composite reflectivity (dB)"
			case 197:
				return "Echo Top (m)"
			case 198:
				return "Hourly Maximum of Simulated Reflectivity at 1 km AGL (dB)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 17: Electrodynamics
		case 17:
			switch number {
			case 192:
				return "Lightning (non-dim)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 18: Nuclear/radiology
		case 18:
			switch number {
			case 0:
				return "Air concentration of Caesium 137 (Bq m-3)"
			case 1:
				return "Air concentration of Iodine 131 (Bq m-3)"
			case 2:
				return "Air concentration of radioactive pollutant (Bq m-3)"
			case 3:
				return "Ground deposition of Caesium 137 (Bq m-2)"
			case 4:
				return "Ground deposition of Iodine 131 (Bq m-2)"
			case 5:
				return "Ground deposition of radioactive pollutant (Bq m-2)"
			case 6:
				return "Time-integrated air concentration of caesium pollutant (Bq s m-3)"
			case 7:
				return "Time-integrated air concentration of iodine pollutant (Bq s m-3)"
			case 8:
				return "Time-integrated air concentration of radioactive pollutant (Bq s m-3)"
			case 10:
				return "Air Concentration (Bq m-3"
			case 11:
				return "Wet Deposition (Bq m-2)"
			case 12:
				return "Dry Deposition (Bq m-2)"
			case 13:
				return "Total Deposition (Wet + Dry)	(Bq m-2)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown", number)
			}
		//Parameter Category 19: Physical atmospheric properties
		case 19:
			switch number {
			case 0:
				return "Visibility (m)"
			case 1:
				return "Albedo (%)"
			case 2:
				return "Thunderstorm probability (%)"
			case 3:
				return "mixed layer depth (m)"
			case 4:
				return "Volcanic ash (code table (4.206))"
			case 5:
				return "Icing top (m)"
			case 6:
				return "Icing base (m)"
			case 7:
				return "Icing (code table (4.207))"
			case 8:
				return "Turbulence top (m)"
			case 9:
				return "Turbulence base (m)"
			case 10:
				return "Turbulence (code table (4.208))"
			case 11:
				return "Turbulent kinetic energy (J kg-1)"
			case 12:
				return "Planetary boundary layer regime (code table (4.209))"
			case 13:
				return "Contrail intensity (code table (4.210))"
			case 14:
				return "Contrail engine type (code table (4.211))"
			case 15:
				return "Contrail top (m)"
			case 16:
				return "Contrail base (m)"
			case 17:
				return "Maximum Snow Albedo (%)"
			case 18:
				return "Snow-Free Albedo	(%)"
			case 19:
				return "Snow Albedo (%)"
			case 20:
				return "Icing (%)"
			case 21:
				return "In-Cloud Turbulence (%)"
			case 22:
				return "Clear Air Turbulence (CAT) (%)"
			case 23:
				return "Supercooled Large Droplet (SLD) Probability (%)"
			case 24:
				return "Convective Turbulent Kinetic Energy (J kg-1)"
			case 25:
				return "Weather Interpretation ww (WMO)"
			case 26:
				return "Convective Outlook"
			case 192:
				return "Maximum Snow Albedo	(%)"
			case 193:
				return "Snow-Free Albedo	(%)"
			case 194:
				return "Slight risk convective outlook (categorical)"
			case 195:
				return "Moderate risk convective outlook (categorical)"
			case 196:
				return "High risk convective outlook (categorical)"
			case 197:
				return "Tornado probability (%)"
			case 198:
				return "Hail probability(%)"
			case 199:
				return "Wind probability	(%)"
			case 200:
				return "Significant Tornado probability	(%)"
			case 201:
				return "Significant Hail probability	(%)"
			case 202:
				return "Significant Wind probability	(%)"
			case 203:
				return "Categorical Thunderstorm (Code table 4.222)"
			case 204:
				return "Number of mixed layers next to surface (integer)"
			case 205:
				return "Flight Category"
			case 206:
				return "Confidence - Ceiling"
			case 207:
				return "Confidence - Visibility"
			case 208:
				return "Confidence - Flight Category"
			case 209:
				return "Low-Level aviation interest"
			case 210:
				return "High-Level aviation interest"
			case 211:
				return "Visible, Black Sky Albedo (%)"
			case 212:
				return "Visible, White Sky Albedo (%)"
			case 213:
				return "Near IR, Black Sky Albedo (%)"
			case 214:
				return "Near IR, White Sky Albedo (%)"
			case 215:
				return "Total Probability of Severe Thunderstorms (Days 2,3) (%)"
			case 216:
				return "Total Probability of Extreme Severe Thunderstorms (Days 2,3)	(%)"
			case 217:
				return "Supercooled Large Droplet (SLD) Icing (See Table 4.207)"
			case 218:
				return "Radiative emissivity"
			case 219:
				return "Turbulence Potential Forecast Index"
			case 220:
				return "Categorical Severe Thunderstorm (Code table 4.222)"
			case 221:
				return "Probability of Convection (%)"
			case 222:
				return "Convection Potential (Code table 4.222)"
			case 232:
				return "Volcanic Ash Forecast Transport and Dispersion (log10 (kg m-3))"
			case 233:
				return "Icing probability (non-dim)"
			case 234:
				return "Icing severity (non-dim)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		// Parameter category 20: Meteorological products, Atmospheric Chemical Constituents
		case 20:
			switch number {
			case 0:
				return "Mass Density (Concentration) (kg m-3)"
			case 1:
				return "Column-Integrated Mass Density (kg m-2)"
			case 2:
				return "Mass Mixing Ratio (kg kg-1)"
			case 3:
				return "Atmosphere Emission Mass Flux (kg m-2 s-1)"
			case 4:
				return "Atmosphere Net Production Mass Flux (kg m-2 s-1)"
			case 5:
				return "Atmosphere Net Production And Emission Mass Flux (kg m-2 s-1)"
			case 6:
				return "Surface Dry Deposition Mass Flux (kg m-2 s-1)"
			case 7:
				return "Surface Wet Deposition Mass Flux (kg m-2 s-1)"
			case 8:
				return "Atmosphere Re-Emission Mass Flux (kg m-2 s-1)"
			case 9:
				return "Wet Deposition by Large-Scale Precipitation Mass Flux (kg m-2 s-1)"
			case 10:
				return "Wet Deposition by Convective Precipitation Mass Flux (kg m-2 s-1)"
			case 11:
				return "Sedimentation Mass Flux (kg m-2 s-1)"
			case 12:
				return "Dry Deposition Mass Flux (kg m-2 s-1)"
			case 13:
				return "Transfer From Hydrophobic to Hydrophilic (kg kg-1 s-1)"
			case 14:
				return "Transfer From SO2 (Sulphur Dioxide) to SO4 (Sulphate) (kg kg-1 s-1)"
			case 50:
				return "Amount in Atmosphere (mol)"
			case 51:
				return "Concentration In Air (mol m-3)"
			case 52:
				return "Volume Mixing Ratio (Fraction in Air) (mol mol-1)"
			case 53:
				return "Chemical Gross Production Rate of Concentration (mol m-3 s-1)"
			case 54:
				return "Chemical Gross Destruction Rate of Concentration (mol m-3 s-1)"
			case 55:
				return "Surface Flux (mol m-2 s-1)"
			case 56:
				return "Changes Of Amount in Atmosphere (mol s-1)"
			case 57:
				return "Total Yearly Average Burden of The Atmosphere (mol)"
			case 58:
				return "Total Yearly Average Atmospheric Loss (mol s-1)"
			case 59:
				return "Aerosol Number Concentration (m-3)"
			case 100:
				return "Surface Area Density (Aerosol) (m-1)"
			case 101:
				return "Atmosphere Optical Thickness (m)"
			case 102:
				return "Aerosol Optical Thickness (Numeric)"
			case 103:
				return "Single Scattering Albedo (Numeric)"
			case 104:
				return "Asymmetry Factor (Numeric)"
			case 105:
				return "Aerosol Extinction Coefficient (m-1)"
			case 106:
				return "Aerosol Absorption Coefficient (m-1)"
			case 107:
				return "Aerosol Lidar Backscatter from Satellite (m-1 sr-1)"
			case 108:
				return "Aerosol Lidar Backscatter from the Ground (m-1 sr-1)"
			case 109:
				return "Aerosol Lidar Extinction from Satellite (m-1)"
			case 110:
				return "Aerosol Lidar Extinction from the Ground (m-1)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 190: ASCII character string
		case 190:
			switch number {
			case 0:
				return "Arbitrary text string(CCITTIA5)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 191: Miscellaneous
		case 191:
			switch number {
			case 0:
				return "Seconds prior to initial reference time (s)"
			case 1:
				return "Geographical Latitude (° N)"
			case 2:
				return "Geographical Longitude (° E)"
			case 192:
				return "Latitude (-90 to 90)	(°)"
			case 193:
				return "East Longitude (0 to 360) (°)"
			case 194:
				return "Seconds prior to initial reference time (s)"
			case 195:
				return "Model Layer number (From bottom up)"
			case 196:
				return "Latitude (nearest neighbor) (-90 to 90)	(°)"
			case 197:
				return "East Longitude (nearest neighbor) (0 to 360)	(°)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 192: Covariance
		case 192:
			switch number {
			case 1:
				return "Covariance between zonal and meridional components of the wind. Defined as [uv]-[u][v], where \"[]\" indicates the mean over the indicated time span.	(m2 s-2)"
			case 2:
				return "Covariance between zonal component of the wind and temperature. Defined as [uT]-[u][T], where \"[]\" indicates the mean over the indicated time span.	(K m s-1)"
			case 3:
				return "Covariance between meridional component of the wind and temperature. Defined as [vT]-[v][T], where \"[]\" indicates the mean over the indicated time span.	(K m s-1)"
			case 4:
				return "Covariance between temperature and vertical component of the wind. Defined as [wT]-[w][T], where \"[]\" indicates the mean over the indicated time span. (K m s-1)"
			case 5:
				return "Covariance between zonal and zonal components of the wind. Defined as [uu]-[u][u], where \"[]\" indicates the mean over the indicated time span. (m2 s-2)"
			case 6:
				return "Covariance between meridional and meridional components of the wind. Defined as [vv]-[v][v], where \"[]\" indicates the mean over the indicated time span.	(m2 s-2)"
			case 7:
				return "Covariance between specific humidity and zonal components of the wind. Defined as [uq]-[u][q], where \"[]\" indicates the mean over the indicated time span.	kg (kg-1 m s-1))"
			case 8:
				return "Covariance between specific humidity and meridional components of the wind. Defined as [vq]-[v][q], where \"[]\" indicates the mean over the indicated time span.	(kg/kg*m/s)"
			case 9:
				return "Covariance between temperature and vertical components of the wind. Defined as [ΩT]-[Ω][T], where \"[]\" indicates the mean over the indicated time span.	(K*Pa/s)"
			case 10:
				return "Covariance between specific humidity and vertical components of the wind. Defined as [Ωq]-[Ω][q], where \"[]\" indicates the mean over the indicated time span. (kg/kg*Pa/s)"
			case 11:
				return "Covariance between surface pressure and surface pressure. Defined as [Psfc]-[Psfc][Psfc], where \"[]\" indicates the mean over the indicated time span. (Pa*Pa)"
			case 12:
				return "Covariance between specific humidity and specific humidy. Defined as [qq]-[q][q], where \"[]\" indicates the mean over the indicated time span. (kg/kg*kg/kg)"
			case 13:
				return "Covariance between vertical and vertical components of the wind. Defined as [ΩΩ]-[Ω][Ω], where \"[]\" indicates the mean over the indicated time span.	(Pa2/s2)"
			case 14:
				return "Covariance between temperature and temperature. Defined as [TT]-[T][T], where \"[]\" indicates the mean over the indicated time span.	(K*K)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		default:
			return fmt.Sprint("Unknown ", category)
		}
	//Product Discipline 1: Hydrologic products,
	case 1:
		switch category {
		//Parameter Category 0: Hydrology basic products
		case 0:
			switch number {
			case 0:
				return "Flash flood guidance (kg m-2)"
			case 1:
				return "Flash flood runoff (kg m-2)"
			case 2:
				return "Remotely sensed snow cover (code table 4.215)"
			case 3:
				return "Elevation of snow covered terrain (code table 4.216)"
			case 4:
				return "Snow water equivalent percent of normal (%)"
			case 5:
				return "Baseflow-Groundwater Runoff (kg m-2)"
			case 6:
				return "Storm Surface Runoff	(kg m-2)"
			case 192:
				return "Baseflow-Groundwater Runoff (kg m-2)"
			case 193:
				return "Storm Surface Runoff	(kg m-2)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 1: Hydrology probabilities
		case 1:
			switch number {
			case 0:
				return "Conditional percent precipitation amount fractile for an overall period (kg m-2)"
			case 1:
				return "Percent precipitation in a sub-period of an overall period (%)"
			case 2:
				return "Probability of 0.01 inch of precipitation (POP) (%)"
			case 192:
				return "Probability of Freezing Precipitation (%)"
			case 193:
				return "Probability of Frozen Precipitation (%)"
			case 194:
				return "Probability of precipitation exceeding flash flood guidance values (%)"
			case 195:
				return "Probability of Wetting Rain, exceeding in 0.10\" in a given time period (%)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 2: Inland water and sediment properties
		case 2:
			switch number {
			case 0:
				return "Water Depth (m)"
			case 1:
				return "Water Temperature (K)"
			case 2:
				return "Water Fraction (Proportion)"
			case 3:
				return "Sediment Thickness (m)"
			case 4:
				return "Sediment Temperature	(K)"
			case 5:
				return "Ice Thickness (m)"
			case 6:
				return "Ice Temperature (K)"
			case 7:
				return "Ice Cover (Proportion)"
			case 8:
				return "Land Cover (0=water, 1=land)	(Proportion)"
			case 9:
				return "Shape Factor with Respect to Salinity Profile"
			case 10:
				return "Shape Factor with Respect to Temperature Profile in Thermocline"
			case 11:
				return "Attenuation Coefficient of Water with Respect to Solar Radiation (m-1)"
			case 12:
				return "Salinity (kg kg-1)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		default:
			return fmt.Sprint("Unknown ", category)
		}
	//Product Discipline 2: Land surface products
	case 2:
		switch category {
		//Parameter Category 0: Vegetation/Biomass
		case 0:
			switch number {
			case 0:
				return "Land cover (1=land, 2=sea) (Proportion)"
			case 1:
				return "Surface roughness (m)"
			case 2:
				return "Soil temperature (K)"
			case 3:
				return "Soil moisture content (kg m-2)"
			case 4:
				return "Vegetation (%)"
			case 5:
				return "Water runoff (kg m-2)"
			case 6:
				return "Evapotranspiration (kg-2 s-1)"
			case 7:
				return "Model terrain height (m)"
			case 8:
				return "Land use (code table (4.212))"
			case 9:
				return "Volumetric Soil Moisture Content	(Proportion)"
			case 10:
				return "Ground Heat Flux (W m-2)"
			case 11:
				return "Moisture Availability (%)"
			case 12:
				return "Exchange Coefficient	(kg m-2 s-1)"
			case 13:
				return "Plant Canopy Surface Water (kg m-2)"
			case 14:
				return "Blackadar's Mixing Length Scale	(m)"
			case 15:
				return "Canopy Conductance (m s-1)"
			case 16:
				return "Minimal Stomatal Resistance (s m-1)"
			case 17:
				return "Wilting Point (Proportion)"
			case 18:
				return "Solar parameter in canopy conductance (Proportion)"
			case 19:
				return "Temperature parameter in canopy (Proportion)"
			case 20:
				return "Humidity parameter in canopy conductance	(Proportion)"
			case 21:
				return "Soil moisture parameter in canopy conductance (Proportion)"
			case 22:
				return "Soil Moisture (kg m-3)"
			case 23:
				return "Column-Integrated Soil Water (kg m-2)"
			case 24:
				return "Heat Flux (W m-2)"
			case 25:
				return "Volumetric Soil Moisture	(m3 m-3)"
			case 26:
				return "Wilting Point (kg m-3)"
			case 27:
				return "Volumetric Wilting Point	(m3 m-3)"
			case 28:
				return "Leaf Area Index (Numeric)"
			case 29:
				return "Evergreen Forest (Numeric)"
			case 30:
				return "Deciduous Forest (Numeric)"
			case 31:
				return "Normalized Differential Vegetation Index (NDVI)  (Numeric)"
			case 32:
				return "Root Depth of Vegetation (m)"
			case 192:
				return "Volumetric Soil Moisture Content	(Fraction)"
			case 193:
				return "Ground Heat Flux	(W m-2)"
			case 194:
				return "Moisture Availability (%)"
			case 195:
				return "Exchange Coefficient	((kg m-3) (m s-1))"
			case 196:
				return "Plant Canopy Surface Water	(kg m-2)"
			case 197:
				return "Blackadar’s Mixing Length Scale	(m)"
			case 198:
				return "Vegetation Type	(Integer (0-13))"
			case 199:
				return "Canopy Conductance (m s-1)"
			case 200:
				return "Minimal Stomatal Resistance	(s m-1)"
			case 201:
				return "Wilting Point (Fraction)"
			case 202:
				return "Solar parameter in canopy conductance (Fraction)"
			case 203:
				return "Temperature parameter in canopy conductance	(Fraction)"
			case 204:
				return "Humidity parameter in canopy conductance	(Fraction)"
			case 205:
				return "Soil moisture parameter in canopy conductance (Fraction)"
			case 206:
				return "Rate of water dropping from canopy to ground"
			case 207:
				return "Ice-free water surface (%)"
			case 208:
				return "Surface exchange coefficients for T and Q divided by delta z	(m s-1)"
			case 209:
				return "Surface exchange coefficients for U and V divided by delta z	(m s-1)"
			case 210:
				return "Vegetation canopy temperature (K)"
			case 211:
				return "Surface water storage (Kg m-2)"
			case 212:
				return "Liquid soil moisture content (non-frozen)	(Kg m-2)"
			case 213:
				return "Open water evaporation (standing water)	(W m-2)"
			case 214:
				return "Groundwater recharge	(Kg m-2)"
			case 215:
				return "Flood plain recharge	(Kg m-2)"
			case 216:
				return "Roughness length for heat (m)"
			case 217:
				return "Normalized Difference Vegetation Index"
			case 218:
				return "Land-sea coverage (nearest neighbor) [land=1,sea=0]"
			case 219:
				return "Asymptotic mixing length scale (m)"
			case 220:
				return "Water vapor added by precip assimilation	(Kg m-2)"
			case 221:
				return "Water condensate added by precip assimilation	(Kg m-2)"
			case 222:
				return "Water Vapor Flux Convergance (Vertical Int)	(Kg m-2)"
			case 223:
				return "Water Condensate Flux Convergance (Vertical Int)	(Kg m-2)"
			case 224:
				return "Water Vapor Zonal Flux (Vertical Int)	(Kg m-2)"
			case 225:
				return "Water Vapor Meridional Flux (Vertical Int)	(Kg m-2)"
			case 226:
				return "Water Condensate Zonal Flux (Vertical Int)	(Kg m-2)"
			case 227:
				return "Water Condensate Meridional Flux (Vertical Int)	(Kg m-2)"
			case 228:
				return "Aerodynamic conductance	(m s-1)"
			case 229:
				return "Canopy water evaporation	(W m-2)"
			case 230:
				return "Transpiration (W m-2)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 1: Agricultural Special Products
		case 1:
			switch number {
			case 192:
				return "Cold Advisory for Newborn Livestock"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 3: Soil Products
		case 3:
			switch number {
			case 0:
				return "Soil type (code table (4.213))"
			case 1:
				return "Upper layer soil temperature (K)"
			case 2:
				return "Upper layer soil moisture (kg m-3)"
			case 3:
				return "Lower layer soil moisture (kg m-3)"
			case 4:
				return "Bottom layer soil temperature (K)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 4: Fire Weather
		case 4:
			switch number {
			case 0:
				return "Fire Outlook	(See Table 4.224)"
			case 1:
				return "Fire Outlook Due to Dry Thunderstorm	(See Table 4.224)"
			case 2:
				return "Haines Index (Numeric)"
			case 255:
				return "Misssing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		default:
			return fmt.Sprint("Unknown ", category)

		}
	//Product discipline 3: Space products,
	case 3:
		switch category {
		//Parameter Category 0: Image format
		case 0:
			switch number {
			case 0:
				return "Scaled radiance (numeric)"
			case 1:
				return "Scaled albedo (numeric)"
			case 2:
				return "Scaled brightness temperature (numeric)"
			case 3:
				return "Scaled precipitable water (numeric)"
			case 4:
				return "Scaled lifted index (numeric)"
			case 5:
				return "Scaled cloud top pressure (numeric)"
			case 6:
				return "Scaled skin temperature (numeric)"
			case 7:
				return "Cloud mask (Code table 4.217)"
			case 8:
				return "Pixel scene type (See Table 4.218)"
			case 9:
				return "Fire Detection Indicator	(See Table 4.223)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 1: Quantitative
		case 1:
			switch number {
			case 0:
				return "Estimated precipitation (kg m-2)"
			case 1:
				return "Instantaneous Rain Rate (kg m-2 s-1)"
			case 2:
				return "Cloud Top Height (m)"
			case 3:
				return "Cloud Top Height Quality Indicator (Code table 4.219)"
			case 4:
				return "Estimated u-Component of Wind (m s-1)"
			case 5:
				return "Estimated v-Component of Wind (m s-1)"
			case 6:
				return "Number Of Pixels Used (Numeric)"
			case 7:
				return "Solar Zenith Angle (°)"
			case 8:
				return "Relative Azimuth Angle (°)"
			case 9:
				return "Reflectance in 0.6 Micron Channel (%)"
			case 10:
				return "Reflectance in 0.8 Micron Channel (%)"
			case 11:
				return "Reflectance in 1.6 Micron Channel (%)"
			case 12:
				return "Reflectance in 3.9 Micron Channel (%)"
			case 13:
				return "Atmospheric Divergence (s-1)"
			case 14:
				return "Cloudy Brightness Temperature (K)"
			case 15:
				return "Clear Sky Brightness Temperature	(K)"
			case 16:
				return "Cloudy Radiance (with respect to wave number) (W m-1 sr-1)"
			case 17:
				return "Clear Sky Radiance (with respect to wave number)	(W m-1 sr-1)"
			case 19:
				return "Wind Speed (m s-1)"
			case 20:
				return "Aerosol Optical Thickness at 0.635 µm"
			case 21:
				return "Aerosol Optical Thickness at 0.810 µm"
			case 22:
				return "Aerosol Optical Thickness at 1.640 µm"
			case 23:
				return "Angstrom Coefficient"
			case 192:
				return "Scatterometer Estimated U Wind Component (m s-1)"
			case 193:
				return "Scatterometer Estimated V Wind Component	(m s-1)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 192: Forecast Satellite imagery
		case 192:
			switch number {
			case 0:
				return "Simulated Brightness Temperature for GOES 12, Channel 2 (K)"
			case 1:
				return "Simulated Brightness Temperature for GOES 12, Channel 3 (K)"
			case 2:
				return "Simulated Brightness Temperature for GOES 12, Channel 4 (K)"
			case 3:
				return "Simulated Brightness Temperature for GOES 12, Channel 6 (K)"
			case 4:
				return "Simulated Brightness Counts for GOES 12, Channel 3 (Byte)"
			case 5:
				return "Simulated Brightness Counts for GOES 12, Channel 4 (Byte)"
			case 6:
				return "Simulated Brightness Temperature for GOES 11, Channel 2 (K)"
			case 7:
				return "Simulated Brightness Temperature for GOES 11, Channel 3 (K)"
			case 8:
				return "Simulated Brightness Temperature for GOES 11, Channel 4 (K)"
			case 9:
				return "Simulated Brightness Temperature for GOES 11, Channel 5 (K)"
			case 10:
				return "Simulated Brightness Temperature for AMSRE on Aqua, Channel 9 (K)"
			case 11:
				return "Simulated Brightness Temperature for AMSRE on Aqua, Channel 10 (K)"
			case 12:
				return "Simulated Brightness Temperature for AMSRE on Aqua, Channel 11 (K)"
			case 13:
				return "Simulated Brightness Temperature for AMSRE on Aqua, Channel 12 (K)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		default:
			return fmt.Sprint("Unknown ", category)
		}
	//Product discipline 4: Space Weather Products
	case 4:
		switch category {
		//Parameter Category 0: Temperature
		case 0:
			switch number {
			case 0:
				return "Temperature (K)"
			case 1:
				return "Electron Temperature	(K)"
			case 2:
				return "Proton Temperature (K)"
			case 3:
				return "Ion Temperature (K)"
			case 4:
				return "Parallel Temperature	(K)"
			case 5:
				return "Perpendicular Temperature (K)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		case 1:
			switch number {
			case 0:
				return "Velocity Magnitude (Speed) (m s-1)"
			case 1:
				return "1st Vector Component of Velocity (Coordinate system dependent) (m s-1)"
			case 2:
				return "2nd Vector Component of Velocity (Coordinate system dependent) (m s-1)"
			case 3:
				return "3rd Vector Component of Velocity (Coordinate system dependent) (m s-1)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		case 2:
			switch number {
			case 0:
				return "Particle Number Density (m-3)"
			case 1:
				return "Electron Density	(m-3)"
			case 2:
				return "Proton Density (m-3)"
			case 3:
				return "Ion Density (m-3)"
			case 4:
				return "Vertical Electron Content (m-2)"
			case 5:
				return "HF Absorption Frequency (Hz)"
			case 6:
				return "HF Absorption (dB)"
			case 7:
				return "Spread F	(m)"
			case 8:
				return "h'F (m)"
			case 9:
				return "Critical Frequency (Hz)"
			case 10:
				return "Scintillation (Numeric)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		case 3:
			switch number {
			case 0:
				return "Magnetic Field Magnitude	(T)"
			case 1:
				return "1st Vector Component of Magnetic Field (T)"
			case 2:
				return "2nd Vector Component of Magnetic Field (T)"
			case 3:
				return "3rd Vector Component of Magnetic Field (T)"
			case 4:
				return "Electric Field Magnitude	(V m-1)"
			case 5:
				return "1st Vector Component of Electric Field (T)"
			case 6:
				return "2nd Vector Component of Electric Field (T)"
			case 7:
				return "3rd Vector Component of Electric Field (T)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		case 4:
			switch number {
			case 0:
				return "Proton Flux (Differential)	((m2 s sr eV)-1)"
			case 1:
				return "Proton Flux (Integral)	((m2 s sr)-1)"
			case 2:
				return "Electron Flux (Differential)	((m2 s sr eV)-1)"
			case 3:
				return "Electron Flux (Integral)	((m2 s sr)-1)"
			case 4:
				return "Heavy Ion Flux (Differential)	((m2 s sr eV / nuc)-1)"
			case 5:
				return "Heavy Ion Flux (iIntegral)	((m2 s sr)-1)"
			case 6:
				return "Cosmic Ray Neutron Flux	(h-1)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		case 5:
			switch number {
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		case 6:
			switch number {
			case 0:
				return "Integrated Solar Irradiance (W m-2)"
			case 1:
				return "Solar X-ray Flux (XRS Long) (W m-2)"
			case 2:
				return "Solar X-ray Flux (XRS Short) (W m-2)"
			case 3:
				return "Solar EUV Irradiance	(W m-2)"
			case 4:
				return "Solar Spectral Irradiance (W m-2 nm-1)"
			case 5:
				return "F10.7 (W m-2 Hz-1)"
			case 6:
				return "Solar Radio Emissions (W m-2 Hz-1)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		case 7:
			switch number {
			case 0:
				return "Limb Intensity (m-2 s-1)"
			case 1:
				return "Disk Intensity (m-2 s-1)"
			case 2:
				return "Disk Intensity Day (m-2 s-1)"
			case 3:
				return "Disk Intensity Night	(m-2 s-1)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		case 8:
			switch number {
			case 0:
				return "X-Ray Radiance (W sr-1 m-2)"
			case 1:
				return "EUV Radiance	(W sr-1 m-2)"
			case 2:
				return "H-Alpha Radiance	(W sr-1 m-2)"
			case 3:
				return "White Light Radiance	(W sr-1 m-2)"
			case 4:
				return "CaII-K Radiance	(W sr-1 m-2)"
			case 5:
				return "White Light Coronagraph Radiance	(W sr-1 m-2)"
			case 6:
				return "Heliospheric Radiance (W sr-1 m-2)"
			case 7:
				return "Thematic Mask (Numeric)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		case 9:
			switch number {
			case 0:
				return "Pedersen Conductivity (S m-1)"
			case 1:
				return "Hall Conductivity (S m-1)"
			case 2:
				return "Parallel Conductivity (S m-1)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		default:
			return fmt.Sprint("Unknown ", category)
		}
	//Product Discipline 10: Oceanographic products
	case 10:
		switch category {
		//Parameter Category 0: Waves
		case 0:
			switch number {
			case 0:
				return "Wave spectra (1)"
			case 1:
				return "Wave spectra (2)"
			case 2:
				return "Wave spectra (3)"
			case 3:
				return "Significant height of combined wind waves and swell (m)"
			case 4:
				return "Direction of wind waves (Degree true)"
			case 5:
				return "Significant height of wind waves (m)"
			case 6:
				return "Mean period of wind waves (s)"
			case 7:
				return "Direction of swell waves (Degree true)"
			case 8:
				return "Significant height of swell waves (m)"
			case 9:
				return "Mean period of swell waves (s)"
			case 10:
				return "Primary wave direction (Degree true)"
			case 11:
				return "Primary wave mean period (s)"
			case 12:
				return "Secondary wave direction (Degree true)"
			case 13:
				return "Secondary wave mean period (s)"
			case 14:
				return "Direction of Combined Wind Waves and Swell (degree true)"
			case 15:
				return "Mean Period of Combined Wind Waves and Swell	(s)"
			case 16:
				return "Coefficient of Drag With Waves"
			case 17:
				return "Friction Velocity (m s-1)"
			case 18:
				return "Wave Stress (N m-2)"
			case 19:
				return "Normalised Waves Stress"
			case 20:
				return "Mean Square Slope of Waves"
			case 21:
				return "U-component Surface Stokes Drift (m s-1)"
			case 22:
				return "V-component Surface Stokes Drift (m s-1)"
			case 23:
				return "Period of Maximum Individual Wave Height (s)"
			case 24:
				return "Maximum Individual Wave Height (m)"
			case 25:
				return "Inverse Mean Wave Frequency (s)"
			case 26:
				return "Inverse Mean Frequency of The Wind Waves (s)"
			case 27:
				return "Inverse Mean Frequency of The Total Swell (s)"
			case 28:
				return "Mean Zero-Crossing Wave Period (s)"
			case 29:
				return "Mean Zero-Crossing Period of The Wind Waves (s)"
			case 30:
				return "Mean Zero-Crossing Period of The Total Swell (s)"
			case 31:
				return "Wave Directional Width"
			case 32:
				return "Directional Width of The Wind Waves"
			case 33:
				return "Directional Width of The Total Swell"
			case 34:
				return "Peak Wave Period (s)"
			case 35:
				return "Peak Period of The Wind Waves (s)"
			case 36:
				return "Peak Period of The Total Swell (s)"
			case 37:
				return "Altimeter Wave Height (m)"
			case 38:
				return "Altimeter Corrected Wave Height (m)"
			case 39:
				return "Altimeter Range Relative Correction"
			case 40:
				return "10 Metre Neutral Wind Speed Over Waves (m s-1)"
			case 41:
				return "10 Metre Wind Direction Over Waves (degree true)"
			case 42:
				return "Wave Engery Spectrum (m-2 s rad-1)"
			case 43:
				return "Kurtosis of The Sea Surface Elevation Due to Waves"
			case 44:
				return "Benjamin-Feir Index"
			case 45:
				return "Spectral Peakedness Factor (s-1)"
			case 46:
				return "2-Dimension Spectral Energy Density E(f,θ)  (m s-2)"
			case 47:
				return "Frequency Spectral Energy Density E(f)=∫E(f,θ)dθ   (m s-2)"
			case 48:
				return "Frequency Spectral Energy Density E(f)=∫E(f,θ)dθ/m0   (m s-2)"
			case 50:
				return "Significant Wave Height (m)"
			case 51:
				return "Peak Direction (degree true)"
			case 52:
				return "Wave Steepness (proportion)"
			case 53:
				return "Mean Wave Directional Spread (degree)"
			case 54:
				return "Wind-Forced Fraction of the Wave Spectrum (proportion)"
			case 55:
				return "Energy Mean Wave Period (TMM1)  (s)"
			case 56:
				return "First Directional Moments Mean Wave Direction  (degree true)"
			case 57:
				return "Second Directional Moments Mean Wave Direction  (degree true)"
			case 58:
				return "First Directional Moments Mean Directional Spread (degree)"
			case 59:
				return "Second Directional Moments Mean Directional Spread (degree)"
			case 60:
				return "Mean Wave Length (m)"
			case 61:
				return "Sxx Component Radiation Stress (N m-2)"
			case 62:
				return "Syy Component Radiation Stress (N m-2)"
			case 63:
				return "Sxy Component Radiation Stress (N m-2)"
			case 192:
				return "Wave Steepness (proportion)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 1: Currents
		case 1:
			switch number {
			case 0:
				return "Current direction (Degree true)"
			case 1:
				return "Current speed (m s-1)"
			case 2:
				return "u-component of current (m s-1)"
			case 3:
				return "v-component of current(m s-1)"
			case 192:
				return "Ocean Mixed Layer U Velocity (m s-1)"
			case 193:
				return "Ocean Mixed Layer V Velocity (m s-1)"
			case 194:
				return "Barotropic U velocity (m s-1)"
			case 195:
				return "Barotropic V velocity (m s-1)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 2: Ice
		case 2:
			switch number {
			case 0:
				return "Ice cover (Proportion)"
			case 1:
				return "Ice thickness (m)"
			case 2:
				return "Direction of ice drift (Degree true)"
			case 3:
				return "Speed of ice drift (m s-1)"
			case 4:
				return "u-component of ice drift (m s-1)"
			case 5:
				return "v-component of ice drift (m s-1)"
			case 6:
				return "Ice growth rate (m s-1)"
			case 7:
				return "Ice divergence (s-1)"
			case 8:
				return "Ice Temperature (K)"
			case 9:
				return "Ice Internal Pressure (Pa m)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 3: Surface Properties
		case 3:
			switch number {
			case 0:
				return "Water temperature (K)"
			case 1:
				return "Deviation of sea level from mean (m)"
			case 192:
				return "Hurricane Storm Surge (m)"
			case 193:
				return "Extra Tropical Storm Surge (m)"
			case 194:
				return "Ocean Surface Elevation Relative to Geoid (m)"
			case 195:
				return "Sea Surface Height Relative to Geoid (m)"
			case 196:
				return "Ocean Mixed Layer Potential Density (Reference 2000m) (kg m-3)"
			case 197:
				return "Net Air-Ocean Heat Flux (W m-2)"
			case 198:
				return "Assimilative Heat Flux (W m-2)"
			case 199:
				return "Surface Temperature Trend (degree per day)"
			case 200:
				return "Surface Salinity Trend (psu per day)"
			case 201:
				return "Kinetic Energy (J kg-1)"
			case 202:
				return "Salt Flux (kg m-2 s-1)"
			case 242:
				return "20% Tropical Cyclone Storm Surge Exceedance (m)"
			case 243:
				return "30% Tropical Cyclone Storm Surge Exceedance (m)"
			case 244:
				return "40% Tropical Cyclone Storm Surge Exceedance (m)"
			case 245:
				return "50% Tropical Cyclone Storm Surge Exceedance (m)"
			case 246:
				return "60% Tropical Cyclone Storm Surge Exceedance (m)"
			case 247:
				return "70% Tropical Cyclone Storm Surge Exceedance (m)"
			case 248:
				return "80% Tropical Cyclone Storm Surge Exceedance (m)"
			case 249:
				return "90% Tropical Cyclone Storm Surge Exceedance (m)"
			case 250:
				return "Extra Tropical Storm Surge Combined Surge and Tide (m)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 4: Sub-surface Properties
		case 4:
			switch number {
			case 0:
				return "Main thermocline depth (m)"
			case 1:
				return "Main thermocline anomaly (m)"
			case 2:
				return "Transient thermocline depth (m)"
			case 3:
				return "Salinity (kg kg-1)"
			case 4:
				return "Ocean Vertical Heat Diffusivity (m2 s-1)"
			case 5:
				return "Ocean Vertical Salt Diffusivity (m2 s-1)"
			case 6:
				return "Ocean Vertical Momentum Diffusivity (m2 s-1)"
			case 7:
				return "Bathymetry (m)"
			case 11:
				return "Shape Factor With Respect To Salinity Profile"
			case 12:
				return "Shape Factor With Respect To Temperature Profile In Thermocline"
			case 13:
				return "Attenuation Coefficient Of Water With Respect to Solar Radiation	(m-1)"
			case 14:
				return "Water Depth (m)"
			case 15:
				return "Water Temperature (K)"
			case 192:
				return "3-D Temperature (° c)"
			case 193:
				return "3-D Salinity (psu)"
			case 194:
				return "Barotropic Kinectic Energy (J kg-1)"
			case 195:
				return "Geometric Depth Below Sea Surface (m)"
			case 196:
				return "Interface Depths (m)"
			case 197:
				return "Ocean Heat Content (J m-2)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		//Parameter Category 191: Miscellaneous
		case 191:
			switch number {
			case 0:
				return "Seconds Prior To Initial Reference Time (Defined In Section 1) (s)"
			case 1:
				return "Meridional Overturning Stream Function (m3 s-1)"
			case 255:
				return "Missing"
			default:
				return fmt.Sprint("Unknown ", number)
			}
		default:
			return fmt.Sprint("Unknown ", category)
		}
	default:
		return fmt.Sprint("Unknown ", discipline)
	}
}

// Type of generating process (code table 4.3)
func ReadGeneratingProcessType(value int) string {
	switch value {
	case 0:
		return "Analysis"
	case 1:
		return "Initialization"
	case 2:
		return "Forecast"
	case 3:
		return "Bias corrected forecast"
	case 4:
		return "Ensemble forecast"
	case 5:
		return "Probability forecast"
	case 6:
		return "Forecast error"
	case 7:
		return "Analysis error"
	case 8:
		return "Observation"
	case 9:
		return "Climatological"
	case 10:
		return "Probability-Weighted Forecast"
	case 11:
		return "Bias-Corrected Ensemble Forecast"
	case 192:
		return "Forecast Confidence Indicator"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Indicator of unit of time range (code table 4.4)
func ReadTimeRangeUnitIndicator(value int) string {
	switch value {
	case 0:
		return "Minute"
	case 1:
		return "Hour"
	case 2:
		return "Day"
	case 3:
		return "Month"
	case 4:
		return "Year"
	case 5:
		return "Decade(10 years)"
	case 6:
		return "Normal(30 years)"
	case 7:
		return "Century(100 years)"
	case 10:
		return "3 hours"
	case 11:
		return "6 hours"
	case 12:
		return "12 hours"
	case 13:
		return "Second"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Fixed surface types and units (code table 4.5)
func ReadSurfaceTypesUnits(value int) string {
	switch value {
	case 1:
		return "Ground or water surface"
	case 2:
		return "Cloud base level"
	case 3:
		return "Level of cloud tops"
	case 4:
		return "Level of 0°C isotherm"
	case 5:
		return "Level of adiabatic condensation lifted from the surface"
	case 6:
		return "Maximum wind level"
	case 7:
		return "Tropopause"
	case 8:
		return "Nominal top of the atmosphere"
	case 9:
		return "Sea bottom"
	case 10:
		return "Entire Atmosphere"
	case 11:
		return "Cumulonimbus Base (CB) (m)"
	case 12:
		return "Cumulonimbus Top (CT) (m)"
	case 20:
		return "Isothermal level (K)"
	case 100:
		return "Isobaric surface (Pa)"
	case 101:
		return "Mean sea level"
	case 102:
		return "Specific altitude above mean sea level (m)"
	case 103:
		return "Specified height level above ground (m)"
	case 104:
		return "Sigma level (\"sigma\" value)"
	case 105:
		return "Hybrid level"
	case 106:
		return "Depth below land surface (m)"
	case 107:
		return "Isentropic(theta) level (K)"
	case 108:
		return "Level at specified pressure difference from ground to level (Pa)"
	case 109:
		return "Potential vorticity surface (K m2 kg-1 s-1)"
	case 111:
		return "Eta* level"
	case 117:
		return "Mixed layer depth (m)"
	case 118:
		return "Hybrid Height Level"
	case 119:
		return "Hybrid Pressure Level"
	case 120:
		return "Pressure Thickness (Pa)"
	case 150:
		return "Generalized Vertical Height Coordinate"
	case 160:
		return "Depth below sea level (m)"
	case 161:
		return "Depth Below Water Surface (m)"
	case 162:
		return "Lake or River Bottom"
	case 163:
		return "Bottom Of Sediment Layer"
	case 164:
		return "Bottom Of Thermally Active Sediment Layer"
	case 165:
		return "Bottom Of Sediment Layer Penetrated By Thermal Wave"
	case 166:
		return "Maxing Layer"
	case 170:
		return "Ionospheric D-region Level"
	case 171:
		return "Ionospheric E-region Level"
	case 172:
		return "Ionospheric F1-region Level"
	case 173:
		return "Ionospheric F2-region Level"
	case 200:
		return "Entire atmosphere (considered as a single layer)"
	case 201:
		return "Entire ocean (considered as a single layer)"
	case 204:
		return "Highest tropospheric freezing level"
	case 206:
		return "Grid scale cloud bottom level"
	case 207:
		return "Grid scale cloud top level"
	case 209:
		return "Boundary layer cloud bottom level"
	case 210:
		return "Boundary layer cloud top level"
	case 211:
		return "Boundary layer cloud layer"
	case 212:
		return "Low cloud bottom level"
	case 213:
		return "Low cloud top level"
	case 214:
		return "Low cloud layer"
	case 215:
		return "Cloud ceiling"
	case 220:
		return "Planetary Boundary Layer"
	case 221:
		return "Layer Between Two Hybrid Levels"
	case 222:
		return "Middle cloud bottom level"
	case 223:
		return "Middle cloud top level"
	case 224:
		return "Middle cloud layer"
	case 232:
		return "High cloud bottom level"
	case 233:
		return "High cloud top level"
	case 234:
		return "High cloud layer"
	case 235:
		return "Ocean Isotherm Level (1/10 ° C)"
	case 236:
		return "Layer between two depths below ocean surface"
	case 237:
		return "Bottom of Ocean Mixed Layer (m)"
	case 238:
		return "Bottom of Ocean Isothermal Layer (m)"
	case 239:
		return "Layer Ocean Surface and 26C Ocean Isothermal Level"
	case 240:
		return "Ocean Mixed Layer"
	case 241:
		return "Ordered Sequence of Data"
	case 242:
		return "Convective cloud bottom level"
	case 243:
		return "Convective cloud top level"
	case 244:
		return "Convective cloud layer"
	case 245:
		return "Lowest level of the wet bulb zero"
	case 246:
		return "Maximum equivalent potential temperature level"
	case 247:
		return "Equilibrium level"
	case 248:
		return "Shallow convective cloud bottom level"
	case 249:
		return "Shallow convective cloud top level"
	case 251:
		return "Deep convective cloud bottom level"
	case 252:
		return "Deep convective cloud top level"
	case 253:
		return "Lowest bottom level of supercooled liquid water layer"
	case 254:
		return "Highest top level of supercooled liquid water layer"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Type of ensemble forecast (code table 4.6)
func ReadEnsembleForecastType(value int) string {
	switch value {
	case 0:
		return "Unperturbed high-resolution control forecast"
	case 1:
		return "Unperturbed low-resolution control forecast"
	case 2:
		return "Negatively perturbed forecast"
	case 3:
		return "Positively perturbed forecast"
	case 4:
		return "Multi-Model Forecast"
	case 192:
		return "Perturbed Ensemble Member"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

//  Derived forecast (code table 4.7)
func ReadDerivedForecast(value int) string {
	switch value {
	case 0:
		return "Unweighted mean of all members"
	case 1:
		return "Weighted mean of all members"
	case 2:
		return "Standard deviation with respect to cluster mean"
	case 3:
		return "Standard deviation with respect to cluster mean, normalized"
	case 4:
		return "Spread of all members"
	case 5:
		return "Large anomaly index of all members"
	case 6:
		return "Unweighted mean of the cluster members"
	case 7:
		return "Interquartile Range (Range between the 25th and 75th quantile)"
	case 8:
		return "Minimum Of All Ensemble Members"
	case 9:
		return "Maximum Of All Ensemble Members"
	case 192:
		return "Unweighted Mode of All Members"
	case 193:
		return "Percentile value (10%) of All Members"
	case 194:
		return "Percentile value (50%) of All Members"
	case 195:
		return "Percentile value (90%) of All Members"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Clustering Method (code table 4.8)
func ReadClusteringMethod(value int) string {
	switch value {
	case 0:
		return "Anomaly correlation"
	case 1:
		return "Root mean square"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Probability Type (code table 4.9)
func ReadProbabilityType(value int) string {
	switch value {
	case 0:
		return "Probability of event below lower limit"
	case 1:
		return "Probability of event above upper limit"
	case 2:
		return "Probability of event between lower and upper limits. The range includes the lower limit but not the upper limit."
	case 3:
		return "Probability of event above lower limit"
	case 4:
		return "Probability of event below upper limit"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Type of statistical processing (code table 4.10)
func ReadStatisticalProcessingType(value int) string {
	switch value {
	case 0:
		return "Average"
	case 1:
		return "Accumulation"
	case 2:
		return "Maximum"
	case 3:
		return "Minimum"
	case 4:
		return "Difference (Value at the end of time range minus value at the beginning)"
	case 5:
		return "Root mean square"
	case 6:
		return "Standard deviation"
	case 7:
		return "Covariance (Temporal variance)"
	case 8:
		return "Difference (Value at the start of time range minus value at the end)"
	case 9:
		return "Ratio"
	case 10:
		return "Standardized Anomaly"
	case 192:
		return "Climatological Mean Value: multiple year averages of quantities which are themselves means over some period of time (P2) less than a year. The reference time (R) indicates the date and time of the start of a period of time, given by R to R + P2, over which a mean is formed; N indicates the number of such period-means that are averaged together to form the climatological value, assuming that the N period-mean fields are separated by one year. The reference time indicates the start of the N-year climatology. N is given in octets 22-23 of the PDS. \n If P1 = 0 then the data averaged in the basic interval P2 are assumed to be continuous, i.e., all available data are simply averaged together. \n If P1 = 1 (the units of time - octet 18, code table 4 - are not relevant here) then the data averaged together in the basic interval P2 are valid only at the time (hour, minute) given in the reference time, for all the days included in the P2 period. The units of P2 are given by the contents of octet 18 and Table 4."
	case 193:
		return "Average of N forecasts (or initialized analyses); each product has forecast period of P1 (P1=0 for initialized analyses); products have reference times at intervals of P2, beginning at the given reference time."
	case 194:
		return "Average of N uninitialized analyses, starting at reference time, at intervals of P2."
	case 195:
		return "Average of forecast accumulations. P1 = start of accumulation period. P2 = end of accumulation period. Reference time is the start time of the first forecast, other forecasts at 24-hour intervals. Number in Ave = number of forecasts used."
	case 196:
		return "Average of successive forecast accumulations. P1 = start of accumulation period. P2 = end of accumulation period. Reference time is the start time of the first forecast, other forecasts at (P2 - P1) intervals. Number in Ave = number of forecasts used"
	case 197:
		return "Average of forecast averages. P1 = start of averaging period. P2 = end of averaging period. Reference time is the start time of the first forecast, other forecasts at 24-hour intervals. Number in Ave = number of forecast used"
	case 198:
		return "Average of successive forecast averages. P1 = start of averaging period. P2 = end of averaging period. Reference time is the start time of the first forecast, other forecasts at (P2 - P1) intervals. Number in Ave = number of forecasts used"
	case 199:
		return "Climatological Average of N analyses, each a year apart, starting from initial time R and for the period from R+P1 to R+P2."
	case 200:
		return "Climatological Average of N forecasts, each a year apart, starting from initial time R and for the period from R+P1 to R+P2."
	case 201:
		return "Climatological Root Mean Square difference between N forecasts and their verifying analyses, each a year apart, starting with initial time R and for the period from R+P1 to R+P2."
	case 202:
		return "Climatological Standard Deviation of N forecasts from the mean of the same N forecasts, for forecasts one year apart. The first forecast starts wtih initial time R and is for the period from R+P1 to R+P2."
	case 203:
		return "Climatological Standard Deviation of N analyses from the mean of the same N analyses, for analyses one year apart. The first analyses is valid for  period R+P1 to R+P2."
	case 204:
		return "Average of forecast accumulations. P1 = start of accumulation period. P2 = end of accumulation period. Reference time is the start time of the first forecast, other forecasts at 6-hour intervals. Number in Ave = number of forecast used"
	case 205:
		return "Average of forecast averages. P1 = start of averaging period. P2 = end of averaging period. Reference time is the start time of the first forecast, other forecasts at 6-hour intervals. Number in Ave = number of forecast used"
	case 206:
		return "Average of forecast accumulations. P1 = start of accumulation period. P2 = end of accumulation period. Reference time is the start time of the first forecast, other forecasts at 12-hour intervals. Number in Ave = number of forecast used"
	case 207:
		return "Average of forecast averages. P1 = start of averaging period. P2 = end of averaging period. Reference time is the start time of the first forecast, other forecasts at 12-hour intervals. Number in Ave = number of forecast used"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Type of time intervals (code table 4.11)
func ReadTimeIntervalsType(value int) string {
	switch value {
	case 1:
		return "Successive times processed have same forecast time, start time of forecast is incremented"
	case 2:
		return "Successive times processed have same start time of forecast, forecast time is incremented"
	case 3:
		return "Successive times processed have start time of forecast incremented and forecast time decremented so that valid time remains constant"
	case 4:
		return "Successive times processed have start time of forecast decremented and forecast time incremented so that valid time remains constant"
	case 5:
		return "Floating subinterval of time between forecast time and end of overall time interval"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Operating Mode (code table 4.12)
func ReadOperatingMode(value int) string {
	switch value {
	case 0:
		return "Maintenance mode"
	case 1:
		return "Clear air"
	case 2:
		return "Precipitation"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Quality Control Indicator (code table 4.13)
func ReadQualityControlIndicator(value int) string {
	switch value {
	case 0:
		return "No quality control applied"
	case 1:
		return "Quality control applied"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Clutter Filter Indicator (code table 4.14)
func ReadClutterFillerIndicator(value int) string {
	switch value {
	case 0:
		return "No clutter filter used"
	case 1:
		return "Clutter filter used"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// TYPE OF SPATIAL PROCESSING USED TO ARRIVE AT A GIVEN DATA VALUE FROM THE SOURCE DATA (code table 4.15)
func ReadSpatialProcessingType(value int) string {
	switch value {
	case 0:
		return "Data is calculated directly from the source grid with no interpolation"
	case 1:
		return "Bilinear interpolation using the 4 source grid grid-point values surrounding the nominal grid-point"
	case 2:
		return "Bicubic interpolation using the 4 source grid grid-point values surrounding the nominal grid-point"
	case 3:
		return "Using the value from the source grid grid-point which is nearest to the nominal grid-point"
	case 4:
		return "Budget interpolation using the 4 source grid grid-point values surrounding the nominal grid-point"
	case 5:
		return "Spectral interpolation using the 4 source grid grid-point values surrounding the nominal grid-point"
	case 6:
		return "Neighbor-budget interpolation using the 4 source grid grid-point values surrounding the nominal grid-point"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

//Type of interval
func ReadIntervalType(value int) string {
	switch value {
	case 0:
		return "Smaller than first limit"
	case 1:
		return "Greater than second limit"
	case 2:
		return "Between first and second limit. The range includes the first limit but not the second limit."
	case 3:
		return "Greater than first limit"
	case 4:
		return "Smaller than second limit"
	case 5:
		return "Smaller or equal first limit"
	case 6:
		return "Greater or equal second limit"
	case 7:
		return "Between first and second limit. The range includes the first limit and the second limit."
	case 8:
		return "Greater or equal first limit"
	case 9:
		return "Smaller or equal second limit"
	case 10:
		return "Between first and second limit. The range includes the second limit but not the first limit."
	case 11:
		return "Equal to first limit"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Precipitation Type (code table 4.201)
func ReadPrecipitationType(value int) string {
	switch value {
	case 1:
		return "Rain"
	case 2:
		return "Thunderstorm"
	case 3:
		return "Freezing rain"
	case 4:
		return "Mixed/ice"
	case 5:
		return "Snow"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Precipitable water category (code table 4.202)
func ReadPrecipitableWaterCategory(value int) string {
	switch value {
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Cloud type (code table 4.203)
func ReadCloudType(value int) string {
	switch value {
	case 0:
		return "Clear"
	case 1:
		return "Cumulonimbus"
	case 2:
		return "Stratus"
	case 3:
		return "Stratocumulus"
	case 4:
		return "Cumulus"
	case 5:
		return "Altostratus"
	case 6:
		return "Nimbostratus"
	case 7:
		return "Altocumulus"
	case 8:
		return "Cirrostratus"
	case 9:
		return "Cirrocumulus"
	case 10:
		return "Cirrus"
	case 11:
		return "Cumulonimbus - ground based fog beneath the lowest layer"
	case 12:
		return "Stratus - ground based fog beneath the lowest layer"
	case 13:
		return "Stratocumulus - ground based fog beneath the lowest layer"
	case 14:
		return "Cumulus - ground based fog beneath the lowest layer"
	case 15:
		return "Altostratus - ground based fog beneath the lowest layer"
	case 16:
		return "Nimbostratus - ground based fog beneath the lowest layer"
	case 17:
		return "Altocumulus - ground based fog beneath the lowest layer"
	case 18:
		return "Cirrostratus - ground based fog beneath the lowest layer"
	case 19:
		return "Cirrocumulus - ground based fog beneath the lowest layer"
	case 20:
		return "Cirrus - ground based fog beneath the lowest layer"
	case 191:
		return "Unknown type"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Thunderstorm coverage (code table 4.204)
func ReadThunderstormCoverage(value int) string {
	switch value {
	case 0:
		return "None"
	case 1:
		return "Isolated (1% - 2%)"
	case 2:
		return "Few (3% - 15%)"
	case 3:
		return "Scattered (16% - 45%)"
	case 4:
		return "Numerous (> 45%)"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Aerosol presence (code table 4.205)
func ReadAerosolPresence(value int) string {
	switch value {
	case 0:
		return "Aerosol not present"
	case 1:
		return "Aerosol present"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Volcanic ash (code table 4.206)
func ReadVolcanicAsh(value int) string {
	switch value {
	case 0:
		return "Not present"
	case 1:
		return "Present"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Icing (code table 4.207)
func ReadIcing(value int) string {
	switch value {
	case 0:
		return "None"
	case 1:
		return "Light"
	case 2:
		return "Moderate"
	case 3:
		return "Severe"
	case 4:
		return "Trace"
	case 5:
		return "Heavy"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Turbulence (code table 4.208)
func ReadTurbulence(value int) string {
	switch value {
	case 0:
		return "None(smooth)"
	case 1:
		return "Light"
	case 2:
		return "Moderate"
	case 3:
		return "Severe"
	case 4:
		return "Extreme"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Planetary boundary layer regime (code table 4.209)
func ReadPlanetaryBoundaryLayerRegime(value int) string {
	switch value {
	case 1:
		return "Stable"
	case 2:
		return "Mechanically driven turbulence"
	case 3:
		return "Forced convection"
	case 4:
		return "Free convection"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Contrail intensity (code table 4.210)
func ReadContrailIntensity(value int) string {
	switch value {
	case 0:
		return "Contrail not present"
	case 1:
		return "Contrail present"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Contrail engine type (code table 4.211)
func ReadContrailEngineType(value int) string {
	switch value {
	case 0:
		return "Low bypass"
	case 1:
		return "High bypass"
	case 2:
		return "Non bypass"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Land use (code table 4.212)
func ReadLandUse(value int) string {
	switch value {
	case 1:
		return "Urban land"
	case 2:
		return "Agriculture"
	case 3:
		return "Range land"
	case 4:
		return "Deciduous forest"
	case 5:
		return "Coniferous forest"
	case 6:
		return "Forest/wetland"
	case 7:
		return "Water"
	case 8:
		return "Wetlands"
	case 9:
		return "Desert"
	case 10:
		return "Tundra"
	case 11:
		return "Ice"
	case 12:
		return "Tropical forest"
	case 13:
		return "Savannah"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Soil type (code table 4.213)
func ReadSoilType(value int) string {
	switch value {
	case 1:
		return "Sand"
	case 2:
		return "Loamy sand"
	case 3:
		return "Sandy loam"
	case 4:
		return "Silt loam"
	case 5:
		return "Organic(redefined)"
	case 6:
		return "Sandy clay loam"
	case 7:
		return "Silt clay loam"
	case 8:
		return "Clay loam"
	case 9:
		return "Sandy clay"
	case 10:
		return "Silty clay"
	case 11:
		return "Clay"
	case 12:
		return "Loam"
	case 13:
		return "Peat"
	case 14:
		return "Rock"
	case 15:
		return "Ice"
	case 16:
		return "Water"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Remotely Sensed Snow Coverage (code table 4.215)
func ReadRemotelySensedSnowCoverage(value int) string {
	switch value {
	case 50:
		return "No-snow/no-cloud"
	case 100:
		return "Clouds"
	case 250:
		return "Snow"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Elevation of Snow Covered Terrain (code table 4.216)
func ReadSnowCoveredTerrainElevation(value int) string {
	if value >= 0 && value <= 90 {
		return "Elevation in increments of 100 m"
	} else if value == 254 {
		return "Clouds"
	} else if value == 255 {
		return "Missing"
	} else {
		return fmt.Sprint("Unknown ", value)
	}
}

// Cloud mask type (code table 4.217)
func ReadCloudMaskType(value int) string {
	switch value {
	case 0:
		return "Clear over water"
	case 1:
		return "Clear over land"
	case 2:
		return "Cloud"
	case 3:
		return "No data"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Pixel scene type (code table 4.218)
func ReadPixelSceneType(value int) string {
	switch value {
	case 0:
		return "No Scene Identified"
	case 1:
		return "Green Needle-Leafed Forest"
	case 2:
		return "Green Broad-Leafed Forest"
	case 3:
		return "Deciduous Needle-Leafed Forest"
	case 4:
		return "Deciduous Broad-Leafed Forest"
	case 5:
		return "Deciduous Mixed Forest"
	case 6:
		return "Closed Shrub-Land"
	case 7:
		return "Open Shrub-Land"
	case 8:
		return "Woody Savannah"
	case 9:
		return "Savannah"
	case 10:
		return "Grassland"
	case 11:
		return "Permanent Wetland"
	case 12:
		return "Cropland"
	case 13:
		return "Urban"
	case 14:
		return "Vegetation / Crops"
	case 15:
		return "Permanent Snow / Ice"
	case 16:
		return "Barren Desert"
	case 17:
		return "Water Bodies"
	case 18:
		return "Tundra"
	case 97:
		return "Snow / Ice on Land"
	case 98:
		return "Snow / Ice on Water"
	case 99:
		return "Sun-Glint"
	case 100:
		return "General Cloud"
	case 101:
		return "Low Cloud / Fog / Stratus"
	case 102:
		return "Low Cloud / Stratocumulus"
	case 103:
		return "Low Cloud / Unknown Type"
	case 104:
		return "Medium Cloud / Nimbostratus"
	case 105:
		return "Medium Cloud / Altostratus"
	case 106:
		return "Medium Cloud / Unknown Type"
	case 107:
		return "High Cloud / Cumulus"
	case 108:
		return "High Cloud / Cirrus"
	case 109:
		return "High Cloud / Unknown Type"
	case 110:
		return "Unknown Cloud Type"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Cloud top height quality indicator (code table 4.219)
func ReadCloudTopHeightQuality(value int) string {
	switch value {
	case 0:
		return "Nominal Cloud Top Height Quality"
	case 1:
		return "Fog In Segment"
	case 2:
		return "Poor Quality Height Estimation"
	case 3:
		return "Fog In Segment and Poor Quality Height Estimation"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Horizontal dimension processed (code table 4.220)
func ReadHorizontalDimensionProcessed(value int) string {
	switch value {
	case 0:
		return "Latitude"
	case 1:
		return "Longitude"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Treatment of missing data (code table 4.221)
func ReadMissingDataTreatment(value int) string {
	switch value {
	case 0:
		return "Not included"
	case 1:
		return "Extrapolated"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Categorical Result (code table 4.222)
func ReadCategoricalResult(value int) string {
	switch value {
	case 0:
		return "No"
	case 1:
		return "Yes"
	case 4:
		return "Low"
	case 6:
		return "Medium"
	case 8:
		return "High"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Fire Detection Indicator (code table 4.223)
func ReadFireDetection(value int) string {
	switch value {
	case 0:
		return "No Fire Detected"
	case 1:
		return "Possible Fire Detected"
	case 2:
		return "Probable Fire Detected"
	case 3:
		return "Missing"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Categorical Outlook (code table 4.224)
func ReadCategoricalOutlook(value int) string {
	switch value {
	case 0:
		return "No Risk Area"
	case 2:
		return "General Thunderstorm Risk Area"
	case 4:
		return "Slight Risk Area"
	case 6:
		return "Moderate Risk Area"
	case 8:
		return "High Risk Area"
	case 11:
		return "Dry Thunderstorm (Dry Lightning) Risk Area"
	case 14:
		return "Critical Risk Area"
	case 18:
		return "Extremely Critical Risk Area"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Aerosol Type (code table 4.233)
func ReadAerosolType(value int) string {
	switch value {
	case 0:
		return "Ozone (O3)"
	case 1:
		return "Water Vapour (H2O)"
	case 2:
		return "Methane (CH4)"
	case 3:
		return "Carbon Dioxide	(CO2)"
	case 4:
		return "Carbon Monoxide (CO)"
	case 5:
		return "Nitrogen Dioxide (NO2)"
	case 6:
		return "Nitrous Oxide	(N2O)"
	case 7:
		return "Formaldehyde (HCHO)"
	case 8:
		return "Sulphur Dioxide (SO2)"
	case 9:
		return "Ammonia (NH3)"
	case 10:
		return "Ammonium (NH4+)"
	case 11:
		return "Nitrogen Monoxide	(NO)"
	case 12:
		return "Atomic Oxygen	(O)"
	case 13:
		return "Nitrate Radical (NO3)"
	case 14:
		return "Hydroperoxyl Radical (HO2)"
	case 15:
		return "Dinitrogen Pentoxide (H2O5)"
	case 16:
		return "Nitrous Acid (HONO)"
	case 17:
		return "Nitric Acid (HNO3)"
	case 18:
		return "Peroxynitric Acid	(HO2NO2)"
	case 19:
		return "Hydrogen Peroxide	(H2O2)"
	case 20:
		return "Molecular Hydrogen (H)"
	case 21:
		return "Atomic Nitrogen (N)"
	case 22:
		return "Sulphate (SO42-)"
	case 23:
		return "Radon	(Rn)"
	case 24:
		return "Elemental Mercury	(Hg(O))"
	case 25:
		return "Divalent Mercury (Hg2+)"
	case 26:
		return "Atomic Chlorine (Cl)"
	case 27:
		return "Chlorine Monoxide	(ClO)"
	case 28:
		return "Dichlorine Peroxide (Cl2O2)"
	case 29:
		return "Hypochlorous Acid	(HClO)"
	case 30:
		return "Chlorine Nitrate (ClONO2)"
	case 31:
		return "Chlorine Dioxide (ClO2)"
	case 32:
		return "Atomic Bromide (Br)"
	case 33:
		return "Bromine Monoxide (BrO)"
	case 34:
		return "Bromine Chloride (BrCl)"
	case 35:
		return "Hydrogen Bromide (HBr)"
	case 36:
		return "Hypobromous Acid (HBrO)"
	case 37:
		return "Bromine Nitrate (BrONO2)"
	case 10000:
		return "Hydroxyl Radical (OH)"
	case 10001:
		return "Methyl Peroxy Radical (CH3O2)"
	case 10002:
		return "Methyl Hydroperoxide (CH3O2H)"
	case 10004:
		return "Methanol (CH3OH)"
	case 10005:
		return "Formic Acid (CH3OOH)"
	case 10006:
		return "Hydrogen Cyanide (HCN)"
	case 10007:
		return "Aceto Nitrile (CH3CN)"
	case 10008:
		return "Ethane	(C2H6)"
	case 10009:
		return "Ethene (= Ethylene) (C2H4)"
	case 10010:
		return "Ethyne (= Acetylene) (C2H2)"
	case 10011:
		return "Ethanol (C2H5OH)"
	case 10012:
		return "Acetic Acid (C2H5OOH)"
	case 10013:
		return "Peroxyacetyl Nitrate (CH3C(O)OONO2)"
	case 10014:
		return "Propane (C3H8)"
	case 10015:
		return "Propene (C3H6)"
	case 10016:
		return "Butanes (C4H10)"
	case 10017:
		return "Isoprene (C5H10)"
	case 10018:
		return "Alpha Pinene (C10H16)"
	case 10019:
		return "Beta Pinene (C10H16)"
	case 10020:
		return "Limonene (C10H16)"
	case 10021:
		return "Benzene (C6H6)"
	case 10022:
		return "Toluene (C7H8)"
	case 10023:
		return "Xylene	(C8H10)"
	case 10500:
		return "Dimethyl Sulphide (CH3SCH3)"
	case 20001:
		return "Hydrogen Chloride (HCL)"
	case 20002:
		return "CFC-11"
	case 20003:
		return "CFC-12"
	case 20004:
		return "CFC-113"
	case 20005:
		return "CFC-113a"
	case 20006:
		return "CFC-114"
	case 20007:
		return "CFC-115"
	case 20008:
		return "HCFC-22"
	case 20009:
		return "HCFC-141b"
	case 20010:
		return "HCFC-142b"
	case 20011:
		return "Halon-1202"
	case 20012:
		return "Halon-1211"
	case 20013:
		return "Halon-1301"
	case 20014:
		return "Halon-2402"
	case 20015:
		return "Methyl Chloride (HCC-40)"
	case 20016:
		return "Carbon Tetrachloride (HCC-10)"
	case 20017:
		return "HCC-140a (CH3CCl3)"
	case 20018:
		return "Methyl Bromide (HBC-40B1)"
	case 20019:
		return "Hexachlorocyclohexane (HCH)"
	case 20020:
		return "Alpha Hexachlorocyclohexane"
	case 20021:
		return "Hexachlorobiphenyl (PCB-153)"
	case 30000:
		return "Radioactive Pollutant (Tracer, defined by originating centre)"
	case 60000:
		return "HOx Radical (OH+HO2)"
	case 60001:
		return "Total Inorganic and Organic Peroxy Radicals (HO2+RO2)	(RO2)"
	case 60002:
		return "Passive Ozone"
	case 60003:
		return "NOx Expressed As Nitrogen (NOx)"
	case 60004:
		return "All Nitrogen Oxides (NOy) Expressed As Nitrogen (NOy)"
	case 60005:
		return "Total Inorganic Chlorine (Clx)"
	case 60006:
		return "Total Inorganic Bromine (Brx)"
	case 60007:
		return "Total Inorganic Chlorine Except HCl, ClONO2: ClOx"
	case 60008:
		return "Total Inorganic Bromine Except Hbr, BrONO2:BrOx"
	case 60009:
		return "Lumped Alkanes"
	case 60010:
		return "Lumped Alkenes"
	case 60011:
		return "Lumped Aromatic Coumpounds"
	case 60012:
		return "Lumped Terpenes"
	case 60013:
		return "Non-Methane Volatile Organic Compounds Expressed as Carbon	(NMVOC)"
	case 60014:
		return "Anthropogenic Non-Methane Volatile Organic Compounds Expressed as Carbon (aNMVOC)"
	case 60015:
		return "Biogenic Non-Methane Volatile Organic Compounds Expressed as Carbon (bNMVOC)"
	case 60016:
		return "Lumped Oxygenated Hydrocarbons	(OVOC)"
	case 62000:
		return "Total Aerosol"
	case 62001:
		return "Dust Dry"
	case 62002:
		return "water In Ambient"
	case 62003:
		return "Ammonium Dry"
	case 62004:
		return "Nitrate Dry"
	case 62005:
		return "Nitric Acid Trihydrate"
	case 62006:
		return "Sulphate Dry"
	case 62007:
		return "Mercury Dry"
	case 62008:
		return "Sea Salt Dry"
	case 62009:
		return "Black Carbon Dry"
	case 62010:
		return "Particulate Organic Matter Dry"
	case 62011:
		return "Primary Particulate Organic Matter Dry"
	case 62012:
		return "Secondary Particulate Organic Matter Dry"
	case 65535:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Atmospheric chemical or physical constituent type (code table 4.230)
//func ReadAtmosphericChemicalOrPhysicalConstituentType = func ReadAerosolType

// Wind-Generated Wave Sectral Description (code table 4.235)
func ReadWindGeneratedWaveSpectralDescription(value int) string {
	switch value {
	case 0:
		return "Total Wave Spectrum (combined wind waves and swell)"
	case 1:
		return "Generalized Partition"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Data Representation Template Number (code table 5.0)
func ReadDataRepresentationTemplateNumber(value int) string {
	switch value {
	case 0:
		return "Grid point data - simple packing (see Template 5.0)"
	case 1:
		return "Matrix value - simple packing (see Template 5.1)"
	case 2:
		return "Grid point data - complex packing (see Template 5.2)"
	case 3:
		return "Grid point data - complex packing and spatial differencing (see Template 5.3)"
	case 4:
		return "Grid Point Data - IEEE Floating Point Data (see Template 5.4)"
	case 42:
		return "Grid Point and Spectral Data - CCSDS szip (see Template 5.42)"
	case 50:
		return "Spectral data -simple packing (see Template 5.50)"
	case 51:
		return "Spherical harmonics data - complex packing (see Template 5.51)"
	case 61:
		return "Grid Point Data - Simple Packing With Logarithm Pre-processing (see Template 5.61)"
	case 200:
		return "Run Length Packing With Level Values (see Template 5.200)"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Type of original field values (code table 5.1)
func ReadOriginalFieldValuesType(value int) string {
	switch value {
	case 0:
		return "Floating point"
	case 1:
		return "Integer"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Matrix coordinate value function definition (code table 5.2)
func ReadMatrixCoordinateValueFunctionDefinition(value int) string {
	switch value {
	case 0:
		return "Explicit coordinate values set"
	case 1:
		return "Linear coordinates"
	case 11:
		return "Geometric coordinates"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Matrix coordinate parameter (code table 5.3)
func ReadMatrixCoordinateParameter(value int) string {
	switch value {
	case 1:
		return "Direction Degrees true"
	case 2:
		return "Frequency (s-1)"
	case 3:
		return "Radial number (2pi/lambda) (m-1)"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Group Splitting Method (code table 5.4)
func ReadGroupSplittingMethod(value int) string {
	switch value {
	case 0:
		return "Row by row splitting"
	case 1:
		return "General group splitting"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Missing Value Management for Complex Packing (code table 5.5)
func ReadMissingValueManagement(value int) string {
	switch value {
	case 0:
		return "No explicit missing values included within data values"
	case 1:
		return "Primary missing values included within data values"
	case 2:
		return "Primary and secondary missing values included within data values"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Order of Spatial Differencing (code table 5.6)
func ReadSpatialDifferencingOrder(value int) string {
	switch value {
	case 1:
		return "First-order spatial differencing"
	case 2:
		return "Second-order spatial differencing"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Precision of floating-point numbers (code table 5.7)
func ReadFloatingPointNumbersPrecision(value int) string {
	switch value {
	case 1:
		return "IEEE 32-bit (I=4 in Section 7)"
	case 2:
		return "IEEE 64-bit (I=8 in Section 7)"
	case 3:
		return "IEEE 128-bit (I=16 in Section 7)"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Type of compression (code table 5.40)
func ReadCompressionType(value int) string {
	switch value {
	case 0:
		return "Lossless"
	case 1:
		return "Lossy"
	case 255:
		return "Missing"
	default:
		return fmt.Sprint("Unknown ", value)
	}
}

// Bit Map Indicator (code table 6.0)
func ReadBitMapIndicator(value int) string {
	if value == 0 {
		return "A bit map applies to this product and is specified in this Section"
	} else if value >= 1 && value <= 253 {
		return "A bit map pre-determined by the originating/generating Centre applies to this product and is not specified in this Section"
	} else if value == 254 {
		return "A bit mapdefined previously in the same \"GRIB\" message applies to this product"
	} else if value == 255 {
		return "Missing"
	} else {
		return fmt.Sprint("Unknown ", value)
	}
}
