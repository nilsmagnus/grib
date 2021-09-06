package gribtest

import (
	"github.com/nilsmagnus/grib/griblib"
	"testing"
)

func TestCalculateAverageValue_0_values(t *testing.T) {

	filter := griblib.GeoFilter{MinLong: 10_000_000, MinLat: 10_000_000, MaxLat: 20_000_000, MaxLong: 20_000_000}
	grid := griblib.Grid0{Di: 2_500_000, Dj: 2_500_000}
	data := make([]float64, 100)

	calculatedValue, err := griblib.AverageValueBasic(filter, &grid, data)

	if err != nil {
		t.Fatalf("Error calculating value: %v", err)
	}

	if calculatedValue != 0.0 {
		t.Errorf("Average value should have been 0.0, was %f", calculatedValue)
	}
}

func TestCalculateAverageValue_incrementing_values(t *testing.T) {

	filter := griblib.GeoFilter{MinLong: 10_000_000, MinLat: 85_000_000, MaxLat: 70_000_000, MaxLong: 15_000_000}

	grid := griblib.Grid0{Di: 2_500_000, Dj: 2_500_000, Lo1: 0, Lo2: 357_500_000, La1: 90_000_000, La2: -2057483648, Ni: 144, Nj: 73}
	data := make([]float64, grid.Ni*grid.Nj)

	for i := 0; i < int(grid.Ni); i++ {
		for j := 0; j < int(grid.Nj); j++ {
			index := i*int(grid.Nj) + j
			data[index] = float64(index)
		}
	}
	/*
		[]data is  now
			0	1	2	3	4	5	6	7	8	9	...
			10	11	12	13	14	15	16	17	18	19	...
			20	21	22	23'	24'	25	26	27	28	29	...
			30	31	32	33'	34'	35	36	37	38	39	...
			40	41	42	43	44	45	46	47	48	49	...
			50	51	52	53	54	55	56	57	58	59	...
			60	61	62	63	64	65	66	67	68	69	...
			70	71	72	73	74	75	76	77	78	79	...
			80	81	82	83	84	85	86	87	88	89	...
			90	91	92	93	94	95	96	97	98	99	...
			...
	*/

	calculatedValue, err := griblib.AverageValueBasic(filter, &grid, data)

	if err != nil {
		t.Fatalf("Error calculating value: %v", err)
	}

	if calculatedValue != 333 {
		t.Errorf("Average value should have been 333, was %f", calculatedValue)
	}
}
