---
title: Threshold Secret Sharing Scheme
type: homepage
menu:
  main:
    Name: Top
date: 2017-02-09
lastmod: 2017-03-20
weight: 0
description: >-
  This software provides both a Go library and command line tools implementing
  a threshold Secret Sharing scheme.
---
[![GPLv3](https://img.shields.io/badge/license-GPLv3-blue.svg)](https://www.gnu.org/copyleft/gpl.html)
[![CircleCI](https://circleci.com/gh/itslab-kyushu/sss/tree/master.svg?style=svg)](https://circleci.com/gh/itslab-kyushu/sss/tree/master)
[![wercker status](https://app.wercker.com/status/16562999f1f803486bd8893c1dec21e6/s/master "wercker status")](https://app.wercker.com/project/byKey/16562999f1f803486bd8893c1dec21e6)
[![Release](https://img.shields.io/badge/release-0.3.1-brightgreen.svg)](https://github.com/itslab-kyushu/sss/releases/tag/v0.3.1)

## Summary
This software provides both a [Go](https://golang.org/)
[library](https://godoc.org/github.com/itslab-kyushu/sss/sss) and
command line tools implementing the threshold Secret Sharing scheme:

* Adi Shamir, "[How to share a secret](http://dl.acm.org/ft_gateway.cfm?id=359176),"
  Communications of the ACM, 22(11):pp.612–613, 1979.

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


### Contents
* To use the threshold secret sharing from another go application,
  see the [API Reference](api) page.  
* To compute shares and reconstruct a secret in a computer,
  see the [local mode usage](local) page.
* To use a secret sharing based data storage service,
  see the [client usage](remote) and [server usage](server) pages.

## License
This software is released under The GNU General Public License Version 3,
see [license](./licenses/) for more detail.
