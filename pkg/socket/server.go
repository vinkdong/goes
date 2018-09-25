package socket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/vinkdong/gox/log"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func StartSocketServer(address string) {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/ws", wsHandler)
	log.Infof("started server at %s", address)
	http.ListenAndServe(address, nil)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {

	var err error
	conn, err := upgrader.Upgrade(w, r, nil)

	defer conn.Close()
	if err != nil {
		log.Info(err)
	}

	var (
		msgType int
		data    []byte
	)

	for {
		msgType, data, err = conn.ReadMessage()
		log.Infof("data type is %d", msgType)
		if err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
		if msgType == -1 {
			break
		}
	}
ERR:
	conn.Close()
}
