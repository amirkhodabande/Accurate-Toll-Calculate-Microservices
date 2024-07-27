package main

import (
	"aggregator/services/client"
	"distance_calculator/services"
	"log"
)

func main() {
	calculatorService := services.NewCalculatorService()

	// httpInvoiceAggregator := client.NewHTTPAggregator()
	grpcInvoiceAggregator := client.NewGRPCAggregator()

	KafkaConsumer, err := services.NewKafkaConsumer("obudata", calculatorService, grpcInvoiceAggregator)
	if err != nil {
		log.Fatal(err)
	}
	KafkaConsumer.Start()
}
