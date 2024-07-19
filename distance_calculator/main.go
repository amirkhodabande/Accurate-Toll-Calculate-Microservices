package main

import (
	"distance_calculator/services"
	"distance_calculator/services/api"
	"log"
)

func main() {
	calculatorService := services.NewCalculatorService()
	aggregatorClient := api.NewAggregator()

	KafkaConsumer, err := services.NewKafkaConsumer("obudata", calculatorService, aggregatorClient)
	if err != nil {
		log.Fatal(err)
	}
	KafkaConsumer.Start()
}
