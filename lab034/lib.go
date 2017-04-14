package lab034

import "reflect"

func StringSliceReflectEqual(a, b []string) bool {
	return reflect.DeepEqual(a, b)
}

func StringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for n, v := range a {
		if v != b[n] {
			return false
		}
	}

	return true
}

func StringSliceEqualBCE(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for n, v := range a {
		if v != b[n] {
			return false
		}
	}

	return true
}
