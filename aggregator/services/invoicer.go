package services

import (
	"aggregator/store"
	"aggregator/types"
	"fmt"
)

type Invoicer interface {
	CalculateInvoice(obuID int) error
}

type InvoiceProcessor struct {
	store store.Storer
}

func NewInvoiceProcessor(store store.Storer) *InvoiceProcessor {
	return &InvoiceProcessor{
		store: store,
	}
}

func (i *InvoiceProcessor) CalculateInvoice(obuID int) (*types.Invoice, error) {
	distance, err := i.store.Get(obuID)
	if err != nil {
		return nil, fmt.Errorf("obu is invalid")
	}

	inv := &types.Invoice{
		OBUID:    obuID,
		Distance: distance,
		Amount:   0.15 * distance,
	}

	return inv, nil
}
