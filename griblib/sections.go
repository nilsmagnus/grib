package griblib

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"strings"
)

// Message is the entire message for a data-layer
type Message struct {
	Section0 Section0
	Section1 Section1
	Section2 Section2
	Section3 Section3
	Section4 Section4
	Section5 Section5
	Section6 Section6
	Section7 Section7
}

// Options is used to filter messages.
type Options struct {
	Operation               string    `json:"operation"`
	Discipline              int       `json:"discipline"` // -1 means all disciplines
	DataExport              bool      `json:"dataExport"`
	Category                int       `json:"category"` // -1 means all categories
	Filepath                string    `json:"filePath"`
	ReduceFilePath          string    `json:"reduceFilePath"`
	ExportType              int       `json:"exportType"`
	MaximumNumberOfMessages int       `json:"maximumNumberOfMessages"`
	GeoFilter               GeoFilter `json:"geoFilter"`
	Surface                 Surface   `json:"surfaceFilter"`
	// empty filter , GeoFilter{},  means no filter
}

const (
	//Grib is a magic-number specifying the grib2-format
	Grib = 0x47524942
	//EndSectionLength is the binary length of the end-section. Fixed number.
	EndSectionLength = 926365495
	//SupportedGribEdition is 2
	SupportedGribEdition = 2
)

// Data returns the data as an array of float64
func (message Message) Data() []float64 {
	return message.Section7.Data
}

// ReadNMessages reads at most n first messages from gribFile
// if an error occurs, the read messages and the error is returned
func ReadNMessages(gribFile io.Reader, n int) ([]*Message, error) {
	messages := make([]*Message, 0)

	for {
		message, messageErr := ReadMessage(gribFile)

		if messageErr != nil {
			if strings.Contains(messageErr.Error(), "EOF") {
				return messages, nil
			}
			log.Println("Error when parsing a message, ", messageErr.Error())
			return messages, messageErr
		}
		messages = append(messages, message)
		if len(messages) >= n {
			return messages, nil
		}
	}
}

// ReadMessages reads all message from gribFile
func ReadMessages(gribFile io.Reader) ([]*Message, error) {

	messages := make([]*Message, 0)

	for {
		message, messageErr := ReadMessage(gribFile)
		if messageErr != nil {
			if strings.Contains(messageErr.Error(), "EOF") {
				return messages, nil
			}
			log.Println("Error when parsing a message, ", messageErr.Error())
			return messages, messageErr
		}
		messages = append(messages, message)
	}
}

// ReadMessage reads the actual messages from a gribfile-reader (io.Reader from either file, http or any other io.Reader)
func ReadMessage(gribFile io.Reader) (*Message, error) {

	message := Message{}
	section0, headError := ReadSection0(gribFile)

	if headError != nil {
		return &message, headError
	}

	messageBytes := make([]byte, section0.MessageLength-16)

	numBytes, readError := gribFile.Read(messageBytes)

	if readError != nil {
		log.Println("Error reading message")
		return &message, readError
	}

	if numBytes != int(section0.MessageLength-16) {
		log.Println("Did not read full message")
	}

	return readMessage(bytes.NewReader(messageBytes), section0)

}

func readMessage(gribFile io.Reader, section0 Section0) (*Message, error) {

	message := Message{
		Section0: section0,
	}
	for {

		// pre-parse section head to decide which struct use
		sectionHead, headErr := ReadSectionHead(gribFile)
		if headErr != nil {
			log.Println("Error reading header", headErr.Error())
			return &message, headErr
		}

		if sectionHead.ContentLength() > 0 {
			var rawData = make([]byte, sectionHead.ContentLength())
			err := binary.Read(gribFile, binary.BigEndian, &rawData)
			if err != nil {
				return &message, err
			}
			byteReader := bytes.NewBuffer(rawData)

			switch sectionHead.Number {

			case 1:
				message.Section1, err = ReadSection1(byteReader, sectionHead.ContentLength())
			case 2:
				message.Section2, err = ReadSection2(byteReader, sectionHead.ContentLength())
			case 3:
				message.Section3, err = ReadSection3(byteReader, sectionHead.ContentLength())
			case 4:
				message.Section4, err = ReadSection4(byteReader, sectionHead.ContentLength())
			case 5:
				message.Section5, err = ReadSection5(byteReader, sectionHead.ContentLength())
			case 6:
				message.Section6, err = ReadSection6(byteReader, sectionHead.ContentLength())
			case 7:
				message.Section7, err = ReadSection7(byteReader, sectionHead.ContentLength(), message.Section5)
			case 8:
				// end-section, return
				return &message, nil
			default:
				err = fmt.Errorf("unknown section number %d  (Something bad with parser or files)", sectionHead.Number)
			}
			if err != nil {
				return &message, err
			}
		} else {
			return &message, nil
		}
	}
}

// Section0 is the indicator section http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_sect0.shtml
// This section serves to identify the start of the record in a human readable form, indicate the total length of the message,
// and indicate the Edition number of GRIB used to construct or encode the message. For GRIB2, this section is always 16 octets
// long.
//
//	| Octet Number | Content
//	-----------------------------------------------------------------------------------------
//	| 1-4          | 'GRIB' (Coded according to the International Alphabet Number 5)
//	| 5-6          | reserved
//	| 7            | Discipline (From Table 0.0)
//	| 8            | Edition number - 2 for GRIB2
//	| 9-16         | Total length of GRIB message in octets (All sections);
type Section0 struct {
	Indicator     uint32 `json:"indicator"`
	Reserved      uint16 `json:"reserved"`
	Discipline    uint8  `json:"discipline"`
	Edition       uint8  `json:"edition"`
	MessageLength uint64 `json:"messageLength"`
}

// ReadSection0 reads Section0 from an io.reader
func ReadSection0(reader io.Reader) (section0 Section0, err error) {
	section0.Indicator = 255
	err = binary.Read(reader, binary.BigEndian, &section0)
	if err != nil {
		return section0, err
	}

	if section0.Indicator == Grib {
		if section0.Edition != SupportedGribEdition {
			return section0, fmt.Errorf("unsupported grib edition %d", section0.Edition)
		}
	} else {
		return section0, fmt.Errorf("unsupported grib indicator %d", section0.Indicator)
	}

	return

}

// SectionHead is the common header for each section1-8
//
//	| Octet Number | Content
//	-----------------------------------------------------------
//	| 1-4          | Length of the section in octets (21 or N)
//	| 5            | Number of the section (1)
type SectionHead struct {
	ByteLength uint32 `json:"byteLength"`
	Number     uint8  `json:"number"`
}

// ReadSectionHead is poorly documented other than code
func ReadSectionHead(section io.Reader) (head SectionHead, err error) {
	var length uint32
	err = binary.Read(section, binary.BigEndian, &length)
	if err != nil {
		return head, fmt.Errorf("read of Length failed: %s", err.Error())
	}
	if length == EndSectionLength {
		return SectionHead{
			ByteLength: 4,
			Number:     8,
		}, nil
	}
	var sectionNumber uint8
	err = binary.Read(section, binary.BigEndian, &sectionNumber)
	if err != nil {
		return head, err
	}

	return SectionHead{
		ByteLength: length,
		Number:     sectionNumber,
	}, nil
}

// SectionNumber returns the number of the sectionhead
func (s SectionHead) SectionNumber() uint8 {
	return s.Number
}

// ContentLength returns the binary length of the sectionhead
func (s SectionHead) ContentLength() int {
	return int(s.ByteLength) - binary.Size(s)
}

func (s SectionHead) String() string {
	return fmt.Sprint("Section ", s.Number)
}

// Time is the time of section 1
//
//	| Octet Number | Content
//	---------------------------------
//	| 13-14        | Year (4 digits)
//	| 15           | Month
//	| 16           | Day
//	| 17           | Hour
//	| 18           | Minute
//	| 19           | Second
type Time struct {
	Year   uint16 `json:"year"`   // year
	Month  uint8  `json:"month"`  // month + 1
	Day    uint8  `json:"day"`    // day
	Hour   uint8  `json:"hour"`   // hour
	Minute uint8  `json:"minute"` // minute
	Second uint8  `json:"second"` // second
}

// Section1 is the Identification section http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_doc/grib2_sect1.shtml
//    | Octet Number | Content
//    -----------------------------------------------------------------------------------------
//    | 1-4          | Length of the section in octets (21 or N)
//    | 5            | Number of the section (1)
//    | 6-7          | Identification of originating/generating center (See Table 0) (See note 4)
//    | 8-9          | Identification of originating/generating subcenter (See Table C)
//    | 10           | GRIB master tables version number (currently 2) (See Table 1.0) (See note 1)
//    | 11           | Version number of GRIB local tables used to augment Master Tables (see Table 1.1)
//    | 12           | Significance of reference time (See Table 1.2)
//    | 13-14        | Year (4 digits)
//    | 15           | Month
//    | 16           | Day
//    | 17           | Hour
//    | 18           | Minute
//    | 19           | Second
//    | 20           | Production Status of Processed data in the GRIB message (See Table 1.3)
//    | 21           | Type of processed data in this GRIB message (See Table 1.4)
//    | 22-N         | Reserved
//
// Local tables define those parts of the master table which are reserved for local use except for the case described below. In any case,
// the use of local tables in the messages are intended for non-local or international exchange is strongly discouraged.
//
// If octet 10 is set to 255 then only local tables are in use.  In this case, the local table version number (octet 11) must not be zero
// nor missing.  Local tables may include entries from the entire range of the tables.
//
// If octet 11 is zero, octet 10 must contain a valid master table version number and only those parts of the tables not reserved for local
// use may be used.
//
// If octets 8-9 is zero, Not a sub-center, the originating/generating center is the center defined by octets 6-7.

type Section1 struct {
	OriginatingCenter         uint16 `json:"ooriginatingCenter"`
	OriginatingSubCenter      uint16 `json:"originatingSubCenter"`
	MasterTablesVersion       uint8  `json:"masterTablesVersion"`
	LocalTablesVersion        uint8  `json:"localTablesVersion"`
	ReferenceTimeSignificance uint8  `json:"referenceTimeSignificance"` // Table 1.2, value 1 is start of forecast
	ReferenceTime             Time   `json:"referenceTime"`
	ProductionStatus          uint8  `json:"productionStatus"`
	Type                      uint8  `json:"type"` // data type, Table 1.4, value 1 is forecast products
}

// ReadSection1 is poorly documented other than code
func ReadSection1(f io.Reader, length int) (section Section1, err error) {
	return section, binary.Read(f, binary.BigEndian, &section)
}

// Section2 is the Local Use section http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_doc/grib2_sect2.shtml
//
//	| Octet Number | Content
//	-----------------------------------------------------------------------------------------
//	| 1-4          | Length of the section in octets (N)
//	| 5            | Number of the section (2)
//	| 6-N          | Local Use
//
// Center=7 (NCEP), subcenter=14(NWS Meteorological Development Laboratory (MDL)) used octet 6 to indicate which local use table
// to use. For MDL, octet 6=1 indicates use: "MDL Template 2.1"
type Section2 struct {
	LocalUse []uint8 `json:"localUse"`
}

// ReadSection2 is for "Local use"
func ReadSection2(f io.Reader, len int) (section Section2, err error) {
	section.LocalUse = make([]uint8, len)
	return section, read(f, &section.LocalUse)
}

// Section3 is the Grid Definition section http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_doc/grib2_sect3.shtml
// It contains information of the grid(earth shape, long, lat, etc)
//
//	| Octet Number | Content
//	-----------------------------------------------------------------------------------------
//	| 1-4          | Length of the section in octets (nn)
//	| 5            | Number of the section (3)
//	| 6            | Source of grid definition (See Table 3.0) (See note 1 below)
//	| 7-10         | Number of data points
//	| 11           | Number of octets for optional list of numbers defining number of points (See note 2 below)
//	| 12           | Interpetation of list of numbers defining number of points (See Table 3.11)
//	| 13-14        | Grid definition template number (= N) (See Table 3.1)
//	| 15-xx        | Grid definition template (See Template 3.N, where N is the grid definition template
//	|              | number given in octets 13-14)
//	| [xx+1]-nn    | Optional list of numbers defining number of points (See notes 2, 3, and 4 below)
//
// If octet 6 is not zero, octets 15-xx (15-nn if octet 11 is zero) may not be supplied.  This should be documented with all bits set
// to 1 in the grid definition template number.
//
// An optional list of numbers defining number of points is used to document a quasi-regular grid, where the number of points may vary
// from one row to another.  In such a case, octet 11 is non zero and gives the number octets on which each number of points is encoded.
// For all other cases, such as regular grids, octets 11 and 12 are zero and no list is appended to the grid definition template.
//
// If a list of numbers defining the number of points is preset, it is appended at the end of the grid definition template ( or directly
// after the grid definition number if the template is missing).  When the grid definition template is present, the length is given
// according to bit 3 of the scanning mode flag octet (length is Nj or Ny for flag value 0).  List ordering is implied by data scanning.
//
// Depending on the code value given in octet 12, the list of numbers either:
//   - Corresponds to the coordinate lines as given in the grid definition, or
//   - Corresponds to a full circle, or
//   - Does not apply.
type Section3 struct {
	Source                   uint8       `json:"source"`
	DataPointCount           uint32      `json:"datapointCount"`
	PointCountOctets         uint8       `json:"pointCountOctets"`
	PointCountInterpretation uint8       `json:"pointCountInterpretation"`
	TemplateNumber           uint16      `json:"templateNumber"`
	Definition               interface{} `json:"definition"`
}

func (s Section3) String() string {
	return fmt.Sprint("Point count: ", s.DataPointCount, " Definition: ", GridDefinitionTemplateDescription(int(s.TemplateNumber)))
}

// ReadSection3 reads section3 from reader(f). the Lenght parameter is ignored
func ReadSection3(f io.Reader, _ int) (section Section3, err error) {

	err = read(f, &section.Source, &section.DataPointCount, &section.PointCountOctets, &section.PointCountInterpretation, &section.TemplateNumber)
	if err != nil {
		return section, err
	}

	//
	section.Definition, err = ReadGrid(f, section.TemplateNumber)
	return section, err
}

// Section4 is the Product Definition Section http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_doc/grib2_sect4.shtml
//
//	| Octet Number | Content
//	-----------------------------------------------------------------------------------------
//	| 1-4          | Length of the section in octets (nn)
//	| 5            | Number of the section (4)
//	| 6-7          | Number of coordinate values after template (See note 1 below)
//	| 8-9          | Product definition template number (See Table 4.0)
//	| 10-xx        | Product definition template (See product template 4.X, where X is
//	|              | the number given in octets 8-9)
//	| [xx+1]-nn    | Optional list of coordinate values (See notes 2 and 3 below)
//
// Coordinate values are intended to document the vertical discretization associated with model data on hybrid coordinate vertical
// levels.  A value of zero in octets 6-7 indicates that no such values are present.  Otherwise the number corresponds to the whole
// set of values.
//
// Hybrid systems employ a means of representing vertical coordinates in terms of a mathematical combination of pressure and sigma
// coordinates.  When used in conjunction with a surface pressure field and an appropriate mathematical expression, the vertical
// coordinate parameters may be used to interpret the hybrid vertical coordinate.
//
// Hybrid coordinate values, if present, should be encoded in IEEE 32-bit floating point format.  They are intended to be encoded as
// pairs.
type Section4 struct {
	CoordinatesCount                uint16   `json:"coordinatesCount"`
	ProductDefinitionTemplateNumber uint16   `json:"productDefinitionTemplateNumber"`
	ProductDefinitionTemplate       Product0 `json:"productDefinitionTemplate"` // FIXME, support more products
	Coordinates                     []byte   `json:"coordinates"`
}

// ReadSection4 reads section4 from an io.Reader
func ReadSection4(f io.Reader, length int) (section Section4, err error) {
	err = read(f, &section.CoordinatesCount, &section.ProductDefinitionTemplateNumber)
	if err != nil {
		return section, err
	}

	switch section.ProductDefinitionTemplateNumber {
	case 0:
		err = read(f, &section.ProductDefinitionTemplate)
	default:
		//return section, fmt.Errorf("Category definition template number %d not implemented yet", section.ProductDefinitionTemplateNumber)
		return section, nil
	}

	if err != nil {
		return section, err
	}

	section.Coordinates = make([]byte, section.CoordinatesCount)

	return section, read(f, &section.Coordinates)
}

// Section5 is Data Representation section http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_doc/grib2_sect5.shtml
//
//	| Octet Number | Content
//	-----------------------------------------------------------------------------------------
//	| 1-4          | Length of the section in octets (nn)
//	| 5            | Number of the section (5)
//	| 6-9          | Number of data points where one or more values are specified in Section 7 when a bit map is present,
//	|              | total number of data points when a bit map is absent.
//	| 10-11        | Data representation template number (See Table 5.0 http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_doc/grib2_table5-0.shtml)
//	| 12-nn        | Data representation template (See Template 5.X, where X is the number given in octets 10-11)
type Section5 struct {
	PointsNumber       uint32 `json:"pointsNumber"`
	DataTemplateNumber uint16 `json:"dataTemplateNumber"`
	//DataTemplate       Data3  `json:"dataTemplate"` // FIXME, support more data-types
	Data []byte `json:"dataTemplate"`
}

// ReadSection5 is poorly documented other than code
func ReadSection5(f io.Reader, length int) (section Section5, err error) {

	section.Data = make([]byte, length-6)

	err = read(f, &section.PointsNumber, &section.DataTemplateNumber, &section.Data)
	if err != nil {
		return section, err
	}

	if section.DataTemplateNumber != 0 && section.DataTemplateNumber != 2 && section.DataTemplateNumber != 3 {
		return section, fmt.Errorf("template number not supported: %d", section.DataTemplateNumber)
	}

	return section, nil
}

// GetDataTemplate extract DataTemplate from the section
func (section Section5) GetDataTemplate() (interface{}, error) {
	switch section.DataTemplateNumber {
	case 0:
		data := Data0{}
		err := read(bytes.NewReader(section.Data), &data)
		return data, err
	case 2:
		data := Data2{}
		err := read(bytes.NewReader(section.Data), &data)
		return data, err
	case 3:
		data := Data3{}
		err := read(bytes.NewReader(section.Data), &data)
		return data, err
	default:
		return struct{}{}, fmt.Errorf("unknown data format")
	}
}

// Section6 is the Bit-Map section http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_doc/grib2_sect6.shtml
//
//	| Octet Number | Content
//	-----------------------------------------------------------------------------------------
//	| 1-4          | Length of the section in octets (nn)
//	| 5            | Number of the section (6)
//	| 6            | Bit-map indicator (See Table 6.0) (See note 1 below)
//	| 7-nn         | Bit-map
//
// If octet 6 is not zero, the length of this section is 6 and octets 7-nn are not present.
type Section6 struct {
	BitmapIndicator uint8  `json:"bitmapIndicator"`
	Bitmap          []byte `json:"bitmap"`
}

// ReadSection6 is poorly documented other than code
func ReadSection6(f io.Reader, length int) (section Section6, err error) {
	section.Bitmap = make([]byte, length-1)

	return section, read(f, &section.BitmapIndicator, &section.Bitmap)
}

// Section7 is the Data section http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_doc/grib2_sect7.shtml
//
//	| Octet Number | Content
//	-----------------------------------------------------------------------------------------
//	| 1-4          | Length of the section in octets (nn)
//	| 5            | Number of the section (7)
//	| 6-nn         | Data in a format described by data Template 7.X, where X is the data representation template number
//	|              | given in octets 10-11 of Section 5.
type Section7 struct {
	Data []float64 `json:"data"`
}

// ReadSection7 reads the actual data
func ReadSection7(f io.Reader, length int, section5 Section5) (section Section7, sectionError error) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Corrupt message %q\n", err)
		}
	}()

	data, sectionError := section5.GetDataTemplate()

	if sectionError != nil {
		return Section7{}, sectionError
	}

	if length != 0 {

		switch x := data.(type) {
		case Data0:
			section.Data, sectionError = ParseData0(f, length, &x)
		case Data2:
			section.Data, sectionError = ParseData2(f, length, &x)
		case Data3:
			section.Data, sectionError = ParseData3(f, length, &x)
		default:
			sectionError = fmt.Errorf("unknown data type")
			return
		}
	}

	return section, sectionError
}

// read bytes from reader and serialize the bytes into the given data .. pointers
func read(reader io.Reader, data ...interface{}) (err error) {
	for _, what := range data {
		err = binary.Read(reader, binary.BigEndian, what)
		if err != nil {
			return err
		}
	}
	return nil
}
