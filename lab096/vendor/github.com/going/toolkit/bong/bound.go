/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          bound.go
 * Description:   Protorpc common functions
 */

package bong

import (
	"bufio"
	"encoding/binary"
	"io"
)

func read(r *bufio.Reader) (f []byte, err error) {
	b := make([]byte, 4)
	if _, err = io.ReadFull(r, b); err != nil {
		return
	}
	f = make([]byte, binary.BigEndian.Uint32(b))
	if _, err = io.ReadFull(r, f); err != nil {
		return
	}
	return
}

func write(w *bufio.Writer, f []byte) (err error) {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(len(f)))
	if _, err = w.Write(b); err != nil {
		return
	}
	if _, err = w.Write(f); err != nil {
		return
	}
	if err = w.Flush(); err != nil {
		return
	}
	return
}
