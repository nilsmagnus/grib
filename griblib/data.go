package griblib

import (
	"bytes"
	"fmt"
	"io"
	"math"
)

//Data3 is a Grid point data - complex packing and spatial differencing
// http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_temp5-3.shtml
type Data3 struct {
	Reference              float32 `json:"reference"`
	BinaryScale            uint16  `json:"binaryScale"`
	DecimalScale           uint16  `json:"decimalScale"`
	Bits                   uint8   `json:"bits"`
	Type                   uint8   `json:"type"`
	GroupMethod            uint8   `json:"groupMethod"`
	MissingValue           uint8   `json:"missingValue"`
	MissingSubstitute1     uint32  `json:"missingSubstitute1"`
	MissingSubstitute2     uint32  `json:"missingSubstitute2"`
	NG                     uint32  `json:"ng"`
	GroupWidths            uint8   `json:"groupWidths"`
	GroupWidthsBits        uint8   `json:"groupWidthsBits"`
	GroupLengthsReference  uint32  `json:"groupLengthsReference"` // 13
	GroupLengthIncrement   uint8   `json:"groupLengthIncrement"`  // 14
	GroupLastLength        uint32  `json:"groupLastLength"`       // 15
	GroupScaledLengthsBits uint8   `json:"groupScaledLengthsBits"`
	SpatialOrderDifference uint8   `json:"spatialOrderDifference"`
	OctetsNumber           uint8   `json:"octetsNumber"`
}

// ParseData3 parses data3 struct from the reader into the template
func ParseData3(dataReader io.Reader, dataLength int, data3 *Data3) []float64 {

	rawData := make([]byte, dataLength)
	dataReader.Read(rawData)
	reader := newBitReader(bytes.NewReader(rawData))
	var ival1, ival2, minsd uint64

	os := data3.SpatialOrderDifference
	nbitsd := data3.OctetsNumber
	var sign uint64

	nbitsd = nbitsd * 8
	if nbitsd > 0 { // first order spatial differencing g1 and gMin
		sign = reader.readBits64(1)
		ival1 = reader.readBits64(uint(nbitsd) - 1)
		if sign == 1 {
			ival1 = -ival1
		}
		if os == 2 { //second order spatial differencing h1, h2, hMin
			sign = reader.readBits64(1)
			ival2 = reader.readBits64(uint(nbitsd) - 1)
			if sign == 1 {
				ival2 = -ival2
			}
		}
		sign = reader.readBits64(1)
		minsd = reader.readBits64(uint(nbitsd) - 1)
		if sign == 1 {
			minsd = -minsd
		}

	} else {
		fmt.Println("unsupported spatial differencing, returning zero-data")
		// TODO fill with missing value
		return make([]float64, dataLength) // missing data
	}

	if data3.NG == 0 {
		fmt.Println("no groups, returning zero-data")
		// TODO fill with missing

		return make([]float64, dataLength) // missing data
	}

	// [ww +1]-xx  Get reference values for groups (X1's)
	// X1 == gref
	X1 := make([]uint64, data3.NG) // initialized to zero
	if data3.Bits != 0 {
		reader.incrByte()
		for i := 0; i < int(data3.NG); i++ {
			X1[i] = reader.readBits64(uint(data3.Bits))
		}
	}

	// [xx +1 ]-yy Get number of bits used to encode each group
	// NB == gwidth
	NB := make([]uint64, data3.NG) // initialized to zero
	if data3.GroupWidthsBits != 0 {
		reader.incrByte()
		for i := 0; i < int(data3.NG); i++ {
			NB[i] = reader.readBits64(uint(data3.GroupWidthsBits))
		}
	}

	for i := 0; i < int(data3.NG); i++ {
		NB[i] += uint64(data3.GroupWidths)
	}

	// [yy +1 ]-zz Get the scaled group lengths using formula
	//     Ln = ref + Kn * len_inc, where n = 1-NG,
	//          ref = referenceGroupLength, and  len_inc = lengthIncrement

	L := make([]uint64, data3.NG) // initialized to zero
	referenceGroupLength := data3.GroupLengthsReference
	nb := data3.GroupScaledLengthsBits
	if nb != 0 {
		reader.incrByte()
		for i := 0; i < int(data3.NG); i++ {
			L[i] = reader.readBits64(uint(nb))
		}
	}

	var totalL uint64
	for i := 0; i < int(data3.NG); i++ {
		L[i] = L[i]*uint64(data3.GroupLengthIncrement) + uint64(referenceGroupLength)
		totalL += L[i]
	}
	totalL -= L[data3.NG-1]
	totalL += uint64(data3.GroupLastLength)

	//enter Length of Last Group
	L[data3.NG-1] = uint64(data3.GroupLastLength)

	// self test
	//TODO add self test?

	data := make([]float64, dataLength)

	count := 0
	reader.incrByte()
	if data3.MissingValue == 0 {
		for i := 0; i < int(data3.NG); i++ {

			if NB[i] != 0 {
				for j := 0; j < int(L[i]); j++ {
					count++
					data[count] = float64(reader.readBits64(uint(NB[i])) + X1[i])
				}
			} else {
				for j := 0; j < int(L[i]); j++ {
					count++
					data[count] = float64(X1[i])
				}
			}
		} // end for i

	} else {
		fmt.Println("Missingvalue 1 and 2 not supported, returing zero-values")
		return make([]float64, dataLength)
	}

	if os == 1 { // g1 and gMin
		// encoded by G(n) = F(n) - F(n -1 )
		// decoded by F(n) = G(n) + F(n -1 )
		// data[] at this point contains G0, G1, G2, ....
		data[0] = float64(ival1)
		for i := 0; i < dataLength; i++ {
			data[i] += float64(minsd)
			data[i] = data[i] + data[i-1]
		}
	} else if os == 2 { // 2nd order
		data[0] = float64(ival1)
		data[1] = float64(ival2)
		for i := 0; i < dataLength; i++ {
			data[i] += float64(minsd)
			data[i] = data[i] + (2 * data[i-1]) - data[i-2]
		}
	}

	D := data3.DecimalScale
	E := data3.BinaryScale
	R := float64(data3.Reference)
	DD := math.Pow(10, float64(D))
	EE := math.Pow(2, float64(E))
	for i := 0; i < dataLength; i++ {
		data[i] = (R + (data[i] * EE)) / DD
	}

	return data

}
func ParseData3_old(dataReader io.Reader, dataLength int, template *Data3) []float64 {

	rawData := make([]byte, dataLength)
	dataReader.Read(rawData)

	var ival1 int64
	var ival2 int64
	var minsd uint64
	var err error

	var missingValueSubstitute1 float64
	var missingValueSubstitute2 float64

	ng := int(template.NG)
	if template.MissingValue == 1 {
		missingValueSubstitute1 = float64(template.MissingSubstitute1)
	} else if template.MissingValue == 2 {
		missingValueSubstitute1 = float64(template.MissingSubstitute1)
		missingValueSubstitute2 = float64(template.MissingSubstitute1)
	}

	//
	// Init reader
	//

	buffer := bytes.NewBuffer(rawData)

	bitReader := newReader(buffer)

	//
	//  Extract Spatial differencing values, if using DRS Template 5.3
	//
	rc := int(template.OctetsNumber) * 8
	if rc != 0 {
		ival1, err = bitReader.readInt(rc)

		if template.SpatialOrderDifference == 2 {
			ival2, err = bitReader.readInt(rc)
		}

		minsd, err = bitReader.readUint(rc)
	}

	//fmt.Println(" ival1", ival1, "ival2", ival2, "minsd", minsd, "octed bytes", rc/8)
	//fmt.Println("")
	//d.Seek(int64(int(dataLength) - oc * 3), 1)
	//return
	//os.Exit(0)

	//
	//  Extract Each Group's reference value
	//
	// fmt.Println("groups", template.NG)
	references, err := bitReader.readUintsBlock(int(template.Bits), ng, true)
	if err != nil {
		panic(err)
	}

	//fmt.Println("references")

	//
	//  Extract Each Group's bit width
	//
	widths, err := bitReader.readUintsBlock(int(template.GroupWidthsBits), ng, true)
	if err != nil {
		panic(err)
	}

	for j := 0; j < ng; j++ {
		widths[j] += uint64(template.GroupWidths)
	}

	//
	//  Extract Each Group's length (number of values in each group)
	//
	lengths, err := bitReader.readUintsBlock(int(template.GroupScaledLengthsBits), ng, true)
	if err != nil {
		panic(err)
	}

	for j := 0; j < ng; j++ {
		lengths[j] = (lengths[j] * uint64(template.GroupLengthIncrement)) + uint64(template.GroupLengthsReference)
	}
	lengths[template.NG-1] = uint64(template.GroupLastLength)

	// debug
	for j := 0; j < ng; j++ {
		//fmt.Println(j, " - Reference", references[j], "Width", widths[j], "length", lengths[j])
	}

	//
	//  Test to see if the group widths and lengths are consistent with number of
	//  values, and length of section 7.
	//

	totBit := 0
	totLen := 0

	for j := 0; j < ng; j++ {
		totBit += int(widths[j]) * int(lengths[j])
		totLen += int(lengths[j])
	}

	//if (totLen != ndpts) {
	//   panic("Checksum err")
	//}
	//fmt.Println(totBit / 8)
	//fmt.Println(dataLength)
	if totBit/8 > int(dataLength) {
		fmt.Println(totLen)
		panic("Checksum err")
	}

	//
	//  For each group, unpack data values
	//
	non := 0
	var section7Data []int64
	var ifldmiss []int64

	if template.MissingValue == 0 {
		n := 0
		for j := 0; j < ng; j++ {
			if widths[j] != 0 {
				tmp, _ := bitReader.readIntsBlock(int(widths[j]), int(lengths[j]))
				section7Data = append(section7Data, tmp...)

				for k := 0; k < int(lengths[j]); k++ {
					section7Data[n] = section7Data[n] + int64(references[j])
					n++
				}
			} else {
				for l := n; l < n+int(lengths[j]); l++ {
					section7Data = append(section7Data, int64(references[j]))
				}
				n = n + int(lengths[j])
			}
		}

	} else if template.MissingValue == 1 || template.MissingValue == 2 {
		// missing values included
		n := 0
		for j := 0; j < ng; j++ {
			if widths[j] != 0 {
				msng1 := math.Pow(2.0, float64(widths[j])) - 1
				msng2 := msng1 - 1

				ifldmiss, err = bitReader.readIntsBlock(int(widths[j]), int(lengths[j]))

				for k := 0; k < int(lengths[j]); k++ {
					if section7Data[n] == int64(msng1) {
						ifldmiss[n] = 1
						//section7Data[n]=0
					} else if template.MissingValue == 2 && section7Data[n] == int64(msng2) {
						ifldmiss[n] = 2
						//section7Data[n]=0
					} else {
						ifldmiss[n] = 0
						//section7Data[non++]=section7Data[n]+references[j]
					}
					n++
				}
			} else {
				msng1 := math.Pow(2.0, float64(template.Bits)) - 1
				msng2 := msng1 - 1
				if references[j] == uint64(msng1) {
					for l := n; l < n+int(lengths[j]); l++ {
						ifldmiss[l] = 1
					}
				} else if template.MissingValue == 2 && references[j] == uint64(msng2) {
					for l := n; l < n+int(lengths[j]); l++ {
						ifldmiss[l] = 2
					}
				} else {
					for l := n; l < n+int(lengths[j]); l++ {
						ifldmiss[l] = 0
					}
					for l := non; l < non+int(lengths[j]); l++ {
						section7Data[l] = int64(references[j])
					}
					non += int(lengths[j])
				}
				n = n + int(lengths[j])
			}
		}
	}

	/*
	   if ( references != 0 ) free(references)
	   if ( widths != 0 ) free(widths)
	   if ( lengths != 0 ) free(lengths)
	*/
	//
	//  If using spatial differences, add overall min value, and
	//  sum up recursively
	//
	//printf("SAGod: %ld %ld\n",idrsnum,idrstmpl[16])

	itemp := non
	ndpts := ng

	if template.SpatialOrderDifference == 1 {
		// first order
		section7Data[0] = int64(ival1)

		if template.MissingValue == 0 {
			itemp = ndpts // no missing values
		}

		for n := 1; n < itemp; n++ {
			section7Data[n] = section7Data[n] + int64(minsd)
			section7Data[n] = section7Data[n] + section7Data[n-1]
		}
	} else if template.SpatialOrderDifference == 2 {
		// second order
		section7Data[0] = int64(ival1)
		section7Data[1] = int64(ival2)
		if template.MissingValue == 0 {
			itemp = ndpts
		}

		for n := 2; n < itemp; n++ {
			section7Data[n] = section7Data[n] + int64(minsd)
			section7Data[n] = section7Data[n] + (2 * section7Data[n-1]) - section7Data[n-2]
		}
	}

	fld := make([]float64, len(section7Data))
	bscale := math.Pow(2.0, float64(template.BinaryScale))
	dscale := math.Pow(10.0, float64(template.DecimalScale))

	if template.MissingValue == 0 {
		// no missing values
		for i, dataValue := range section7Data {
			fld[i] = ((float64(dataValue) * float64(bscale)) + float64(template.Reference)) / dscale
		}
	} else if template.MissingValue == 1 || template.MissingValue == 2 {
		// missing values included
		non := 0
		for n, dataValue := range section7Data {
			if ifldmiss[n] == 0 {
				non++
				fld[n] = ((float64(dataValue) * float64(bscale)) + float64(template.Reference)) / dscale

				//printf(" SAG %d %f %d %f %f %f\n",n,fld[n],section7Data[non-1],bscale,ref,dscale)
			} else if ifldmiss[n] == 1 {
				fld[n] = missingValueSubstitute1
			} else if ifldmiss[n] == 2 {
				fld[n] = missingValueSubstitute2
			}

		}
	}

	return fld
}
