package web

import (
	"net/http"

	"github.com/pragmaticivan/tinyestate-api/state"

	"github.com/gorilla/mux"
	healthcheckHttp "github.com/pragmaticivan/tinyestate-api/healthcheck/delivery/http"
	stateHttp "github.com/pragmaticivan/tinyestate-api/state/delivery/http"
	"github.com/urfave/negroni"
)

// NewWebAdapter -
func NewWebAdapter(suc state.Usecase) http.Handler {
	r := mux.NewRouter()

	healthcheckHttp.NewHealthcheckHTTPHandler(r)
	stateHttp.NewStatecheckHTTPHandler(r, suc)

	n := negroni.Classic()
	n.UseHandler(r)
	return n
}
