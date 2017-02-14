---
title: Command line usage
menu: main
date: 2017-02-09
lastmod: 2017-02-14
weight: 10
description: >-
  The secret sharing scheme this command provides is a simple
  k-out-of-n threshold secret sharing,
  which means totally *n* shares will be made from one secret file,
  and you need to at least *k* shares to reconstruct the secret.
---
The secret sharing scheme this command provides is a simple
[*k*-out-of-*n* threshold secret sharing](http://dx.doi.org/10.1007/978-3-642-20901-7_2),
which means totally *n* shares will be made from one secret file,
and you need to at least *k* shares to reconstruct the secret.

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

This command produces share files and the file name of i-th share has `.i.xz`
as the suffix.

This command also takes an optional flag `--chunk` to specify the byte size of
each chunk.
The given secret file is divided to chunks based on this size and distributed
in shares.

#### Reconstruct
```
$ sss local reconstruct <file>...
```

It produces a file based on the given share's file name by removing the above
suffix.
For example, if the names of share files are `sample.txt.1.xz`,
`sample.txt.2.xz`, ..., then the default file name of the reconstructed secret
will be `sample.txt`.

You can use `--output` flag to use another file name.


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
