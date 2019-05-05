package handler

import (
	"encoding/json"
	"net/http"

	"github.com/morikuni/failure"
)

func (h *Handler) Weather(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var req struct {
		Address string `json:"address"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err := failure.Wrap(err, failure.WithCode(badRequest))
		handleError(w, err)
		return
	}

	ctx := r.Context()
	we, err := h.Reporter.Now(ctx, req.Address)
	if err != nil {
		handleError(w, failure.Wrap(err))
		return
	}

	resp := struct {
		Temperature       float64 `json:"temperature"`
		PrecipProbability float64 `json:"precipProbability"`
		Summary           string  `json:"summary"`
	}{
		Temperature:       float64(we.Temperature),
		PrecipProbability: we.PrecipProbability,
		Summary:           we.Summary,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		handleError(w, err)
		return
	}
}
