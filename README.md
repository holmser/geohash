# Location Search

Simple script to learn location based queries with redis.

- Insert location database into redis
  - http://download.geonames.org/export/dump/
  ```sh
  wget http://download.geonames.org/export/dump/US.zip
  ```
- Begin to query them
- Play with goroutines and channels

# How Geohashing works

Get geohash in resolution you need

> **Haversine** formula is a standard way to calculate distance between 2 points on the surface of a sphere.  Earth is not a perfect sphere, so this method may introduce errors of up to .5%.  This is usually an acceptable error rate for social proximity searches.

![GeoHash + Haversine](img/GeoHashing.png?raw=true)

High level process is:
- Select point
- retrieve all points within 8 boxes
- apply haversine formula to results and drop from radius.


- Redis uses Haversine formula, error rate may be up to 0.5%
- Add via name, lat, lon
Sample data:
```
4085315	Ragland Cemetery	Ragland Cemetery		34.67398	-86.82695	S	CMTY	US		AL	083			0	180	178	America/Chicago	2006
```

### Real life query patterns

- Get places near me with activities on specific date/range
- Get places near me with activities 

# Links
- [Best GeoHashing explanation I've found](https://gis.stackexchange.com/questions/18330/using-geohash-for-proximity-searches)
