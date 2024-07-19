package api

import (
	"bytes"
	"distance_calculator/types"
	"encoding/json"
	"net/http"
)

type Aggregator struct{}

func NewAggregator() *Aggregator {
	return &Aggregator{}
}

func (a *Aggregator) AggregateInvoice(distance types.Distance) (*http.Response, error) {
	payload, err := json.Marshal(distance)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST", "http://localhost:3000/aggregate", bytes.NewReader(payload),
	)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
