package Game

//https://github.com/Shopify/go-lua/blob/master/doc/presentations/golangmtl-march-2016/presentation.md
//^ actually useful docs

import "github.com/Shopify/go-lua"
import "github.com/Shopify/goluago"

var l *lua.State

func InitialzeLua(manager *GameManager) {
	l = lua.NewState()

	lua.OpenLibraries(l)
	goluago.Open(l)
	registerLuaFunctions(l, manager)

}

func registerLuaFunctions(state *lua.State, manager *GameManager) {

}
