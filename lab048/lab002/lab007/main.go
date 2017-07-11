package main

import (
	"fmt"
	"github.com/yuin/gopher-lua"
)

type Hero struct {
	HeroId   int
	HeroLv   int
	HeroStar int
}

const luaHeroTypeName = "hero"

func registerHeroType(L *lua.LState) {
	mt := L.NewTypeMetatable(luaHeroTypeName)
	L.SetGlobal("hero", mt)
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), heroMethods))
}

func checkHero(L *lua.LState) *Hero {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*Hero); ok {
		return v
	}
	L.ArgError(1, "hero expected")
	return nil
}

var heroMethods = map[string]lua.LGFunction{
	"heroId":   heroGetSetHeroId,
	"heroLv":   heroGetSetHeroLv,
	"heroStar": heroGetSetHeroStar,
}

func heroGetSetHeroId(L *lua.LState) int {
	h := checkHero(L)
	if L.GetTop() == 2 {
		h.HeroId = L.CheckInt(2)
	}
	L.Push(lua.LNumber(h.HeroId))
	return 1
}

func heroGetSetHeroLv(L *lua.LState) int {
	h := checkHero(L)
	if L.GetTop() == 2 {
		h.HeroLv = L.CheckInt(2)
	}
	L.Push(lua.LNumber(h.HeroLv))
	return 1
}

func heroGetSetHeroStar(L *lua.LState) int {
	h := checkHero(L)
	if L.GetTop() == 2 {
		h.HeroStar = L.CheckInt(2)
	}
	L.Push(lua.LNumber(h.HeroStar))
	return 1
}

func main() {
	L := lua.NewState()
	defer L.Close()

	registerHeroType(L)

	if err := L.DoFile("hero.lua"); err != nil {
		panic(err)
	}

	for i := 1; i < 100; i++ {
		//hero
		h1 := &Hero{HeroId: i, HeroLv: 100 + i, HeroStar: 200 + i}
		fmt.Println(h1)

		ud := L.NewUserData()
		ud.Value = h1
		L.SetMetatable(ud, L.GetTypeMetatable(luaHeroTypeName))

		if err := L.CallByParam(lua.P{
			Fn:      L.GetGlobal("heroLevel"),
			NRet:    1,
			Protect: true,
		}, ud); err == nil {
			fmt.Println("执行成功")
		}

		if ret, ok := L.Get(-1).(lua.LNumber); ok {
			fmt.Println("ret:", ret)
		} else {
			fmt.Println(ok)
		}

		fmt.Println("修改后的hero,", h1)

		L.Pop(1)
	}
}
