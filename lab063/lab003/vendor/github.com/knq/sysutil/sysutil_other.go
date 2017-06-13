// +build !linux,!windows,!darwin,!freebsd,!netbsd,!openbsd

package sysutil

import "time"

func init() {
	btime = time.Now()
}
