# libracore – A Simple, Powerful Library for Libra Blockchain.

[![Build Status](https://travis-ci.org/keyfuse/libracore.png)](https://travis-ci.org/keyfuse/libracore) [![Go Report Card](https://goreportcard.com/badge/github.com/keyfuse/libracore)](https://goreportcard.com/report/github.com/keyfuse/libracore) [![codecov.io](https://codecov.io/gh/keyfuse/libracore/graphs/badge.svg)](https://codecov.io/gh/keyfuse/libracore/branch/master) [![GPL License](http://img.shields.io/badge/license-GPL-blue.svg?style=flat)](LICENSE) 

## libracore

*libracore* is a simple Go (golang) library for creating and manipulating Libra blockchain data structures like creating keys and addresses, creating and signing transactions.

## Overview

* Address creation
* Account query
* Transaction creation, signature and verification
* Two-Party Ed25519 Threshold Signature Scheme (TSS)

## Focus

* Simple and easy to use
* Full test coverage

## Tests

```
$ export GOPATH=`pwd`
$ go get -u github.com/keyfuse/libracore/xcore
$ cd src/github.com/keyfuse/libracore/
$ make test
```

## Examples

- [Addresses](examples/address.go)
- [Accounts](examples/account_query.go)
- [Mint Coins](examples/account_mint.go)
- [Transaction](examples/p2p_transaction.go)
- [Transfer Coins](examples/p2p_transfer.go)
- [Two-Party-Threshold Transaction](examples/two_party_ed25519_transaction.go)

Try Examples:
```
$ make runexamples
```

## Can I trust this code?
> Don't trust. Verify.

## License
libracore is released under the GPLv3 License.
