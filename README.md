Circle-ci build: [![CircleCI](https://circleci.com/gh/nilsmagnus/grib/tree/master.svg?style=svg)](https://circleci.com/gh/nilsmagnus/grib/tree/master)

GRIB2 Golang parser application and library
================================

Parser and library for grib2 file format. 

Forked from github.com/analogic/grib which is now abandoned by the author (see [comment on my pull request](https://github.com/analogic/grib/pull/1)).

## Usage

Install by typing

    go get -u github.com/nilsmagnus/grib


### Application Usage:

    $ grib -h 
    
    Usage of grib:
     -category int
       	Filters on Category within discipline. -1 means all categories (default -1)
     -dataExport
       	Export data values. (default true)
     -discipline int
       	Filters on Discipline. -1 means all disciplines (default -1)
     -export int
       	Export format. Valid types are 0 (none) 1(print discipline names) 2(print categories) 3(json) 
     -file string
       	Grib filepath
     -latMax int
       	Maximum latitude multiplied with 100000. (default 36000000)
     -latMin int
       	Minimum latitude multiplied with 100000.
     -longMax int
       	Maximum longitude multiplied with 100000. (default 9000000)
     -longMin int
       	Minimum longitude multiplied with 100000. (default -9000000)
     -maxmsg int
       	Maximum number of messages to parse. Does not work in combination with filters. (default 2147483647)
     -operation string
       	Operation. Valid values: 'parse', 'reduce'. (default "parse")
     -reducefile string
       	Destination for reduced file. (default "reduced.grib2")

#### Examples:

Reduce input file to default output-file with discipline 0 (Meteorology):

    grib -operation reduce -file testdata/reduced.grib2 -discipline 0

Filter on area on size of norway+sweden, output to json:
      
    grib -file testdata/gfs.t00z.pgrb2.2p50.f003  -latMin 57000000 -latMax 71000000 -longMin 4400000 -longMax 32000000 -export 3

Filter on temperature only:

    grib -file testdata/gfs.t00z.pgrb2.2p50.f003 -discipline 0 -category 0 

## Library examples

Have a look at 'main.go' for main usage:

    gribfile, err := os.Open("somegrib2file.grib2")
	if err != nil {
		b.Fatalf("Could not open test-file %v", err)
	}
    griblib.ReadMessages(gribfile)
	

## What works?

- basic binary parsing of GRIB2 GFS files from NOAA
- implemented only "Grid point data - complex packing and spatial differencing"
- Parsing Data3 type
- Parsing Data0 type (Thanks to Cyrille Meichel)

## Development

### Dependencies

grib itself has _no dependencies_ and wants to stay that way to keep it simple. Therefore, there is no dep/glide and no dependency-configuration-files.

### Build

Grib uses make to build. You probably need go-lint installed in order to build.

To build, simply type

    $ make
    
to test:

    $ make test


## TODOs

- Support different types of grids, not only grid0
- Support different types of products, not only product0
- Tests for reduction
- Tests for reading all sections


# Contributors

 - "analogic", https://github.com/analogic
 - Cyrille Meichel, https://github.com/landru29
 - Nils Larsgård, https://github.com/nilsmagnus
 
## Help appreciated

Feel free to fork and submit pull requests or simply create issues for improvements :)

### Performance


to run a performance-test go to griblib/gribtest and run `make benchmark`. 

#### CPU
Sample output:

```
File: gribtest.test
Type: cpu
Time: Feb 7, 2019 at 9:10pm (CET)
Duration: 9.31s, Total samples = 9.54s (102.47%)
Showing nodes accounting for 8.64s, 90.57% of 9.54s total
Dropped 126 nodes (cum <= 0.05s)
      flat  flat%   sum%        cum   cum%
     1.37s 14.36% 14.36%      2.14s 22.43%  github.com/nilsmagnus/grib/griblib.(*BitReader).readBit
     1.23s 12.89% 27.25%      3.37s 35.32%  github.com/nilsmagnus/grib/griblib.(*BitReader).readUint
     0.72s  7.55% 34.80%      0.72s  7.55%  runtime.memmove
     0.72s  7.55% 42.35%      0.80s  8.39%  syscall.Syscall
     0.54s  5.66% 48.01%      0.54s  5.66%  runtime.memclrNoHeapPointers
     0.48s  5.03% 53.04%      0.48s  5.03%  github.com/nilsmagnus/grib/griblib.(*BitReader).currentBit (inline)
     0.38s  3.98% 57.02%      0.89s  9.33%  encoding/binary.(*decoder).value
     0.31s  3.25% 60.27%      3.97s 41.61%  github.com/nilsmagnus/grib/griblib.(*BitReader).readIntsBlock
 ```

* As you can see from the sample output, it is the readBit and readUint that takes most of the time. 
* Performance can be better if you only want to read certain kinds of messages. The filter that is implemented will filter only after reading all the messages. (help appreciated :) )

#### Memory


```
go tool pprof -top memprofile.out
File: gribtest.test
Type: alloc_space
Time: Feb 7, 2019 at 9:20pm (CET)
Showing nodes accounting for 9.31GB, 99.36% of 9.37GB total
Dropped 44 nodes (cum <= 0.05GB)
      flat  flat%   sum%        cum   cum%
    4.83GB 51.57% 51.57%     5.94GB 63.37%  github.com/nilsmagnus/grib/griblib.(*Data2).extractData
    2.39GB 25.47% 77.04%     2.39GB 25.47%  github.com/nilsmagnus/grib/griblib.(*Data2).scaleValues
    0.98GB 10.47% 87.50%     0.98GB 10.47%  github.com/nilsmagnus/grib/griblib.(*BitReader).readIntsBlock
    0.30GB  3.20% 90.71%     9.36GB 99.84%  github.com/nilsmagnus/grib/griblib.ReadMessage
    0.25GB  2.67% 93.38%     0.49GB  5.22%  github.com/nilsmagnus/grib/griblib.(*Data2).extractBitGroupParameters

```

* extractData is hogging memory
* scaleValues is hoggig memory


# Grib Documentation

Grib specification:

http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_doc.shtml

Documentation from noaa.gov :

http://www.nco.ncep.noaa.gov/pmb/docs/on388/


Daily grib2 files from NOAA can be found at

http://www.ftp.ncep.noaa.gov/data/nccf/com/gfs/prod/
