package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func scrapData(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	fmt.Println("conexion aceptada: " + r.RemoteAddr)

	_, mensaje, err := ws.ReadMessage()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mensaje: " + string(mensaje))

	info := make(map[string]dataStruct)
	scraperData(info)

	ws.WriteJSON(info)

	fmt.Println("terminando conexion: " + r.RemoteAddr)
}
