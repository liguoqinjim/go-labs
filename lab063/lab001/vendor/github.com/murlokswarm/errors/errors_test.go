package errors

import "testing"

func TestNew(t *testing.T) {
	t.Log(New("test", "for", "New"))
}

func TestNewf(t *testing.T) {
	t.Log(Newf("test for %v", "Newf"))
}
