package lib_test

import (
	"strconv"
	"testing"
)

func benchmarkAppendFloat(b *testing.B, f float64, fmt byte, prec, bitSize int) {
	dst := make([]byte, 30)
	b.ResetTimer() // Overkill here, but for illustrative purposes.
	for i := 0; i < b.N; i++ {
		strconv.AppendFloat(dst[:0], f, fmt, prec, bitSize)
	}
}

func BenchmarkAppendFloatDecimal(b *testing.B) { benchmarkAppendFloat(b, 33909, 'g', -1, 64) }
func BenchmarkAppendFloat(b *testing.B)        { benchmarkAppendFloat(b, 339.7784, 'g', -1, 64) }
func BenchmarkAppendFloatExp(b *testing.B)     { benchmarkAppendFloat(b, -5.09e75, 'g', -1, 64) }
func BenchmarkAppendFloatNegExp(b *testing.B)  { benchmarkAppendFloat(b, -5.11e-95, 'g', -1, 64) }
func BenchmarkAppendFloatBig(b *testing.B) {
	benchmarkAppendFloat(b, 123456789123456789123456789, 'g', -1, 64)
}
