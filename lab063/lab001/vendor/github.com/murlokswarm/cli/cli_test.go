package cli

import "testing"

func TestExec(t *testing.T) {
	if err := Exec("ls", "-la"); err != nil {
		t.Fatal(err)
	}
}
