# An implementation of Shamir's secret sharing
[![GPLv3](https://img.shields.io/badge/license-GPLv3-blue.svg)](https://www.gnu.org/copyleft/gpl.html)
[![Build Status](https://travis-ci.org/itslab-kyushu/sss.svg?branch=master)](https://travis-ci.org/itslab-kyushu/sss)
[![Release](https://img.shields.io/badge/release-0.1.0-brightgreen.svg)](https://github.com/itslab-kyushu/sss/releases/tag/v0.1.0)

## Installation
```sh
$ go get github.com/itslab-kyushu/sss
```

Compiled binaries are also available on
[Github](https://github.com/itslab-kyushu/sss/releases).

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
