package handlers

import (
	"aggregator/services/client"
	"fmt"
	"net/http"
	"strconv"
)

type InvoiceHandler struct {
	invoicer client.Invoicer
}

func NewInvoiceHandler(invoicer client.Invoicer) *InvoiceHandler {
	return &InvoiceHandler{
		invoicer: invoicer,
	}
}

func (h *InvoiceHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(r.URL.Query().Get("obu"))
	if err != nil {
		return fmt.Errorf("obu id is missing")
	}

	res, err := h.invoicer.GetInvoice(id)
	if err != nil {
		return fmt.Errorf("internal server error")
	}

	writeJson(w, http.StatusOK, res)

	return nil
}
