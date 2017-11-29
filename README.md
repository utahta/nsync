# nsync

[![GoDoc Reference](https://godoc.org/github.com/utahta/nsync?status.svg)](http://godoc.org/github.com/utahta/nsync)
[![Build Status](https://travis-ci.org/utahta/nsync.svg?branch=master)](https://travis-ci.org/utahta/nsync)
[![codecov](https://codecov.io/gh/utahta/nsync/branch/master/graph/badge.svg)](https://codecov.io/gh/utahta/nsync)
[![Go Report Card](https://goreportcard.com/badge/github.com/utahta/nsync)](https://goreportcard.com/report/github.com/utahta/nsync)

A named mutual exclusion lock written in Go.

## Usage

```go
import "github.com/utahta/nsync"

func main() {
	var mux nsync.Mutex
	name := "lock name"
	mux.Lock(name)
	defer mux.Unlock(name)
}
```
