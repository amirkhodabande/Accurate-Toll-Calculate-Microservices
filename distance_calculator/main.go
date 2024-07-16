package main

import (
	"distance_calculator/services"
	"log"
)

func main() {
	calculatorService := services.NewCalculatorService()
	
	KafkaConsumer, err := services.NewKafkaConsumer("obudata", calculatorService)
	if err != nil {
		log.Fatal(err)
	}
	KafkaConsumer.Start()	
}
