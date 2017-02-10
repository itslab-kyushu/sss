---
title: Command line usage
menu: main
date: 2017-02-09
weight: 10
description: >-
  The command line tool which this software provides is called sss.
  It has two commands, distribute and reconstruct,
  to create shares from a secret file and reconstruct the secret files from
  shares, respectively.
---
The command line tool which this software provides is called `sss`.
It has two commands, `distribute` and `reconstruct`,
to create shares from a secret file and reconstruct the secret files from
shares, respectively.

The secret sharing scheme this command provides is a simple
[*k*-out-of-*n* threshold secret sharing](http://dx.doi.org/10.1007/978-3-642-20901-7_2),
which means totally *n* shares will be made from one secret file,
and you need to at least *k* shares to reconstruct the secret.


## Distribute
```
$ sss distribute <file> <number of shares> <threshold>
```

This command produces share files and the file name of i-th share has `.i.xz`
as the suffix.

This command also takes an optional flag `--chunk` to specify the byte size of
each chunk.
The given secret file is divided to chunks based on this size and distributed
in shares.


## Reconstruct
```
$ sss reconstruct <file>...
```

It produces a file based on the given share's file name by removing the above
suffix.
For example, if the names of share files are `sample.txt.1.xz`,
`sample.txt.2.xz`, ..., then the default file name of the reconstructed secret
will be `sample.txt`.

You can use `--output` flag to use another file name.
