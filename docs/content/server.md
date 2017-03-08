---
title: Server application
menu: main
date: 2017-03-08
lastmod: 2017-03-08
weight: 20
description: >-
  Server application provides a data storage service based on the threshold
  secret sharing scheme.
---
[![Release](https://img.shields.io/badge/release-0.3.1-brightgreen.svg)](https://github.com/itslab-kyushu/sss/releases/tag/v0.3.1)
[![Dockerhub](https://img.shields.io/badge/dockerhub-itslabq%2Fsss-blue.svg)](https://hub.docker.com/r/itslabq/sss/)
[![MicroBadger](https://images.microbadger.com/badges/image/itslabq/sss.svg)](https://microbadger.com/images/itslabq/sss)

## Summary
Server application provides a data storage service based on the threshold secret
sharing scheme.
The server application is based on [gRPC](http://www.grpc.io/), and you can
find the service definition in [here](https://github.com/itslab-kyushu/sss/blob/master/kvs/kvs.proto).

[![gRPC](../img/grpc.png)](http://www.grpc.io/)

## Usage
The server application takes three flags,

* `--port`: the port number the server will listen,
* `--root`: the document root path to store uploaded shares,
* `--no-compress`: if set, all shares will be stored without compression.

If those flags are omitted, default values are used.
Thus, you can start a server by just run `sss-server`.

## Docker support
A docker image to run the server, `itslabq/sss`, is also available.

[![docker](../img/small_h-trans.png)](https://www.docker.com/)

Containers based on this image expose port 13009 and store uploaded shares into
`/data`. The following command runs a container with mapping local port 13009
and `$(pwd)/data` to the container's port 13009 and `/data`:

```shell
$ docker run -it -p 13009:13009 -v $(pwd)/data:/data itslabq/sss
```

## Installation
Compiled binaries of the server application `sss-server` are available on
[Github](https://github.com/itslab-kyushu/sss/releases).
To use these binaries, after downloading a binary to your environment, decompress and put it in a directory included in your $PATH.

You can also compile by yourself.
First, you need to download the code base:

```shell
$ git clone https://github.com/itslab-kyushu/sss $GOPATH/src/itslab-kyushu/sss
```

Then, build server command `sss-server`:

```shell
$ cd $GOPATH/src/itslab-kyushu/sss/server
$ go get -d -t -v .
$ go build -o sss-server
```

To build it, [Go](https://golang.org/) > 1.7.4 is required.
