package main

import (
	"github.com/yuin/gopher-lua"
)

func square(L *lua.LState) int {
	i := L.ToInt(1)
	ln := lua.LNumber(i * i)
	L.Push(ln)
	return 1
}

func main() {
	L := lua.NewState()
	defer L.Close()

	L.SetGlobal("square", L.NewFunction(square))
	if err := L.DoFile("test.lua"); err != nil {
		panic(err)
	}
}
