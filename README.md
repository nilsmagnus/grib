Circle-ci build: [![CircleCI](https://circleci.com/gh/nilsmagnus/grib/tree/master.svg?style=svg)](https://circleci.com/gh/nilsmagnus/grib/tree/master)

GRIB2 Golang parser application and library
================================

Parser and library for grib2 file format. 

Forked from github.com/analogic/grib which is now abandoned by the author (see [comment on my pull request](https://github.com/analogic/grib/pull/1)).

## Usage

Install by typing

    go get -u github.com/nilsmagnus/grib

## Library Usage:

Have a look at 'main.go' for main usage:

    gribfile, err := os.Open("somegrib2file.grib2")
	if err != nil {
		b.Fatalf("Could not open test-file %v", err)
	}
    messages, err := griblib.ReadMessages(gribfile)
    
    for _, message := range messages {
    	// do your thing
    }

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
go version go1.11.2 linux/amd64
go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out
category number 0,parameter number 0,surface type 1, surface value 0 max: 329.500000 min: 197.900000
goos: linux
goarch: amd64
pkg: github.com/nilsmagnus/grib/griblib/gribtest
BenchmarkReadMessages-4   	     100	  11681028 ns/op	 6587281 B/op	    9344 allocs/op
PASS
ok  	github.com/nilsmagnus/grib/griblib/gribtest	5.456s
go tool pprof -top profile.out
File: gribtest.test
Type: cpu
Time: Feb 9, 2019 at 6:44pm (CET)
Duration: 5.41s, Total samples = 5.34s (98.77%)
Showing nodes accounting for 5.12s, 95.88% of 5.34s total
Dropped 69 nodes (cum <= 0.03s)
      flat  flat%   sum%        cum   cum%
     1.58s 29.59% 29.59%      1.78s 33.33%  github.com/nilsmagnus/grib/griblib.(*BitReader).readBit
     1.07s 20.04% 49.63%      2.85s 53.37%  github.com/nilsmagnus/grib/griblib.(*BitReader).readUint
     0.53s  9.93% 59.55%      0.53s  9.93%  runtime.memclrNoHeapPointers
     0.36s  6.74% 66.29%      0.85s 15.92%  encoding/binary.(*decoder).value
     0.15s  2.81% 69.10%      2.78s 52.06%  github.com/nilsmagnus/grib/griblib.(*BitReader).readIntsBlock
     0.15s  2.81% 71.91%      0.15s  2.81%  github.com/nilsmagnus/grib/griblib.(*Data3).applySpacialDifferencing
     0.14s  2.62% 74.53%      0.26s  4.87%  reflect.Value.Index
     0.13s  2.43% 76.97%      0.20s  3.75%  bytes.(*Buffer).ReadByte
     0.11s  2.06% 79.03%      0.28s  5.24%  github.com/nilsmagnus/grib/griblib.(*Data2).scaleValues
     0.11s  2.06% 81.09%      0.19s  3.56%  reflect.Value.SetUint
     0.08s  1.50% 82.58%      0.08s  1.50%  github.com/nilsmagnus/grib/griblib.Data0.scaleFunc.func1
     0.08s  1.50% 84.08%      0.08s  1.50%  reflect.(*rtype).Kind (inline)
     0.07s  1.31% 85.39%      0.07s  1.31%  bytes.(*Buffer).empty (inline)
     0.06s  1.12% 86.52%      0.68s 12.73%  runtime.mallocgc
     0.06s  1.12% 87.64%      0.06s  1.12%  runtime.memmove
     0.05s  0.94% 88.58%      0.05s  0.94%  reflect.flag.mustBeAssignable
     0.05s  0.94% 89.51%      0.05s  0.94%  runtime.nextFreeFast (inline)
     0.04s  0.75% 90.26%      0.04s  0.75%  encoding/binary.(*decoder).uint8 (inline)
   
 ```

* As you can see from the sample output, it is the readBit and readUint that takes most of the time. If anyone know how to optimize these functions further, please let me know :)
* Performance can be better if you only want to read certain kinds of messages. The filter that is implemented will filter only after reading all the messages. (help appreciated :) )

#### Memory


```
go tool pprof -top memprofile.out
File: gribtest.test
Type: alloc_space
Time: Feb 12, 2019 at 8:19pm (CET)
Showing nodes accounting for 3.39GB, 99.31% of 3.41GB total
Dropped 43 nodes (cum <= 0.02GB)
      flat  flat%   sum%        cum   cum%
    1.52GB 44.71% 44.71%     2.05GB 60.09%  github.com/nilsmagnus/grib/griblib.(*Data2).extractData
    0.76GB 22.19% 66.90%     0.76GB 22.19%  github.com/nilsmagnus/grib/griblib.(*Data2).scaleValues
    0.48GB 14.09% 80.99%     0.48GB 14.09%  github.com/nilsmagnus/grib/griblib.(*BitReader).readIntsBlock
    0.26GB  7.68% 88.67%     0.34GB 10.01%  github.com/nilsmagnus/grib/griblib.(*Data2).extractBitGroupParameters
    0.08GB  2.38% 91.05%     0.08GB  2.38%  github.com/nilsmagnus/grib/griblib.(*BitReader).readUintsBlock
    0.06GB  1.90% 92.95%     3.34GB 97.93%  github.com/nilsmagnus/grib/griblib.readMessage
    0.06GB  1.76% 94.70%     0.06GB  1.77%  encoding/binary.Read
    0.06GB  1.70% 96.40%     3.40GB 99.64%  github.com/nilsmagnus/grib/griblib.ReadMessage
    0.06GB  1.62% 98.02%     0.06GB  1.62%  github.com/nilsmagnus/grib/griblib.makeBitReader
    0.04GB  1.29% 99.31%     0.04GB  1.29%  github.com/nilsmagnus/grib/griblib.(*bitGroupParameter).zeroGroup (inline)
```

* extractData is the most memory-hungry function. It could probably be more efficient, but is now optimized 80%.  
* scaleValues could possibly be optimized, but is much better now than previous versions.


# Grib Documentation

Grib specification:

http://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_doc.shtml

Documentation from noaa.gov :

http://www.nco.ncep.noaa.gov/pmb/docs/on388/


Daily grib2 files from NOAA can be found at

http://www.ftp.ncep.noaa.gov/data/nccf/com/gfs/prod/
