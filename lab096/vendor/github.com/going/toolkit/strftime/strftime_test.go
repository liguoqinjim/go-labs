package strftime

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	t.Log(Format("%Y-%m-%d %H:%M:%S", time.Now()))
}

func TestParse(t *testing.T) {
	t.Log(Parse("%Y-%m-%d %H:%M:%S", "2013-04-15 02:10:06"))
	t.Log(Parse("%d/%b/%Y:%H:%M:%S %z", "14/Apr/2013:18:38:28 +0800"))
}
