epoch
---------------------

A tiniest CLI tool to convert timestamps in seconds or milliseconds or datetime strings to timestamps and datetime strings.

Having to deal with timestamps a lot at work I occasionally need to quickly translate timestamps to dates and vice versa. I thought it will be handy to have it as a command line utility instead of having to go to a relevant website (e.g. epochconverter.com), especially when network connection is not available.

The datetime format parsing powered by the brilliant [dateparse](https://github.com/araddon/dateparse) library.

Usage
-----
```
Usage:
	epoch <timestamp or date>

Parameters:
	timestamp - A Unix timestamp in second or milliseconds
	date - A date in most common formats, e.g. 2019-04-16 12:05:00

Examples:
	epoch 1555416300
	epoch 1555416300000
	epoch 2019-04-16 12:05:00
	epoch 04/16/2019 12:05:00
	epoch 2019-04-16 15:05:00 +0300 UTC
```

Example Output
-----
```
$ epoch 2019-04-16 07:05:00 EST-05

Milliseconds:  1555416300000
         GMT:  2019-04-16 12:05:00 +0000 UTC
       Local:  2019-04-16 07:05:00 -0500 EST
```
