package data

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

type ScaledValue struct {
	Scale uint8
	Value uint32
}

type BasicAngle struct {
	BasicAngle    uint32
	BasicAngleSub uint32
}

type Grid interface {
	Export() map[string]string
}

func ReadGrid(f io.Reader, templateNumber uint16) (Grid, error) {
	switch templateNumber {
	case 0:
		var grid Grid0
		return &grid, binary.Read(f, binary.BigEndian, &grid)
	case 10:
		var grid Grid10
		return &grid, binary.Read(f, binary.BigEndian, &grid)
	case 20:
		var grid Grid20
		return &grid, binary.Read(f, binary.BigEndian, &grid)
	case 30:
		var grid Grid30
		return &grid, binary.Read(f, binary.BigEndian, &grid)
	case 40:
		var grid Grid40
		return &grid, binary.Read(f, binary.BigEndian, &grid)
	case 90:
		var grid Grid90
		return &grid, binary.Read(f, binary.BigEndian, &grid)
	default:
		var grid Grid90
		return &grid, errors.New(fmt.Sprint("Unknown grid template number ", templateNumber))
	}
}

type GridHeader struct {
	EarthShape      uint8
	SphericalRadius ScaledValue
	MajorAxis       ScaledValue
	MinorAxis       ScaledValue
}

func (h *GridHeader) Export() (d map[string]string) {
	return map[string]string{
		"earth": ReadEarthShape(int(h.EarthShape)),
	}
}

// Grid Definition Template 3.0: Latitude/longitude (or equidistant cylindrical, or Plate Carree)
type Grid0 struct {
	//Name :=  "Latitude/longitude (or equidistant cylindrical, or Plate Carree) "
	GridHeader
	Ni                          uint32
	Nj                          uint32
	BasicAngle                  BasicAngle
	La1                         int32
	Lo1                         int32
	ResolutionAndComponentFlags uint8
	La2                         int32
	Lo2                         int32
	Di                          int32
	Dj                          int32
	ScanningMode                uint8
}

func (h *Grid0) Export() (d map[string]string) {
	return map[string]string{
		"earth":         ReadEarthShape(int(h.EarthShape)),
		"ni":            fmt.Sprint(h.Ni),
		"nj":            fmt.Sprint(h.Nj),
		"basicAngle":    fmt.Sprint(h.BasicAngle.BasicAngle),
		"basicAngleSub": fmt.Sprint(h.BasicAngle.BasicAngleSub),
		"la1":           fmt.Sprint(h.La1),
		"lo1":           fmt.Sprint(h.Lo1),
		"la2":           fmt.Sprint(h.La2),
		"lo2":           fmt.Sprint(h.Lo2),
		"di":            fmt.Sprint(h.Di),
		"dj":            fmt.Sprint(h.Dj),
		"scanningMode":  fmt.Sprint(h.ScanningMode),
	}
}

// Grid Definition Template 3.10: Mercator
type Grid10 struct {
	//name :=  "Mercator"
	GridHeader
	Ni                          uint32
	Nj                          int32
	La1                         int32
	Lo1                         int32
	ResolutionAndComponentFlags uint8
	Lad                         int32
	La2                         int32
	Lo2                         int32
	ScanningMode                uint8
	GridOrientation             uint32
	Di                          int32
	Dj                          int32
}

// Grid Definition Template 3.20: Polar stereographic projection
type Grid20 struct {
	//name =  "Polar stereographic projection ";
	GridHeader
	Nx                          uint32
	Ny                          uint32
	La1                         int32
	Lo1                         int32
	ResolutionAndComponentFlags uint8
	Lad                         int32
	Lov                         int32
	Dx                          int32
	Dy                          int32
	ProjectionCenter            uint8
	ScanningMode                uint8
}

// Grid Definition Template 3.30: Lambert conformal
type Grid30 struct {
	//name =  "Polar stereographic projection ";
	GridHeader
	Nx                          uint32
	Ny                          uint32
	La1                         int32
	Lo1                         int32
	ResolutionAndComponentFlags uint8
	Lad                         int32
	Lov                         int32
	Dx                          int32
	Dy                          int32
	ProjectionCenter            uint8
	ScanningMode                uint8
	Latin1                      uint32
	Latin2                      uint32
	LaSouthPole                 uint32
	LoSouthPole                 uint32
}

// Grid Definition Template 3.40: Gaussian latitude/longitude
type Grid40 struct {
	//name =  "Gaussian latitude/longitude ";
	GridHeader
	Ni                          uint32
	Nj                          uint32
	BasicAngle                  uint32
	La1                         int32
	Lo1                         int32
	ResolutionAndComponentFlags uint8
	La2                         int32
	Lo2                         int32
	Di                          int32
	N                           uint32
	ScanningMode                uint8
}

// Grid Definition Template 3.90: Space view perspective or orthographic
// FIXME: implement properly
type Grid90 struct {
	//name =  "Space view perspective or orthographic ";
	GridHeader
	Nx uint32
	Ny uint32
	//BasicAngle                  BasicAngle
	Lap                         int32
	Lop                         int32
	ResolutionAndComponentFlags uint8

	Dx uint32
	Dy uint32

	Xp uint32
	Yp uint32

	// fix byte size
	ScanningMode uint8
	Orientation  uint32
	Nr           uint32

	Xo uint32
	Yo uint32
}
