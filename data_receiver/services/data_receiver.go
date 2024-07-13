package services

import (
	"fmt"
	"log"

	"data-receiver.com/types"
	"github.com/gorilla/websocket"
)

var kafkaTopic string = "obudata"

type DataReceiver struct {
	Conn *websocket.Conn
	Prod DataProducer
}

func NewDataReceiver() (*DataReceiver, error) {
	p, err := NewKafkaProducer()
	if err != nil {
		return nil, err
	}

	return &DataReceiver{
		Prod: p,
	}, nil
}

func (dr *DataReceiver) WsReceiveLoop() {
	fmt.Println("client connected")
	for {
		var data types.OBU
		if err := dr.Conn.ReadJSON(&data); err != nil {
			log.Println("read error:", err)

			continue
		}

		if err := dr.Prod.ProduceData(data); err != nil {
			fmt.Println("kafka produce error:", err)
		}
	}
}
