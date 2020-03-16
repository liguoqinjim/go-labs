package lab008

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	fmt.Println(testing.Short())
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	cases := []struct {
		a, b, result int
	}{
		{1, 2, 3},
		{2, 3, 5},
	}

	for _, c := range cases {
		r := Add(c.a, c.b)
		if c.result != r {
			t.Errorf("Add expect %d,but got %d", c.result, r)
		}
	}
}
