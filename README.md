# An implementation of Shamir's secret sharing
[![GPLv3](https://img.shields.io/badge/license-GPLv3-blue.svg)](https://www.gnu.org/copyleft/gpl.html)
[![Build Status](https://travis-ci.org/itslab-kyushu/sss.svg?branch=master)](https://travis-ci.org/itslab-kyushu/sss)
[![wercker status](https://app.wercker.com/status/16562999f1f803486bd8893c1dec21e6/s/master "wercker status")](https://app.wercker.com/project/byKey/16562999f1f803486bd8893c1dec21e6)
[![Release](https://img.shields.io/badge/release-0.2.0-brightgreen.svg)](https://github.com/itslab-kyushu/sss/releases/tag/v0.2.0)

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


## Usage
### Distribute
```
$ sss distribute <file> <number of shares> <threshold>
```

It produces share files and the file name of i-th share has `.i.json` as the
suffix.

### Reconstruct
```
$ sss reconstruct <file>...
```

It produces a file based on the given share's file name by removing the above
suffix.


## License
This software is released under The GNU General Public License Version 3,
see [COPYING](COPYING) for more detail.
