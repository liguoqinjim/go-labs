package mac

/*
#include "menu.h"
*/
import "C"
import "github.com/murlokswarm/app"

type contextMenu struct {
	*menu
}

func newContextMenu(m app.ContextMenu) *contextMenu {
	cm := &contextMenu{
		menu: newMenu(app.Menu(m)),
	}
	return cm
}

func (m *contextMenu) Close() error {
	return nil
}

func (m *contextMenu) Mount(c app.Componer) {
	m.menu.Mount(c)
	C.Menu_Show(m.ptr)
}
