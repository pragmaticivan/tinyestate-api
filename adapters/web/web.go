package web

import (
	"net/http"

	"github.com/pragmaticivan/tinyestate-api/canonical"
	"github.com/pragmaticivan/tinyestate-api/state"

	"github.com/gorilla/mux"
	canonicalHttp "github.com/pragmaticivan/tinyestate-api/canonical/delivery/http_transport"
	healthcheckHttp "github.com/pragmaticivan/tinyestate-api/healthcheck/delivery/http_transport"
	stateHttp "github.com/pragmaticivan/tinyestate-api/state/delivery/http_transport"
	"github.com/urfave/negroni"
)

// NewWebAdapter -
func NewWebAdapter(suc state.Usecase, ccu canonical.Usecase) http.Handler {
	r := mux.NewRouter()

	healthcheckHttp.NewHealthcheckHTTPHandler(r)
	stateHttp.NewStateHTTPHandler(r, suc)
	canonicalHttp.NewCanonicalHTTPHandler(r, ccu)

	n := negroni.Classic()
	n.UseHandler(r)
	return n
}
