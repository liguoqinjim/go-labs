log4go
======

Simplified Logging API for go

Installation
============

    go get github.com/hugozhu/log4go

Usage
=====
```
import (
    "github.com/hugozhu/log4go"
)

var log = log4go.New(os.Stdout)

func main() {
    log.Debug("Hello, World!")
}
```    
