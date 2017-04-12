package lab003

import (
	"testing"
)

type AddArray struct {
	result int
	add_1  int
	add_2  int
}

func BenchmarkLoops(b *testing.B) {
	var test ForTest
	ptr := &test
	for i := 0; i < b.N; i++ {
		ptr.Loops()
	}
}

//测试并发效率
func BenchmarkLoopsParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var test ForTest
		ptr := &test
		for pb.Next() {
			ptr.Loops()
		}
	})
}

func TestAdd(t *testing.T) {
	var test_data = [3]AddArray{
		{2, 1, 1},
		{5, 2, 3},
		{4, 2, 2},
	}
	for _, v := range test_data {
		if v.result != Add(v.add_1, v.add_2) {
			t.Errorf("Add(%d,%d) != %d \n", v.add_1, v.add_2, v.result)
		}
	}
}
