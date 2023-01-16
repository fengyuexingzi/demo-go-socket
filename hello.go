package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	var upgrader = websocket.Upgrader{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {
			for {
				mType, msg, err := conn.ReadMessage()
				if err != nil {
					println("close with error")
					println(string(err.Error()))
					conn.Close()
					return
				}
				println((string(msg)))
				switch mType {
				case 1:
					var sendmsg []byte = []byte("hi,I'm Go")
					conn.WriteMessage(mType, sendmsg)
				case 8:
					fmt.Println("client close")
				default:
					fmt.Println("type::" + string(rune(mType)))
				}

			}
		}(conn)
	})
	println("serve start")
	http.ListenAndServe(":3000", nil)
}
