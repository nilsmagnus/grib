# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a GRIB2 (GRIdded Binary) file format parser written in Go. GRIB2 is a standard format for meteorological data. The library parses GRIB2 files from sources like NOAA GFS (Global Forecast System) and can be used both as a library and as a CLI tool.

## Build Commands

```bash
make              # Run tests, lint, and install
make build        # Build the binary
make test         # Run tests (short mode)
make lint         # Run go vet
make test-all     # Run lint + tests
make fmt          # Format code with gofmt
```

Run a single test:
```bash
go test -v ./griblib/gribtest -run TestName
```

Run benchmarks:
```bash
cd griblib/gribtest && make benchmark
```

## Architecture

### Message Structure

A GRIB2 file contains multiple **Messages**, each representing a data layer. Each Message has 8 sections (Section0-Section7):

- **Section0**: Indicator section - identifies GRIB format, discipline (meteorology, hydrology, etc.), message length
- **Section1**: Identification - originating center, reference time, production status
- **Section2**: Local use (optional)
- **Section3**: Grid definition - earth shape, lat/long, number of data points (uses `Grid0` for latitude/longitude grids)
- **Section4**: Product definition - what the data represents (temperature, wind, etc.) via `Product0`
- **Section5**: Data representation - how data is packed (template number determines packing method)
- **Section6**: Bit-map section
- **Section7**: Actual data values (unpacked into `[]float64`)

### Data Packing Templates

The library supports three data representation templates (Section5):

- **Data0** (`data0.go`): Simple packing - basic binary scale/decimal scale
- **Data2** (`data2.go`): Complex packing - group splitting method with bit groups
- **Data3** (`data3.go`): Complex packing with spatial differencing - extends Data2 with first/second order differencing to improve compression

Data2 and Data3 use **bit groups** for efficient encoding of meteorological fields.

### Key Packages

- `griblib/` - Main library code: sections parsing, data templates, filters, exports
- `internal/reader/` - BitReader for reading variable-bit-width integers from binary streams
- `griblib/gribtest/` - Tests and benchmarks

### Entry Points

**Library usage** (`griblib/sections.go`):
- `ReadMessages(io.Reader)` - Read all messages from a GRIB2 file
- `ReadNMessages(io.Reader, n)` - Read first n messages
- `ReadMessage(io.Reader)` - Read single message

**CLI** (`main.go`):
- `parse` operation: Read and filter messages, export to JSON
- `reduce` operation: Filter and write reduced GRIB2 file

### Filtering

`griblib/filters.go` supports filtering by:
- Discipline (meteorology=0, hydrology=1, etc.)
- Category (temperature=0, moisture=1, etc.)
- Surface type
- Geographic bounds (GeoFilter with lat/long in millionths of degrees)

## Dependencies

The library intentionally has no runtime dependencies. Test dependencies are `github.com/stretchr/testify` and `github.com/golang/mock`.

## External Documentation

- GRIB2 specification: https://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_doc/
- Sample data: http://www.ftp.ncep.noaa.gov/data/nccf/com/gfs/prod/
