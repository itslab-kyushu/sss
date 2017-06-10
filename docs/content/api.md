---
title: API Reference
menu: main
date: 2017-02-09
lastmod: 2017-03-08
weight: 30
---
[![GoDoc](https://godoc.org/github.com/itslab-kyushu/sss/sss?status.svg)](https://godoc.org/github.com/itslab-kyushu/sss/sss)

Package `github.com/itslab-kyushu/sss/sss` provides Distribute and Reconstruct
functions, which execute distribute and reconstruct procedures defined in the
threshold Secret Sharing scheme.
It also provides useful structures, Field and Polynomial.

This page explains a basic usage of those functions.
See [godoc](https://godoc.org/github.com/itslab-kyushu/sss/sss) for the detailed
information.

## Installation
```shell
$ go get -d github.com/itslab-kyushu/sss
```
## Example
### Compute shares from a secret
Distribute function takes secret, chunk size, the total number of shares,
and minimum number of shares to reconstruct the secret, in this order,
and returns a slice of shares.

The following example reads a file, creates shares and stores them in JSON
format.

```go
secret, err := ioutil.ReadFile("secret-file")
if err != nil {
  return err
}

shares, err := sss.Distribute(secret, chunksize, totalShares, threshold)
if err != nil {
  return err
}

for i, s := range shares {
  data, err := json.Marshal(s)
  if err != nil {
    return err
  }
  filename := fmt.Sprintf("%s.%d.json", "share-", i)
  if err = ioutil.WriteFile(filename, data, 0644); err != nil {
    return err
  }
}
```

### Reconstruct the secret
Reconstruct takes a slice of shares and returns the secret in a slice of bytes.

The following example reads a set of share files, reconstruct the secret, and
writes it to a file.


```go
// filenames is a slice of file names of shares.
shares := make([]sss.Share, len(filenames))
for i, f := range filenames {

  data, err := ioutil.ReadFile(f)
  if err != nil {
    return err
  }

  if err = json.Unmarshal(data, &shares[i]); err != nil {
    return err
  }

}

secret, err := sss.Reconstruct(shares)
if err != nil {
  return err
}
return ioutil.WriteFile("secret-file", secret, 0644)
```
