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
	Reference              float32
	BinaryScale            uint16
	DecimalScale           uint16
	Bits                   uint8
	Type                   uint8
	GroupMethod            uint8
	MissingValue           uint8
	MissingSubstitute1     uint32
	MissingSubstitute2     uint32
	NG                     uint32
	GroupWidths            uint8
	GroupWidthsBits        uint8
	GroupLengthsReference  uint32 // 13
	GroupLengthIncrement   uint8  // 14
	GroupLastLength        uint32 // 15
	GroupScaledLengthsBits uint8
	SpatialOrderDifference uint8
	OctetsNumber           uint8
}

func ParseData3(dataReader io.Reader, dataLength int, template *Data3) []int64 {

	rawData := make([]byte, dataLength)
	dataReader.Read(rawData)

	var ival1 int64
	var ival2 int64
	var minsd uint64
	var err error

	var rmiss1 float64
	var rmiss2 float64

	ng := int(template.NG)

	if template.MissingValue == 1 {
		rmiss1 = float64(template.MissingSubstitute1)
	} else if template.MissingValue == 2 {
		rmiss1 = float64(template.MissingSubstitute1)
		rmiss2 = float64(template.MissingSubstitute1)
	}

	//
	// Init reader
	//

	b := bytes.NewBuffer(rawData)

	r := NewReader(b)

	//
	//  Extract Spatial differencing values, if using DRS Template 5.3
	//
	rc := int(template.OctetsNumber) * 8
	if rc != 0 {
		ival1, err = r.ReadInt(rc)

		if template.SpatialOrderDifference == 2 {
			ival2, err = r.ReadInt(rc)
		}

		minsd, err = r.ReadUint(rc)
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
	references, err := r.readUintsBlock(int(template.Bits), ng, true)
	if err != nil {
		panic(err)
	}

	//fmt.Println("references")

	//
	//  Extract Each Group's bit width
	//
	widths, err := r.readUintsBlock(int(template.GroupWidthsBits), ng, true)
	if err != nil {
		panic(err)
	}

	for j := 0; j < ng; j++ {
		widths[j] += uint64(template.GroupWidths)
	}

	//
	//  Extract Each Group's length (number of values in each group)
	//
	lengths, err := r.readUintsBlock(int(template.GroupScaledLengthsBits), ng, true)
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
	//ifld := make([]int64, template.NG) // ndpts originalne
	var ifld []int64
	var ifldmiss []int64

	if template.MissingValue == 0 {
		n := 0
		for j := 0; j < ng; j++ {
			if widths[j] != 0 {

				//fmt.Println("reading", int(widths[j]), "bits", int(lengths[j]), "times")
				tmp, _ := r.readIntsBlock(int(widths[j]), int(lengths[j]), false)
				ifld = append(ifld, tmp...)

				//fmt.Println("----> ")
				//fmt.Println(int(widths[j]), int(lengths[j]))
				//fmt.Println(ifld)

				for k := 0; k < int(lengths[j]); k++ {
					ifld[n] = ifld[n] + int64(references[j])
					n++
				}
			} else {

				//fmt.Println(n)
				for l := n; l < n+int(lengths[j]); l++ {
					ifld = append(ifld, int64(references[j]))
				}
				n = n + int(lengths[j])
			}
		}

	} else if template.MissingValue == 1 || template.MissingValue == 2 {
		// missing values included

		n := 0

		for j := 0; j < ng; j++ {
			//printf(" SAGNGP %d %d %d %d\n",j,widths[j],lengths[j],references[j]);

			if widths[j] != 0 {
				msng1 := math.Pow(2.0, float64(widths[j])) - 1
				msng2 := msng1 - 1

				ifldmiss, err = r.readIntsBlock(int(widths[j]), int(lengths[j]), false)

				for k := 0; k < int(lengths[j]); k++ {
					if ifld[n] == int64(msng1) {
						ifldmiss[n] = 1
						//ifld[n]=0;
					} else if template.MissingValue == 2 && ifld[n] == int64(msng2) {
						ifldmiss[n] = 2
						//ifld[n]=0;
					} else {
						ifldmiss[n] = 0
						//ifld[non++]=ifld[n]+references[j];
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
						ifld[l] = int64(references[j])
					}
					non += int(lengths[j])
				}
				n = n + int(lengths[j])
			}
		}
	}

	/*
	   if ( references != 0 ) free(references);
	   if ( widths != 0 ) free(widths);
	   if ( lengths != 0 ) free(lengths);
	*/
	//
	//  If using spatial differences, add overall min value, and
	//  sum up recursively
	//
	//printf("SAGod: %ld %ld\n",idrsnum,idrstmpl[16]);

	itemp := non
	ndpts := ng

	if template.SpatialOrderDifference == 1 {
		// first order
		ifld[0] = int64(ival1)

		if template.MissingValue == 0 {
			itemp = ndpts // no missing values
		}

		for n := 1; n < itemp; n++ {
			ifld[n] = ifld[n] + int64(minsd)
			ifld[n] = ifld[n] + ifld[n-1]
		}
	} else if template.SpatialOrderDifference == 2 {
		// second order
		ifld[0] = int64(ival1)
		ifld[1] = int64(ival2)
		if template.MissingValue == 0 {
			itemp = ndpts
		}

		for n := 2; n < itemp; n++ {
			ifld[n] = ifld[n] + int64(minsd)
			ifld[n] = ifld[n] + (2 * ifld[n-1]) - ifld[n-2]
		}
	}

	//
	//  Scale data back to original form
	//
	//printf("SAGT: %f %f %f\n",ref,bscale,dscale);
	fld := make([]float64, ng)
	bscale := math.Pow(2.0, float64(template.BinaryScale))
	dscale := math.Pow(10.0, -float64(template.DecimalScale))

	if template.MissingValue == 0 {
		// no missing values
		for n := 0; n < ndpts; n++ {
			fld[n] = ((float64(ifld[n]) * float64(bscale)) + float64(template.Reference)) * dscale
		}
	} else if template.MissingValue == 1 || template.MissingValue == 2 {
		// missing values included
		non := 0
		for n := 0; n < ndpts; n++ {

			if ifldmiss[n] == 0 {
				non++
				fld[n] = ((float64(ifld[non]) * float64(bscale)) + float64(template.Reference)) * dscale

				//printf(" SAG %d %f %d %f %f %f\n",n,fld[n],ifld[non-1],bscale,ref,dscale);
			} else if ifldmiss[n] == 1 {
				fld[n] = rmiss1
			} else if ifldmiss[n] == 2 {
				fld[n] = rmiss2
			}

		}
	}

	return ifld
}
