package main

import "github.com/yuin/gopher-lua"

type Hero struct {
	HeroId   int
	HeroLv   int
	HeroStar int
}

const luaHeroTypeName = "hero"

func registerHeroType(L *lua.LState) {
	mt := L.NewTypeMetatable(luaHeroTypeName)
	L.SetGlobal("hero", mt)

	L.SetField(mt, "new", L.NewFunction(newHero))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), heroMethods))
}

func newHero(L *lua.LState) int {
	hero := &Hero{HeroId: L.CheckInt(1), HeroLv: L.CheckInt(2), HeroStar: L.CheckInt(3)}
	ud := L.NewUserData()
	ud.Value = hero
	L.SetMetatable(ud, L.GetTypeMetatable(luaHeroTypeName))
	L.Push(ud)
	return 1
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
	"heroId": heroGetSetHeroId,
	"heroLv": heroGetSetHeroLv,
}

func heroGetSetHeroId(L *lua.LState) int {
	h := checkHero(L)
	if L.GetTop() == 2 { //set
		h.HeroId = L.CheckInt(2)
		return 0
	}
	L.Push(lua.LNumber(h.HeroId))
	return 1
}

func heroGetSetHeroLv(L *lua.LState) int {
	h := checkHero(L)
	if L.GetTop() == 2 {
		h.HeroLv = L.CheckInt(3)
		return 0
	}
	L.Push(lua.LNumber(h.HeroLv))
	return 1
}

func main() {
	L := lua.NewState()
	defer L.Close()

	registerHeroType(L)

	if err := L.DoString(`
        h = hero.new(55,66,77)
        print(h:heroId())
        print(h:heroLv())
        h:heroId(23)
        print(h:heroId())
    `); err != nil {
		panic(err)
	}
}
