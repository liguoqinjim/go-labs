package sysutil

import (
	"testing"
	"time"
)

func TestBootTime(t *testing.T) {
	t.Logf("boot time: %s", BootTime().Format(time.RFC3339Nano))
}
