package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pragmaticivan/tinyestate-api/city"
	log "github.com/sirupsen/logrus"
)

// CityHandler represent the httphandler for city
type CityHandler struct {
	cityUsecase city.Usecase
}

// NewCityHTTPHandler -
func NewCityHTTPHandler(r *mux.Router, cuc city.Usecase) {
	handler := &CityHandler{
		cityUsecase: cuc,
	}
	r.HandleFunc("/v1/states/{id}/cities", handler.Fetch).Methods("GET")
}

// Fetch -
func (s *CityHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idp, _ := strconv.Atoi(vars["id"])
	id := int64(idp)

	listAr, _ := s.cityUsecase.GetByStateID(context.Background(), id)
	sj, _ := json.Marshal(listAr)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "%s", sj)
	if err != nil {
		log.Warn(err)
	}
}
