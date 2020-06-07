// auth stuff https://blog.rapid7.com/2016/07/13/quick-security-wins-in-golang/
package main

import (
	"Mud-Engine/Server/Game"
	"flag"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{}

var t *template.Template
var GMan Game.GameManager
var Config Game.EngineConfig

func logic(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Upgrade:", err)
		return
	}
	defer c.Close()
	pName := login(c)
	p := Game.Player{
		Name:       pName,
		Connection: c,
		GMan:       &GMan,
	}
	err = GMan.AddPlayer(pName, &p)
	GMan.StartingRoom.Enter(nil, &p)
	if err != nil {
		log.Println("Login:", err)
		c.WriteMessage(websocket.TextMessage, []byte("Character name invalid. Closing connection"))
	} else {
		p.Logic()
	}
}

func login(conn *websocket.Conn) string {
	err := conn.WriteMessage(websocket.TextMessage, []byte("sys::Please enter your name"))
	if err != nil {
		log.Println("write:", err)
	}
	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Println("read:", err)
	}
	log.Printf("recv: %s", message)
	return string(message)
}

func main() {
	ctor()
	flag.Parse()
	log.SetFlags(0)
	fs := http.FileServer(http.Dir("Server/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	lp := "Server/templates/index.html"

	t = template.Must(template.ParseFiles(lp))
	if t == nil {
		print("templ didnt load")
		return
	}
	http.HandleFunc("/", serveTemplate)

	http.HandleFunc("/logic", logic)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func ctor() {
	GMan.Players = make(map[string]*Game.Player)
	GMan.Rooms = make(map[string]*Game.Room)
	Game.InitialzeLua(&GMan)
	//load config
	Config = Game.LoadConfig()
	//load rooms
	Game.LoadRooms(Config.PathToRooms)
	//start tick
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {

	if err := t.Execute(w, "ws://"+r.Host+"/logic"); err != nil {
		log.Println("Parse error: ", err)
	}
}
