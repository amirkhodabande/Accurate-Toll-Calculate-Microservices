package handlers

import (
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

func (h *InvoiceHandler) Handle(w http.ResponseWriter, r *http.Request) {
	invoiceProcessor := services.NewInvoiceProcessor(h.store)
	
	id, err := strconv.Atoi(r.URL.Query().Get("obu"))
	if err != nil {
		writeJson(w, http.StatusBadRequest, map[string]any{"success": false, "message": "obu id is missing"})
		return
	}

	invoice, err := invoiceProcessor.CalculateInvoice(id)
	if err != nil {
		writeJson(w, http.StatusInternalServerError, map[string]any{"success": false, "message": err.Error()})
		return
	}

	writeJson(w, http.StatusOK, map[string]any{"success": true, "data": invoice})
}
