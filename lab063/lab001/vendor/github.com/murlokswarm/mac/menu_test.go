package mac

import (
	"testing"
	"time"
	"unsafe"

	"github.com/murlokswarm/app"
)

type MenuComponent struct {
	Greeting                  string
	ErrorBadMarkup            bool
	ErrorInvalidTag           bool
	ErrorCompositionContainer bool
	ErrorCompositionItem      bool
	ErrorIconNonexistent      bool
	ErrorIconExt              bool
}

func (m *MenuComponent) Render() string {
	return `
<menu>
	<menuitem label="hello" separator="true" icon="logo.png"/>
	<menuitem label="boo" separator="true" {{if .ErrorIconNonexistent}}icon="logosh.png"{{end}}/>
	<menuitem label="bar" separator="true" {{if .ErrorIconExt}}icon="logo.bmp"{{end}}/>
	<menuitem label="{{if .Greeting}}{{.Greeting}}{{else}}world{{end}}" />
	<SubMenuComponent />

	{{if .ErrorBadMarkup}}
		<div>Pouette</span>
	{{end}}

	{{if .ErrorInvalidTag}}
		<div>Pouette</div>
	{{end}}

	{{if .ErrorCompositionContainer}}
		<menuitem>
			<menu></menu>
		</menuitem>
	{{end}}

	{{if .ErrorCompositionItem}}
		<menuitem>
			<menuitem></menuitem>
		</menuitem>
	{{end}}
</menu>
	`
}

type SubMenuComponent struct {
	Placeholder bool
}

func (m *SubMenuComponent) Render() string {
	return `
<menu>
	<menuitem label="foo" />
	<menuitem label="bar" />
</menu>
	`
}

func init() {
	app.RegisterComponent(&MenuComponent{})
	app.RegisterComponent(&SubMenuComponent{})
}

func TestMenuBar(t *testing.T) {
	driver.running = true
	defer func() { driver.running = false }()

	c := &SubMenuComponent{}
	m := newMenuBar()
	m.Mount(c)
}

func TestContextMenu(t *testing.T) {
	driver.running = true
	defer func() { driver.running = false }()

	c := &SubMenuComponent{}
	m := newContextMenu(app.ContextMenu{})
	m.Mount(c)
}

func TestMenu(t *testing.T) {
	newMenu(app.Menu{})
}

func TestMenuMount(t *testing.T) {
	m := newMenu(app.Menu{})

	c := &MenuComponent{}
	m.Mount(c)

	c2 := &MenuComponent{Greeting: "Maxoo"}
	m.Mount(c2)
}

func TestMenuMountBadMarkup(t *testing.T) {
	defer func() { recover() }()

	m := newMenu(app.Menu{})

	c := &MenuComponent{ErrorBadMarkup: true}
	m.Mount(c)
	t.Error("should panic")
}

func TestMenuMountInvalidTag(t *testing.T) {
	defer func() { recover() }()

	m := newMenu(app.Menu{})
	c := &MenuComponent{ErrorInvalidTag: true}
	m.Mount(c)
	t.Error("should panic")
}

func TestMenuMountErrorCompositionContainer(t *testing.T) {
	defer func() { recover() }()

	m := newMenu(app.Menu{})
	c := &MenuComponent{ErrorCompositionContainer: true}
	m.Mount(c)
	t.Error("should panic")
}

func TestMenuMountErrorCompositionItem(t *testing.T) {
	defer func() { recover() }()

	m := newMenu(app.Menu{})
	c := &MenuComponent{ErrorCompositionItem: true}
	m.Mount(c)
	t.Error("should panic")
}

func TestMenuMountErrorNonexistentIcon(t *testing.T) {
	defer func() { recover() }()

	m := newMenu(app.Menu{})
	c := &MenuComponent{ErrorIconNonexistent: true}
	m.Mount(c)
	t.Error("should panic")
}

func TestMenuMountErrorIconExt(t *testing.T) {
	defer func() { recover() }()

	m := newMenu(app.Menu{})
	c := &MenuComponent{ErrorIconExt: true}
	m.Mount(c)
	t.Error("should panic")
}

func TestMenuRender(t *testing.T) {
	m := newMenu(app.Menu{})
	c := &MenuComponent{}
	m.Mount(c)

	c.Greeting = "Maxence"
	app.Render(c)

	// Error.
	c.ErrorIconExt = true
	app.Render(c)
}

func TestOnMenuCloseFinal(t *testing.T) {
	m := newMenu(app.Menu{})

	cid := cString(m.ID().String())
	defer free(unsafe.Pointer(cid))

	onMenuCloseFinal(cid)
	time.Sleep(time.Millisecond * 50)
	onMenuCloseFinal(cid)
}
