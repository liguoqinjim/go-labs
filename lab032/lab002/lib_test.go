package lab002

import (
	"testing"
)

func BenchmarkLoops(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Division(6, 7)
	}
}
func BenchmarkDivision(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Division(6, 7)
	}
}
func TestDivision1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("除法函数测试没通过")
	} else {
		t.Log("第一个测试通过了")
	}
}
