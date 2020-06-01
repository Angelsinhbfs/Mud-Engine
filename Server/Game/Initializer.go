package Game

//https://github.com/Shopify/go-lua/blob/master/doc/presentations/golangmtl-march-2016/presentation.md
//^ actually useful docs
//http://www.lua.org/manual/5.2/manual.html#lua_call

import (
	"github.com/Shopify/go-lua"
	"io/ioutil"
	"log"
	"strings"
)
import "github.com/Shopify/goluago"

var l *lua.State
var gMan *GameManager

func InitialzeLua(manager *GameManager) {
	l = lua.NewState()
	gMan = manager
	lua.OpenLibraries(l)
	goluago.Open(l)
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
	if err := lua.DoFile(l, path+"/"+fName); err != nil {
		log.Print("LUA: room error:", err)
		return
	}
	f := strings.Split(fName, ".")
	if len(f) > 0 {
		l.Global(f[0])
		l.Table(-1)
		r.UID = getStringField("UID")
		r.Name = getStringField("Name")
		r.Description = getStringField("Description")
	}

	gMan.AddRoom(r)
}

func getStringField(key string) string {
	l.PushString(key)
	l.Table(-2)
	if !l.IsString(-1) {
		log.Print("LUA: get field error. " + key + " is not a string")
	}
	s, _ := l.ToString(-1)
	l.Pop(1)
	return s
}

func getNumberField(key string) float64 {
	l.PushString(key)
	l.Table(-2)
	if !l.IsNumber(-1) {
		log.Print("LUA: get field error. " + key + " is not a number")
	}
	s, _ := l.ToNumber(-1)
	l.Pop(1)
	return s
}
