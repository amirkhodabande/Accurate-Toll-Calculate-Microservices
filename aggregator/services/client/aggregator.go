package client

import (
	"aggregator/types"
	"net/http"
)

type Aggregator interface {
	AggregateInvoice(distance types.Distance) (*http.Response, error)
}
