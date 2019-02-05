package http_transport

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pragmaticivan/tinyestate-api/domain"
)

// HealthcheckHandler represent the httphandler for healthcheck
type HealthcheckHandler struct {
}

// NewHealthcheckHTTPHandler -
func NewHealthcheckHTTPHandler(r *mux.Router) {
	handler := &HealthcheckHandler{}
	r.HandleFunc("/_health", handler.Healthcheck).Methods("GET")
	r.HandleFunc("/", handler.Healthcheck).Methods("GET")

}

// Healthcheck -
func (a *HealthcheckHandler) Healthcheck(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(domain.HealthCheck{Status: "Ok"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_, err := w.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}
