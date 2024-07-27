package client

import (
	"aggregator/types"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type HTTPInvoiceAggregator struct{}

func NewHTTPAggregator() *HTTPInvoiceAggregator {
	return &HTTPInvoiceAggregator{}
}

func (a *HTTPInvoiceAggregator) AggregateInvoice(distance types.Distance) (*types.AggregationClientResponse, error) {
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

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return &types.AggregationClientResponse{
		Success: true,
		Msg:     fmt.Sprintf("Aggregating: %.2f", distance.Value),
	}, nil
}
