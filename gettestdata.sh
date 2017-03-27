#!/bin/bash

mkdir testdata
date="2017032200"
for hour in "000" "003" "006" "009" "012"
do
    wget -O testdata/gfs.t00z.pgrb2.2p50.f$hour http://www.ftp.ncep.noaa.gov/data/nccf/com/gfs/prod/gfs.$date/gfs.t00z.pgrb2.2p50.f$hour
done