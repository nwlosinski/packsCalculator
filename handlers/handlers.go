package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nwlosinski/packsCalculator/calculator"
)

type Handler struct {
	service *calculator.Service
}

func NewHandler(service *calculator.Service) *Handler {
	return &Handler{service: service}
}

type CalculateRequest struct {
	Amount int `json:"amount"`
}

type PackSizesRequest struct {
	PackSizes []int `json:"packSizes"`
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/calculate", h.calculateHandler)

	mux.HandleFunc("/packsizes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.getPackSizesHandler(w, r)
		case http.MethodPost:
			h.setPackSizesHandler(w, r)
		default:
			http.Error(w, "method not allowed", 405)
		}
	})
}

func (h *Handler) calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", 405)
		return
	}

	var req CalculateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", 400)
		return
	}

	res, err := h.service.CalculatePacks(req.Amount)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) getPackSizesHandler(w http.ResponseWriter, r *http.Request) {
	sizes := h.service.GetPackSizes()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"packSizes": sizes,
	})
}

func (h *Handler) setPackSizesHandler(w http.ResponseWriter, r *http.Request) {
	var req PackSizesRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", 400)
		return
	}

	if len(req.PackSizes) == 0 {
		http.Error(w, "packSizes cannot be empty", 400)
		return
	}

	err := h.service.UpdatePackSizes(req.PackSizes)
	if err != nil {
		http.Error(w, "something went wrong", 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "pack sizes updated",
	})
}
