package client

import (
	"aggregator/types"
)

type Aggregator interface {
	AggregateInvoice(distance types.Distance) (*types.AggregationClientResponse, error)
}
