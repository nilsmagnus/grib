GRIB2 Golang experimental parser
================================

Unfinished dirty parser for meteo data

### Usage

Install by typing

    go get -u github.com/nilsmagnus/grib

Usage:

    Usage of grib:
      -export int
        	Export format. Valid types are 0 (none) 1 (json)
      -file string
        	Grib filename
      -maxmsg int
        	Maximum number of messages to parse. (default 2147483647)


### What works?

- basic binary parsing of GRIB2 GFS files from NOAA
- implemented only "Grid point data - complex packing and spatial differencing"

### TODO?

- implement and test output values
- refactor
- add some kind of tool for exporting values for world/place to json

# Grib Documentation

Grib specification:

http://www.wmo.int/pages/prog/www/WMOCodes/Guides/GRIB/GRIB2_062006.pdf

Documentation from noaa.gov :

http://www.nco.ncep.noaa.gov/pmb/docs/on388/


Examples can be found at

http://www.ftp.ncep.noaa.gov/data/nccf/com/gfs/prod/
