package utils

import "testing"

var cases = []struct {
	a, b   int
	result int
}{
	{1, 2, 3},
	{1, 3, 4},
}

func TestAdd(t *testing.T) {
	for _, c := range cases {
		r := Add(c.a, c.b)
		if r != c.result {
			t.Errorf("Add(%d,%d)= %d , want %d", c.a, c.b, r, c.result)
		}
	}
}
