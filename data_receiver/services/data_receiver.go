package services

import (
	"encoding/json"
	"fmt"
	"log"

	"data-receiver.com/types"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gorilla/websocket"
)

var kafkaTopic string = "obudata"

type DataReceiver struct {
	Conn *websocket.Conn
	Prod *kafka.Producer
}

func NewDataReceiver() (*DataReceiver, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		return nil, err
	}

	// Delivery report
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

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

		if err := dr.ProduceData(data); err != nil {
			fmt.Println("kafka produce error:", err)
		}
	}
}

func (dr *DataReceiver) ProduceData(data types.OBU) error {
	_, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return dr.Prod.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kafkaTopic, Partition: kafka.PartitionAny},
		Value:          []byte("test producing"),
	}, nil)
}
