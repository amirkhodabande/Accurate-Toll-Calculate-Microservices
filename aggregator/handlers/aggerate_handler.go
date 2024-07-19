package handlers

import (
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

func (h *AggregateHandler) Handle(w http.ResponseWriter, r *http.Request) {
	invoiceAggregatorService := services.NewInvoiceAggregator(h.store)

	var distance types.Distance
	if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "bad request"})
		return
	}

	if err := invoiceAggregatorService.AggregateDistance(distance); err != nil {
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJson(w, http.StatusOK, map[string]string{"success": "true"})
}

func writeJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
