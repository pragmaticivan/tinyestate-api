package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pragmaticivan/tinyestate-api/state"
)

// HTTPStatecheckHandler  represent the httphandler for state
type HTTPStatecheckHandler struct {
	stateUsecase state.Usecase
}

// NewStatecheckHTTPHandler -
func NewStatecheckHTTPHandler(r *mux.Router, suc state.Usecase) {
	handler := &HTTPStatecheckHandler{
		stateUsecase: suc,
	}
	r.HandleFunc("/v1/states", handler.Fetch).Methods("GET")
}

// Fetch -
func (s *HTTPStatecheckHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	listAr, _ := s.stateUsecase.Fetch(context.Background())
	fmt.Printf("%#v\n", listAr)
	sj, _ := json.Marshal(listAr)
	fmt.Printf("%#v\n", sj)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", sj)
}
