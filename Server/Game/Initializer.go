package Game

//https://github.com/Shopify/go-lua/blob/master/doc/presentations/golangmtl-march-2016/presentation.md
//^ actually useful docs
//http://www.lua.org/manual/5.2/manual.html#lua_call

import (
	"fmt"
	"github.com/Shopify/go-lua"
	"github.com/Shopify/goluago/util"
	"io/ioutil"
	"log"
)
import "github.com/Shopify/goluago"

var l *lua.State
var gMan *GameManager

func InitialzeLua(manager *GameManager) {
	l = lua.NewState()
	gMan = manager
	lua.OpenLibraries(l)
	goluago.Open(l)
	util.Open(l)
	registerLuaFunctions()
	if err := lua.DoFile(l, "Server/Game/LuaFiles/Definitions.lua"); err != nil {
		log.Print("LUA:", err)
	}

}

func registerLuaFunctions() {

}

func LoadConfig() EngineConfig {
	if err := lua.DoFile(l, "Server/Game/LuaFiles/Config.lua"); err != nil {
		log.Print("LUA:", err)
	}
	c := EngineConfig{}
	l.Global("PathToRooms")
	l.Global("Port")
	l.Global("TickRate")
	e := false
	if !l.IsString(-3) {
		log.Print("LUA: PathToRooms should be a string")
		e = true
	}
	if !l.IsNumber(-2) {
		log.Print("LUA: Port should be a number")
		e = true
	}
	if !l.IsNumber(-1) {
		log.Print("LUA: TickRate should be a number")
		e = true
	}
	if e {
		panic("Lua Config was not able to load correctly")
	}
	c.PathToRooms, _ = l.ToString(-3)
	c.Port, _ = l.ToInteger(-2)
	c.TickRate, _ = l.ToNumber(-1)
	l.Pop(3)
	return c
}

func LoadRooms(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		loadRoom(path, f.Name())
	}
}

func loadRoom(path string, fName string) {
	r := Room{}
	r.Players = make(map[string]*Player)
	fmt.Println("starting on room " + fName)
	if err := lua.DoFile(l, path+"/"+fName); err != nil {
		log.Print("LUA: room error:", err)
		return
	}
	l.Global("UID")
	l.Global("Name")
	l.Global("Description")
	r.UID, _ = l.ToString(-3)
	r.Name, _ = l.ToString(-2)
	r.Description, _ = l.ToString(-1)
	l.Pop(3)
	loadExits(&r)
	gMan.AddRoom(&r)
}

func loadExits(r *Room) {
	r.Exits = make(map[Direction]string)
	l.Global("Exits")
	l.Field(-1, "D")
	l.Field(-2, "U")
	l.Field(-3, "W")
	l.Field(-4, "S")
	l.Field(-5, "E")
	l.Field(-6, "N")
	for i := 0; i < 6; i++ {
		if l.IsString(-1) {
			r.Exits[Direction(i)], _ = l.ToString(-1)
		}
		l.Pop(1)
	}
}
