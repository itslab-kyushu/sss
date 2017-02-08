# An implementation of Shamir's secret sharing
[![GPLv3](https://img.shields.io/badge/license-GPLv3-blue.svg)](https://www.gnu.org/copyleft/gpl.html)


## Installation
```
$ git clone https://github.com/itslab-kyushu/sss
$ cd sss
$ go build
```

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
