package services

import (
	"fmt"
	"aggregator/store"
	"aggregator/types"
)

type InvoiceAggregator struct {
	store store.Storer
}

func NewInvoiceAggregator(store store.Storer) *InvoiceAggregator {
	return &InvoiceAggregator{
		store: store,
	}
}

func (i *InvoiceAggregator) AggregateDistance(distance types.Distance) error {
	fmt.Println("processing and storing distance in storage:", distance)
	return i.store.Insert(distance)
}
