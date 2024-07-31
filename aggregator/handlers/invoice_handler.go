package handlers

import (
	"aggregator/handlers/custom_errors"
	"aggregator/services"
	"aggregator/store"
	"net/http"
	"strconv"
)

type InvoiceHandler struct {
	store store.Storer
}

func NewInvoiceHandler(store store.Storer) *InvoiceHandler {
	return &InvoiceHandler{
		store: store,
	}
}

func (h *InvoiceHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	invoiceProcessor := services.NewInvoiceProcessor(h.store)

	id, err := strconv.Atoi(r.URL.Query().Get("obu"))
	if err != nil {
		return custom_errors.BadReq()
	}

	invoice, err := invoiceProcessor.CalculateInvoice(id)
	if err != nil {
		return custom_errors.Internal(err.Error())
	}

	return writeJson(w, http.StatusOK, map[string]any{"success": true, "data": invoice})
}
