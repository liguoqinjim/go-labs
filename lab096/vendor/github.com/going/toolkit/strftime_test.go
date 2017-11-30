package toolkit_test

import (
	"github.com/going/toolkit/strftime"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	t.Log(strftime.Format("%Y-%m-%d %H:%M:%S", time.Now()))
}

func TestParse(t *testing.T) {
	t.Log(strftime.Parse("%Y-%m-%d %H:%M:%S", "2013-04-15 02:10:06"))
	t.Log(strftime.Parse("%d/%b/%Y:%H:%M:%S %z", "14/Apr/2013:18:38:28 +0800"))
}
