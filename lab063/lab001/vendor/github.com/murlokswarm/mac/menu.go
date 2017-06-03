package mac

/*
#include "menu.h"
*/
import "C"
import (
	"os"
	"path/filepath"
	"strconv"
	"time"
	"unsafe"

	"github.com/murlokswarm/app"
	"github.com/murlokswarm/errors"
	"github.com/murlokswarm/log"
	"github.com/murlokswarm/markup"
	"github.com/satori/go.uuid"
)

type menu struct {
	id        uuid.UUID
	ptr       unsafe.Pointer
	component app.Componer
}

func newMenu(m app.Menu) *menu {
	id := uuid.NewV1()
	cmenu := C.Menu__{
		ID: C.CString(id.String()),
	}
	defer free(unsafe.Pointer(cmenu.ID))

	menu := &menu{
		id:  id,
		ptr: C.Menu_New(cmenu),
	}
	app.Elements().Add(menu)
	return menu
}

func (m *menu) ID() uuid.UUID {
	return m.id
}

func (m *menu) Mount(c app.Componer) {
	if m.component != nil {
		C.Menu_Clear(m.ptr)
		markup.Dismount(m.component)
	}

	m.component = c
	root, err := markup.Mount(c, m.ID())
	if err != nil {
		log.Panic(err)
	}
	if err := m.mount(root); err != nil {
		log.Panic(err)
	}

	rootID := C.CString(root.ID.String())
	defer free(unsafe.Pointer(rootID))

	C.Menu_Mount(m.ptr, rootID)
}

func (m *menu) mount(n *markup.Node) (err error) {
	switch n.Tag {
	case "menu":
		if err = m.mountContainer(n); err != nil {
			return
		}

	case "menuitem":
		if err = m.mountItem(n); err != nil {
			return
		}

	default:
		return errors.Newf("%v markup is not supported in a menu context. valid tags are menu and menuitem", n)
	}

	for _, child := range n.Children {
		if child.Type == markup.ComponentNode {
			child = markup.Root(child.Component)
		}
		if err = m.mount(child); err != nil {
			return
		}
		m.associate(n, child)
	}
	return
}

func (m *menu) mountContainer(n *markup.Node) error {
	if n.Parent != nil && n.Parent.Tag != "menu" {
		return errors.Newf("%v can only have another menu as parent: %v", n, n.Parent)
	}

	label, _ := n.Attributes["label"]
	container := C.MenuContainer__{
		ID:    C.CString(n.ID.String()),
		Label: C.CString(label),
	}
	defer free(unsafe.Pointer(container.ID))
	defer free(unsafe.Pointer(container.Label))

	C.Menu_MountContainer(m.ptr, container)
	return nil
}

func (m *menu) mountItem(n *markup.Node) (err error) {
	if n.Parent == nil || n.Parent.Tag != "menu" {
		return errors.Newf("%v should have a menu as parent: %v", n, n.Parent)
	}

	label, _ := n.Attributes["label"]
	icon, _ := n.Attributes["icon"]
	shortcut, _ := n.Attributes["shortcut"]
	selector, _ := n.Attributes["selector"]
	onclick, _ := n.Attributes["onclick"]
	disabled, _ := n.Attributes["disabled"]
	separator, _ := n.Attributes["separator"]

	isDisabled, _ := strconv.ParseBool(disabled)
	isSeparator, _ := strconv.ParseBool(separator)

	var iconPath string
	if len(icon) != 0 {
		iconPath = filepath.Join(app.Resources(), icon)
		if !app.FileIsSupportedIcon(iconPath) {
			err = errors.Newf("extension of %v is not supported", iconPath)
			return
		}
		if _, err = os.Stat(iconPath); err != nil {
			return
		}
	}

	item := C.MenuItem__{
		ID:        C.CString(n.ID.String()),
		Label:     C.CString(label),
		Icon:      C.CString(iconPath),
		Shortcut:  C.CString(shortcut),
		Selector:  C.CString(selector),
		OnClick:   C.CString(onclick),
		Disabled:  boolToBOOL(isDisabled),
		Separator: boolToBOOL(isSeparator),
	}
	defer free(unsafe.Pointer(item.ID))
	defer free(unsafe.Pointer(item.Label))
	defer free(unsafe.Pointer(item.Icon))
	defer free(unsafe.Pointer(item.Shortcut))
	defer free(unsafe.Pointer(item.Selector))
	defer free(unsafe.Pointer(item.OnClick))

	C.Menu_MountItem(m.ptr, item)
	return
}

func (m *menu) associate(parent *markup.Node, child *markup.Node) {
	parentID := C.CString(parent.ID.String())
	childID := C.CString(child.ID.String())
	defer free(unsafe.Pointer(parentID))
	defer free(unsafe.Pointer(childID))

	C.Menu_Associate(m.ptr, parentID, childID)
}

func (m *menu) Component() app.Componer {
	return m.component
}

func (m *menu) Render(s markup.Sync) {
	if err := m.mount(s.Node); err != nil {
		log.Error(err)
	}
}

//export onMenuItemClick
func onMenuItemClick(cid *C.char, cmethod *C.char) {
	id := uuid.FromStringOrNil(C.GoString(cid))
	method := C.GoString(cmethod)

	if len(method) == 0 {
		return
	}

	app.UIChan <- func() {
		markup.HandleEvent(id, method, "")
	}
}

//export onMenuCloseFinal
func onMenuCloseFinal(cid *C.char) {
	id := uuid.FromStringOrNil(C.GoString(cid))

	ctx, ok := app.Elements().Get(id)
	if !ok {
		return
	}
	menu := ctx.(*menu)

	go func() {
		time.Sleep(time.Millisecond * 42)

		app.UIChan <- func() {
			markup.Dismount(menu.component)
			app.Elements().Remove(menu)
		}
	}()
}
