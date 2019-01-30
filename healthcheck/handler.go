package healthcheck

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pragmaticivan/tinyestate-api/model"
)

// Handler healthcheck request
func Handler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, model.HealthCheck{Status: "Ok"})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}
