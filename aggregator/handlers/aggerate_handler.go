package handlers

import (
	"aggregator/handlers/custom_errors"
	"aggregator/services"
	"aggregator/store"
	"aggregator/types"
	"encoding/json"
	"net/http"
)

type AggregateHandler struct {
	store store.Storer
}

func NewAggregateHandler(store store.Storer) *AggregateHandler {
	return &AggregateHandler{
		store: store,
	}
}

func (h *AggregateHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	invoiceAggregatorService := services.NewInvoiceAggregator(h.store)

	var distance types.Distance
	if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
		return custom_errors.BadReq()
	}

	if err := invoiceAggregatorService.AggregateDistance(distance); err != nil {
		return custom_errors.Internal("")
	}

	return writeJson(w, http.StatusOK, map[string]string{"success": "true"})
}
