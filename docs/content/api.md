---
title: API Reference
menu: main
date: 2017-02-09
weight: 20
---
This software provides package `github.com/itslab-kyushu/sss/sss`.
This package implements Distribute and Reconstruct functions, which execute
distribute and reconstruct procedures defined in a Secret Sharing scheme.
It also provides useful structures, Field and Polynomial.

This page explains a base usage of those functions.
See [godoc](https://godoc.org/github.com/itslab-kyushu/sss/sss) for the detailed
information.

## Distribute
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

## Reconstruct
Reconstruct takes a slice of shares and returns the secret in a slice of bytes.

The following example reads a set of share files, reconstruct the secret, and
writes it to a file.


```go
// shares is a slice of file names of shares.
shares := make([]sss.Share, len(shares))
for i, f := range shares {

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
