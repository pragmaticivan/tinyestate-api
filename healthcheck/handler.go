package healthcheck

import (
	"encoding/json"
	"net/http"
)

// HealthCheck - status
type HealthCheck struct {
	Status string `json:"status"`
}

// Handler healthcheck request
func Handler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, HealthCheck{Status: "Ok"})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
