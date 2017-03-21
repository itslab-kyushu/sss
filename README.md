# Shamir's Threshold Secret Sharing
[![GPLv3](https://img.shields.io/badge/license-GPLv3-blue.svg)](https://www.gnu.org/copyleft/gpl.html)
[![CircleCI](https://circleci.com/gh/itslab-kyushu/sss/tree/master.svg?style=svg)](https://circleci.com/gh/itslab-kyushu/sss/tree/master)
[![wercker status](https://app.wercker.com/status/16562999f1f803486bd8893c1dec21e6/s/master "wercker status")](https://app.wercker.com/project/byKey/16562999f1f803486bd8893c1dec21e6)
[![Release](https://img.shields.io/badge/release-0.3.1-brightgreen.svg)](https://github.com/itslab-kyushu/sss/releases/tag/v0.3.1)
[![Dockerhub](https://img.shields.io/badge/dockerhub-itslabq%2Fsss-blue.svg)](https://hub.docker.com/r/itslabq/sss/)
[![MicroBadger](https://images.microbadger.com/badges/image/itslabq/sss.svg)](https://microbadger.com/images/itslabq/sss)
[![GoDoc](https://godoc.org/github.com/itslab-kyushu/sss/sss?status.svg)](https://godoc.org/github.com/itslab-kyushu/sss/sss)

This software provides a [Go](https://golang.org/)
[library](https://godoc.org/github.com/itslab-kyushu/sss/sss) implementing
a Secret Sharing scheme, a command line tool which distributes and
reconstructs your secret files, and a client/server datastore service.

This software has been made for comparing performance of secret sharing based
key-value storages in the following article:

* [Hiroaki Anada](http://sun.ac.jp/prof/anada/),
  [Junpei Kawamoto](https://www.jkawamoto.info),
  Chenyutao Ke,
  [Kirill Morozov](http://www.is.c.titech.ac.jp/~morozov/), and
  [Kouichi Sakurai](http://itslab.inf.kyushu-u.ac.jp/~sakurai/),
  "[Cross-Group Secret Sharing Scheme for Secure Usage of Cloud Storage over Different Providers and Regions](http://www.anrdoezrs.net/links/8186671/type/dlg/https://link.springer.com/article/10.1007%2Fs11227-017-2009-7),"
  [The Journal of Supercomputing](http://www.anrdoezrs.net/links/8186671/type/dlg/https://link.springer.com/journal/11227), 2017.

Please consider to refer it, if you will publish any articles using this
software.

## Installation
### Go library
If you are only interested in our secret sharing library for Go,

```sh
$ go get -d github.com/itslab-kyushu/sss
```

### Client/Server application
If you are interested in our client/server application,
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

### Server application as a Docker image
We have a docker image [itslabq/sss](https://hub.docker.com/r/itslabq/sss/)
which includes a compiled binary of the server application.

[![Docker](https://itslab-kyushu.github.io/sss/img/small_h-trans.png)](https://www.docker.com/)

Containers created from this image exposes port 13009 and a volume `/data`
where all uploaded data will be stored.

To run this server image;

```
$ docker run -d --name sss-server -p 13009:13009 -v $(pwd)/data:/data itslabq/sss
```

The above command mounts `./data` to `/data` in the container so that all data
are store in `./data`.

## Client Usage
The client application provides two way to run the threshold Secret Sharing
scheme (SSS).
One of them is local mode, which stores shares into a local file system.
The other one is remote mode, which stores shares into servers provided the
server command.

### Local mode
The local mode provides two sub commands, distribute and reconstruct.
Distribute command reads a file and creates a set of shares,
on the other hand, reconstruct command reads a set of shares and reconstructs
the original file.

#### Distribute
```
$ sss local distribute <file> <number of shares> <threshold>
```

It produces share files and the file name of i-th share has `.i.xz` as the
suffix.

#### Reconstruct
```
$ sss local reconstruct <file>...
```

It produces a file based on the given share's file name by removing the above
suffix.

### Remote mode
Remote mode provides four sub command: get, put, delete, and list.
All commands take a YAML based server configuration file.
The format is as follows:

```yaml
servers:
  - address: 192.168.0.1
    port: 13009
  - address: 192.168.0.2
    port: 13009
  - address: 192.168.1.1
    port: 13009
```

The above example defines three servers.

The get command gathers shares from the servers defined the configuration file,
and put command distributes shares to the servers.

The default name of the configuration file is `sss.yml` but you can set another
name via `--config` flag.

#### Get
```
sss remote get --config sss.yml --output result.dat <file name>
```

Get command gathers shares associated with the given file name from the servers
defined in the configuration file, and then reconstructs and stores them as
the given file name via `--output` flag.

If `--config` flag is omitted, `sss.yml` is used, and if `--output` flag is
omitted, `<file name>` is used.

To find available file names, use list command.

The number of groups and the number of total servers must be greater then or
equal to the group threshold and the data threshold, which are given when those
shares were created.

#### Put
```
sss remote put --config sss.yml <file> <threshold>
```

Put command reads the given file and runs distribute procedure to create shares.
The threshold is a parameter of SSS.
The number of total shares are as same as defined in the server configuration
file.

If `--config` flag is omitted, `sss.yml` is used.

Put command also takes `--chunk` flag to set the byte size of each chunk.
The default value is 256.
The distribute procedure creates a finite filed Z/pZ, where p is a prime number
which has chunk size + 1 bit length.

### Delete
```
sss remote delete --config sss.yml <file name>
```

Delete command deletes all shares associated with the given file name from all
servers defined in the configuration file.

If `--config` flag is omitted, `sss.yml` is used.

### List
```
sss remote list --config sss.yml
```

List command shows all file names stored in the servers.
If `--config` flag is omitted, `sss.yml` is used.


## Server Usage
The server application runs a simple data store service using SSS.

It takes three flags,
* `--port`: the port number the server will listen,
* `--root`: the document root path to store uploaded shares,
* `--no-compress`: if set, all shares will be stored without compression.

If those flags are omitted, default values are used.
Thus, you can start a server by just run `sss-server`.


## Library Usage
See [godoc](https://godoc.org/github.com/itslab-kyushu/sss/sss).


## License
This software is released under The GNU General Public License Version 3,
see [COPYING](COPYING) for more detail.
