# Vimeoserver

##### VimeoServer's main focus is to act as a byte-range proxy. The user may create http requests providing two url parameters:

s: source URL
range: Byte range

A typical curl request to the server would look like the following:

```
curl 'localhost:8000/?s="http://storage.googleapis.com/vimeo-test/work-at-vimeo.mp4"&range=600-1000'
```

## Features:

VimeoServer implements a LRU in-memory cache. The cache implements an interface defined in cache.go. You are free to implement your own cache into the server as long as you adhere to the interface.

The cache itself uses a min-heap to track epoch time of byte entries. Each eviction cycle pop's the object with the lowest epoch timestamp and evicts this object from cache.

The cache implements it's own form of binary search. This need arose as we are dealing with ranges. In order to determine if we search "left" or "right", we needed some further logic.

VimeoServer will also act as a simple proxy, however the source URL will require the Accept-Ranges header and the bytes value for this header declared. If these requirements are met VimeoServer will proxy the request with no need for byte ranges or partial responses.

Tests are provided which confirm the functionality of the server along with the cache. These can be ran by:

```
cd ./vimeoserver/server
go test
```

## ToDo:
* More tests for higher coverage
* Implement security mechanisms
* Implement logging


## Installation
This package should be "go gettable". The fastest way to begin playing with the code is by using docker. The following instructions should work as long as you have the docker daemon running

```
> docker run -it --name golang01 golang:1.6 /bin/bash
root@691aef12be64:/go# go get github.com/ldelossa/vimeoserver
root@691aef12be64:/go/src/github.com/ldelossa/vimeoserver# go run main.go & # You are now running server in the background on port 8000
root@691aef12be64:/go/src/github.com/ldelossa/vimeoserver# curl -I 'localhost:8000/?s="http://storage.googleapis.com/vimeo-test/work-at-vimeo.mp4"&range=600-1000'
HTTP/1.1 200 OK
Content-Type: video/mp4
Date: Sun, 11 Sep 2016 03:23:14 GMT
Content-Length: 401
```

If you do not have docker installed, the typical methodology for pulling down go packages and running them on your local workstation will work fine.
