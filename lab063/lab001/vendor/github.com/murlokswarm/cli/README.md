# cli
[![Build Status](https://travis-ci.org/murlokswarm/cli.svg?branch=master)](https://travis-ci.org/murlokswarm/cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/murlokswarm/cli)](https://goreportcard.com/report/github.com/murlokswarm/cli)
[![GoDoc](https://godoc.org/github.com/murlokswarm/cli?status.svg)](https://godoc.org/github.com/murlokswarm/cli)

Package to launch command lines.

```go
func main() {
	if err := cli.Exec("ls", "-la"); err != nil {
		fmt.Println(err)
	}
}
```