package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pragmaticivan/tinyestate-api/state"
	log "github.com/sirupsen/logrus"
)

// StateHandler  represent the httphandler for state
type StateHandler struct {
	stateUsecase state.Usecase
}

// NewStateHTTPHandler -
func NewStateHTTPHandler(r *mux.Router, suc state.Usecase) {
	handler := &StateHandler{
		stateUsecase: suc,
	}
	r.HandleFunc("/v1/states", handler.Fetch).Methods("GET")
}

// Fetch -
func (s *StateHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	listAr, _ := s.stateUsecase.Fetch(context.Background())
	sj, _ := json.Marshal(listAr)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "%s", sj)
	if err != nil {
		log.Warn(err)
	}
}
