---
title: An Implementation of Secret Sharing
type: homepage
menu:
  main:
    Name: Top
date: 2017-02-09
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
If you want to use this software as a library or you're familiar with Go,

```sh
$ go get github.com/itslab-kyushu/sss
```

If you're a [Homebrew](http://brew.sh/) user,

```sh
$ brew tap itslab-kyushu/sss
$ brew install sss
```

Otherwise, you can fine compiled binaries on
[Github](https://github.com/itslab-kyushu/sss/releases).
After downloading a binary to your environment, decompress and put it in a path
included in $PATH.


## License
This software is released under The GNU General Public License Version 3,
see [license](./licenses/) for more detail.
