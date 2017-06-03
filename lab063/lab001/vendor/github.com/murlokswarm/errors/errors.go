// Package errors is a package to create and format errors.
package errors

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
)

type err struct {
	buffer bytes.Buffer
}

func (e *err) Error() string {
	return e.buffer.String()
}

// New formats using the default formats for its operands and returns the
// string as a value that satisfies error.
// Spaces are always added between operands.
// Last calls are printed to the end.
func New(a ...interface{}) error {
	err := &err{}

	for i, v := range a {
		if i != 0 {
			err.buffer.WriteRune(' ')
		}

		fmt.Fprint(&err.buffer, v)
	}

	printCalls(&err.buffer)
	return err
}

// Newf formats according to a format specifier and returns the string as a
// value that satisfies error.
// Last calls are printed to the end.
func Newf(format string, a ...interface{}) error {
	err := &err{}
	fmt.Fprintf(&err.buffer, format, a...)
	printCalls(&err.buffer)
	return err
}

func printCalls(w io.Writer) {
	fmt.Fprintln(w, "\ncalls:")

	for i := 0; i < 10; i++ {
		_, file, line, ok := runtime.Caller(2 + i)
		if !ok {
			break
		}

		if wd, err := os.Getwd(); err == nil {
			file = strings.TrimPrefix(file, wd)
		}

		file = strings.TrimPrefix(file, os.Getenv("GOPATH"))
		fmt.Fprintf(w, "- line %v\t%v\n", line, file)
	}

	fmt.Fprintln(w)
}
