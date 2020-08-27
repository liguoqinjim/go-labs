package main

import "testing"

var (
	cases = []struct {
		IP     string
		Result bool
	}{
		{IP: "1.2.3.4", Result: true},
		{IP: "266.1.2.3", Result: false},
		{IP: "100.100.100.100", Result: true},
		{IP: "-1", Result: false},
		{IP: "1.2.3;4", Result: false},
		{IP: "1.2.3.4.", Result: false},
		{IP: "-1.2.3.4.", Result: false},
	}
)

func TestIsIpv4(t *testing.T) {
	for _, c := range cases {
		is := IsIpv4(c.IP)
		if is != c.Result {
			t.Errorf("IsIpv4[%s] error:want[%t],got[%t]", c.IP, c.Result, is)
		}
	}
}

func TestIsIpv4Net(t *testing.T) {
	for _, c := range cases {
		is := IsIpv4Net(c.IP)
		if is != c.Result {
			t.Errorf("IsIpv4Net[%s] error:want[%t],got[%t]", c.IP, c.Result, is)
		}
	}
}

func TestIsIpv4Regex(t *testing.T) {
	for _, c := range cases {
		is := IsIpv4Regex(c.IP)
		if is != c.Result {
			t.Errorf("IsIpv4Regex[%s] error:want[%t],got[%t]", c.IP, c.Result, is)
		}
	}
}

//benchmark
func BenchmarkIsIpv4Net(b *testing.B) {
	b.Run("is-valid-ipv4-net-pkg", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = IsIpv4Net("10.41.132.6")
		}
	})
}

func BenchmarkIsIpv4(b *testing.B) {
	b.Run("is-valid-ipv4-custom-method", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = IsIpv4("10.41.132.6")
		}
	})
}

func BenchmarkIsIpv4Regex(b *testing.B) {
	b.Run("is-valid-ipv4-regex", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = IsIpv4Regex("10.41.132.6")
		}
	})
}
