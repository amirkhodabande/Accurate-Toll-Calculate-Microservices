package client

import (
	"aggregator/types"
)

type Invoicer interface {
	GetInvoice(obuID int) (*types.InvoiceGetterClientResponse, error)
}
