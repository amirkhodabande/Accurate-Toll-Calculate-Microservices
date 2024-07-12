package handlers

import (
	"log"
	"net/http"

	"data-receiver.com/services"
	"github.com/gorilla/websocket"
)

type WsHandler struct {
	dr *services.DataReceiver
}

func NewWsHandler() *WsHandler {
	receiver, err := services.NewDataReceiver()
	if err != nil {
		log.Fatal(err)
	}

	return &WsHandler{
		dr: receiver,
	}
}

func (h *WsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1028,
		WriteBufferSize: 1028,
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	h.dr.Conn = conn

	go h.dr.WsReceiveLoop()
}
