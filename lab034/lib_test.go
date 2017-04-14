package lab034

import (
	"testing"
)

func BenchmarkStringSliceReflectEqual(b *testing.B) {
	sa := []string{"q", "w", "e", "r", "t"}
	sb := []string{"q", "w", "a", "s", "z", "x"}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		StringSliceReflectEqual(sa, sb)
	}
}

func BenchmarkStringSliceEqual(b *testing.B) {
	sa := []string{"q", "w", "e", "r", "t"}
	sb := []string{"q", "w", "a", "s", "z", "x"}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		StringSliceEqual(sa, sb)
	}
}

func BenchmarkStringSliceEqualBCE(b *testing.B) {
	sa := []string{"q", "w", "e", "r", "t"}
	sb := []string{"q", "w", "a", "s", "z", "x"}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		StringSliceEqualBCE(sa, sb)
	}
}

//大容量的slice测试
func BenchmarkStringSliceReflectEqual2(b *testing.B) {
	sa := make([]string, 10000)
	for i := 0; i < len(sa); i++ {
		sa[i] = "hello"
	}
	sb := make([]string, 10000)
	for i := 0; i < len(sb); i++ {
		sb[i] = "hello"
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		StringSliceReflectEqual(sa, sb)
	}
}
func BenchmarkStringSliceEqual2(b *testing.B) {
	sa := make([]string, 10000)
	for i := 0; i < len(sa); i++ {
		sa[i] = "hello"
	}
	sb := make([]string, 10000)
	for i := 0; i < len(sb); i++ {
		sb[i] = "hello"
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		StringSliceEqual(sa, sb)
	}
}
func BenchmarkStringSliceEqualBCE2(b *testing.B) {
	sa := make([]string, 10000)
	for i := 0; i < len(sa); i++ {
		sa[i] = "hello"
	}
	sb := make([]string, 10000)
	for i := 0; i < len(sb); i++ {
		sb[i] = "hello"
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		StringSliceEqualBCE(sa, sb)
	}
}
