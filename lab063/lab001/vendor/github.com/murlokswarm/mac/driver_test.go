package mac

import (
	"net/url"
	"sync"
	"testing"

	"github.com/murlokswarm/app"
)

func TestDriver(t *testing.T) {
	t.Log(driver.MenuBar())
	t.Log(driver.Dock())
	t.Log(driver.Resources())
	t.Log(driver.Storage())
	t.Log(driver.JavascriptBridge())
}

func TestDriverNewElement(t *testing.T) {
	driver.running = true
	defer func() { driver.running = false }()

	// Window => would bloc.
	// driver.Element(app.Window{})

	// Menu.
	driver.NewElement(app.ContextMenu{})
}

func TestDriverNewElementNotImplemented(t *testing.T) {
	defer func() { recover() }()

	driver.running = true
	defer func() { driver.running = false }()

	driver.NewElement("not implement")
}

func TestDriverNewElementPanic(t *testing.T) {
	defer func() { recover() }()

	driver.NewElement(app.Window{})
	t.Error("should panic")
}

func TestOnLaunch(t *testing.T) {
	app.OnLaunch = func() {
		t.Log("MacOS driver onLaunch")
	}
	onLaunch()
}

func TestFocused(t *testing.T) {
	app.OnFocus = func() {
		t.Log("MacOS driver onFocus")
	}
	onFocus()
}

func TestOnBlur(t *testing.T) {
	app.OnBlur = func() {
		t.Log("MacOS driver onBlur")
	}
	onBlur()
}

func TestOnReopen(t *testing.T) {
	app.OnReopen = func(v bool) {
		t.Log("MacOS driver onReopen:", v)

		if !v {
			t.Error("v should be true")
		}
	}
	onReopen(true)
}

func TestOnFilesOpen(t *testing.T) {
	app.OnFilesOpen = func(filenames []string) {
		t.Log("MacOS driver onFilesOpen:", filenames)
		if filenames[0] != "zune" {
			t.Error("filenames[0] should be zune")
		}
		if filenames[1] != "mune" {
			t.Error("filenames[1] should be mune")
		}
	}
	onFilesOpen(cString(`["zune", "mune"]`))
}

func TestOnURLOpen(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	app.OnURLOpen = func(URL url.URL) {
		t.Log("MacOS driver onURLOpen:", URL)
		wg.Done()
	}
	onURLOpen(cString("github-mac://openRepo/https://github.com/murlokswarm/app"))

	wg.Wait()

}

func TestOnTerminate(t *testing.T) {
	app.OnTerminate = func() bool {
		t.Log("MacOS driver onTerminate")
		return false
	}

	if ret := onTerminate(); ret {
		t.Error("ret should be false")
	}

	app.OnTerminate = nil

	if ret := onTerminate(); !ret {
		t.Error("ret should be true")
	}
}

func TestOnFinalize(t *testing.T) {
	app.OnFinalize = func() {
		t.Log("MacOS driver onFinalize")
	}
	onFinalize()
}
