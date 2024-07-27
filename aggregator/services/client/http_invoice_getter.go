package client

import (
	"aggregator/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HTTPInvoiceGetter struct{}

func NewHTTPInvoiceGetter() *HTTPInvoiceGetter {
	return &HTTPInvoiceGetter{}
}

func (a *HTTPInvoiceGetter) GetInvoice(obuID int) (*types.InvoiceGetterClientResponse, error) {
	qParam := fmt.Sprintf("obu=%v", obuID)
	req, err := http.NewRequest(
		"GET", fmt.Sprintf("http://localhost:3000/aggregate?%v", qParam), nil,
	)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	encodedRes, _ := io.ReadAll(res.Body)
	var invoice types.Invoice
	if err := json.Unmarshal(encodedRes, &invoice); err != nil {
		fmt.Println("error in decoding data:", err.Error())
		return nil, err
	}

	return &types.InvoiceGetterClientResponse{
		Success: true,
		Data:    invoice,
	}, nil
}
