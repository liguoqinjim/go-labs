package utest

import (
	"io"
	"testing"
)

func Test_All(t *testing.T) {
	Check(t, 1 > 2, 1, 2, 3)
	IsNil(t, io.ErrClosedPipe)
	Equal(t, 1, 2)
	Equal(t, 1.233, 3.333)
	Equal(t, '你', '好')
	Equal(t, "sadkfjsl", "sdfs*\r\n")
	Equal(t, []byte{1, 2, 3, 3}, []byte{3, 4, 5, 6})
	Equal(t, []int{1, 2, 3}, []int{1, 2, 3})
	Equal(t, []int{1, 2, 3}, []int{3, 4, 5})
	DeepEqual(t, []int{1, 2, 3}, []int{3, 4, 5})
	Equal(t, int32(100), 100)
	IsNilNow(t, 111)
	IsNilNow(t, 222)
}
