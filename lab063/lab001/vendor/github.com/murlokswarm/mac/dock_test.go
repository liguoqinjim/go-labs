package mac

import "testing"

func TestNewDock(t *testing.T) {
	newDock()
}

func TestDockMount(t *testing.T) {
	driver.running = true
	defer func() { driver.running = false }()

	d := newDock()
	c := &MenuComponent{}

	d.Mount(c)
}

func TestDockSetIcon(t *testing.T) {
	driver.running = true
	defer func() { driver.running = false }()

	d := newDock()

	// Set.
	d.SetIcon("resources/logo.png")

	// Unset.
	d.SetIcon("")

	// Bad extension.
	d.SetIcon("resources/logo.bmp")

	// Nonexistent.
	d.SetIcon("resources/logosh.png")
}

func TestDockSetBadge(t *testing.T) {
	driver.running = true
	defer func() { driver.running = false }()

	d := newDock()
	d.SetBadge(42)
}
