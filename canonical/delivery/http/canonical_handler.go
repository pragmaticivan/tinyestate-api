package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pragmaticivan/tinyestate-api/canonical"
	log "github.com/sirupsen/logrus"
)

// CanonicalHandler represent the httphandler for canonical
type CanonicalHandler struct {
	canonicalUsecase canonical.Usecase
}

// NewCanonicalHTTPHandler -
func NewCanonicalHTTPHandler(r *mux.Router, cuc canonical.Usecase) {
	handler := &CanonicalHandler{
		canonicalUsecase: cuc,
	}
	r.HandleFunc("/v1/canonicals", handler.Fetch).Methods("GET")
	r.HandleFunc("/v1/canonicals/{id}", handler.FetchByID).Methods("GET")
	r.HandleFunc("/v1/canonical/{canonical}", handler.FetchByCanonical).Methods("GET")
}

// Fetch -
func (s *CanonicalHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	list, _ := s.canonicalUsecase.Fetch(context.Background())
	sjson, _ := json.Marshal(list)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "%s", sjson)
	if err != nil {
		log.Warn(err)
	}
}

// FetchByID -
func (s *CanonicalHandler) FetchByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idp, _ := strconv.Atoi(vars["id"])
	id := int64(idp)

	list, _ := s.canonicalUsecase.FetchByID(context.Background(), id)
	sjson, _ := json.Marshal(list)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "%s", sjson)
	if err != nil {
		log.Warn(err)
	}
}

// FetchByCanonical -
func (s *CanonicalHandler) FetchByCanonical(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	list, _ := s.canonicalUsecase.FetchByCanonical(context.Background(), vars["canonical"])
	sjson, _ := json.Marshal(list)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "%s", sjson)
	if err != nil {
		log.Warn(err)
	}
}
