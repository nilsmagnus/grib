# Template5_x.grib

This document explains how fixtures files were generated

## Getting wgrib2

1. Download ftp://ftp.cpc.ncep.noaa.gov/wd51we/wgrib2/wgrib2.tgz and untar it
1. Install dependencies : on linux environment, install `gfortran-8`, `zlib1g-dev`, `build-essential`
1. exports variables fortran ``export FC=`which gfortran-8` `` and gcc ``export CC=`which gcc` ``
1. `make`

## Protocol

First, a file was downloaded from http://nomads.ncep.noaa.gov/cgi-bin/filter_gfs_1p00.pl\?bottomlat\=-90\&dir\=%2Fgfs.2019010612\&file\=gfs.t12z.pgrb2.1p00.f006\&leftlon\=0\&lev_10_m_above_ground\=on\&rightlon\=360\&toplat\=90\&var_UGRD\=on\&var_VGRD\=on

then, using wgrib2, the file was converted with the following commands

```bash
wgrib2 gfs.t12z.pgrb2.1p00.f006 -set_grib_type simple  -grib_out template5_0.grib2
wgrib2 gfs.t12z.pgrb2.1p00.f006 -set_grib_type complex2  -grib_out template5_2.grib2
wgrib2 gfs.t12z.pgrb2.1p00.f006 -set_grib_type complex3  -grib_out template5_3.grib2
```

And finally, csv files were generated with the following command

```bash
wgrib2 gfs.t12z.pgrb2.1p00.f006 -csv out.csv
```

Last column was extracted for `UGRD` and `VGRD`