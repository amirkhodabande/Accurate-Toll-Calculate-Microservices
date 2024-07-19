package services

import (
	"distance_calculator/services/api"
	"distance_calculator/types"
	"encoding/json"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaConsumer struct {
	consumer         *kafka.Consumer
	isRunning        bool
	calculator       Calculator
	aggregatorClient *api.Aggregator
}

func NewKafkaConsumer(topic string, calculator Calculator, aggregator *api.Aggregator) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return nil, err
	}

	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{
		consumer:         c,
		calculator:       calculator,
		aggregatorClient: aggregator,
	}, nil
}

func (c *KafkaConsumer) Start() {
	fmt.Println("kafka consumer start...")
	c.isRunning = true
	c.readMessageLoop()
}

func (c *KafkaConsumer) readMessageLoop() {
	for c.isRunning {
		msg, err := c.consumer.ReadMessage(time.Second)
		if err != nil {
			if !err.(kafka.Error).IsTimeout() {
				fmt.Println("kafka consume error:", err.Error())
			}

			continue
		}

		var data types.OBU
		if err := json.Unmarshal(msg.Value, &data); err != nil {
			fmt.Println("error in decoding data:", err.Error())
			continue
		}

		distance, err := c.calculator.CalculateDistance(data)
		if err != nil {
			fmt.Println("error in calculating distance:", err.Error())
			continue
		}

		req := types.Distance{
			Value: distance,
			Unix:  time.Now().Unix(),
			OBUID: data.OBUID,
		}
		if _, err = c.aggregatorClient.AggregateInvoice(req); err != nil {
			fmt.Println("error in aggregating distance:", err.Error())
			continue
		}

		fmt.Printf("distance: %.2f \n", distance)
	}
}
