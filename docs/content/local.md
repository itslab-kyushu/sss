---
title: Local mode
menu: main
date: 2017-03-08
lastmod: 2017-03-08
weight: 10
description: >-
  Local mode of the client application provides two commands, distribute and
  reconstruct.
  Distribute command reads a given file and computes shares of
  the k-out-of-n threshold secret sharing scheme.
  It means totally n shares will be made from a secret file,
  and you must have at least k shares to reconstruct the secret.
  Reconstruct command does that phase, i.e. it reads at least *k* share files
  and reconstruct the secret file.
---
[![Release](https://img.shields.io/badge/release-0.3.2-brightgreen.svg)](https://github.com/itslab-kyushu/sss/releases/tag/v0.3.2)

## Summary
Local mode of the client application provides two commands, distribute and
reconstruct.
Distribute command reads a given file and computes shares of
the *k*-out-of-*n* threshold secret sharing scheme.
It means totally *n* shares will be made from a secret file,
and you must have at least *k* shares to reconstruct the secret.
Reconstruct command does that phase, i.e. it reads at least *k* share files and
reconstruct the secret file.

## Distribute command
```shell
$ sss local distribute <file> <number of shares> <threshold>
```

This command reads a secret file `<file>` and makes share files.
Each share file will be compressed by xz,
and the file name of *i*-th share has suffix `.i.xz`.

You need to specify the number of shares this command totally makes by
`<number of shares>`, and how many shares are required to reconstruct the secret
by `<threshold>`.

This command also takes an optional flag `--chunk` to specify the byte size of
each chunk.
The given secret file is divided to chunks based on this size and distributed
in shares.

## Reconstruct command
```shell
$ sss local reconstruct <file>...
```

This command reconstructs a secret from a list of share files.
It produces a file based on the given share's file name by removing the above
suffix.
For example, if the names of share files are `sample.txt.1.xz`,
`sample.txt.2.xz`, ..., then the default file name of the reconstructed secret
will be `sample.txt`.

You can use `--output` flag to use another file name.

## Tutorial
Suppose `secret.dat` is a secret file and distributing it using 3-out-of-10
threshold secret sharing scheme.

```shell
$ sss local distribute secret.dat 10 3
```

The above command creates a set of secrets, `secret.dat.1.xz`, `secret.dat.2.xz`,
..., `secret.dat.10.xz`.
We can store each share file into a different storage in order to prevent
information leakage, and now we can delete the secret file `secret.dat`.

To reconstruct the secret from shares, we must to collect at least 3 share
files. Suppose we have `secret.dat.1.xz`, `secret.dat.2.xz`, and
`secret.dat.5.xz`.

```shell
$ sss local reconstruct secret.dat.1.xz secret.dat.2.xz secret.dat.5.xz
```

The above command reconstructs the secret and stores it as `secret.dat`.


## Installation
If you're a [Homebrew](http://brew.sh/) user,
you can install the client application by

```sh
$ brew tap itslab-kyushu/sss
$ brew install sss
```

Compiled binaries for some platforms are available on
[Github](https://github.com/itslab-kyushu/sss/releases).
To use these binaries, after downloading a binary to your environment, decompress and put it in a directory included in your $PATH.

You can also compile the client application by yourself.
To compile it, you first download the code base:

```shell
$ git clone https://github.com/itslab-kyushu/sss $GOPATH/src/itslab-kyushu/sss
```

Then, build the client application `sss`:

```shell
$ cd $GOPATH/src/itslab-kyushu/sss/client
$ go get -d -t -v .
$ go build -o sss
```

To build the command, [Go](https://golang.org/) > 1.7.4 is required.
