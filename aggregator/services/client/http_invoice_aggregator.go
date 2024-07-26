package client

import (
	"aggregator/types"
	"bytes"
	"encoding/json"
	"net/http"
)

type HTTPInvoiceAggregator struct{}

func NewHTTPAggregator() *HTTPInvoiceAggregator {
	return &HTTPInvoiceAggregator{}
}

func (a *HTTPInvoiceAggregator) AggregateInvoice(distance types.Distance) (*http.Response, error) {
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
