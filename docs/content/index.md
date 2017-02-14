---
title: An Implementation of Secret Sharing
type: homepage
menu:
  main:
    Name: Top
date: 2017-02-09
lastmod: 2017-02-14
weight: 0
---
[![GPLv3](https://img.shields.io/badge/license-GPLv3-blue.svg)](https://www.gnu.org/copyleft/gpl.html)
[![Build Status](https://travis-ci.org/itslab-kyushu/sss.svg?branch=master)](https://travis-ci.org/itslab-kyushu/sss)
[![Release](https://img.shields.io/badge/release-0.2.0-brightgreen.svg)](https://github.com/itslab-kyushu/sss/releases/tag/v0.2.0)

This software provides both a [GO](https://golang.org/)
[library](https://godoc.org/github.com/itslab-kyushu/sss/sss) implementing
a Secret Sharing scheme and a command line tool which distributes and
reconstructs your secret files.


## Installation
If you are only interested in our secret sharing library for Go,

```sh
$ go get -d github.com/itslab-kyushu/sss
```

If you are interested in our server/client application,
compiled binaries of them are available on
[Github](https://github.com/itslab-kyushu/sss/releases).
After downloading a binary to your environment, decompress and put it in a path
included in $PATH.

If you're a [Homebrew](http://brew.sh/) user,
you can install the client application by

```sh
$ brew tap itslab-kyushu/sss
$ brew install sss
```

You can also compile by yourself.
First, you need to download the code base:

```
$ git clone https://github.com/itslab-kyushu/sss $GOPATH/src/itslab-kyushu/sss
```

Then, build client command `sss`:

```
$ cd $GOPATH/src/itslab-kyushu/sss/client
$ go get -d -t -v .
$ go build -o sss
```

and build server command `sss-server`:

```
$ cd $GOPATH/src/itslab-kyushu/sss/server
$ go get -d -t -v .
$ go build -o sss-server
```

To build both commands, [Go](https://golang.org/) > 1.7.4 is required.


## License
This software is released under The GNU General Public License Version 3,
see [license](./licenses/) for more detail.
