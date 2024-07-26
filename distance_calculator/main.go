package main

import (
	"aggregator/services/client"
	"distance_calculator/services"
	"log"
)

func main() {
	calculatorService := services.NewCalculatorService()

	httpInvoiceAggregator := client.NewHTTPAggregator()
	KafkaConsumer, err := services.NewKafkaConsumer("obudata", calculatorService, httpInvoiceAggregator)
	if err != nil {
		log.Fatal(err)
	}
	KafkaConsumer.Start()
}
