package web

import (
	"net/http"

	"github.com/pragmaticivan/tinyestate-api/canonical"
	"github.com/pragmaticivan/tinyestate-api/city"
	"github.com/pragmaticivan/tinyestate-api/state"

	"github.com/gorilla/mux"
	canonicalHttp "github.com/pragmaticivan/tinyestate-api/canonical/delivery/http"
	cityHttp "github.com/pragmaticivan/tinyestate-api/city/delivery/http"
	healthcheckHttp "github.com/pragmaticivan/tinyestate-api/healthcheck/delivery/http"
	stateHttp "github.com/pragmaticivan/tinyestate-api/state/delivery/http"
	"github.com/urfave/negroni"
)

// NewWebAdapter -
func NewWebAdapter(suc state.Usecase, cuc city.Usecase, ccu canonical.Usecase) http.Handler {
	r := mux.NewRouter()

	healthcheckHttp.NewHealthcheckHTTPHandler(r)
	stateHttp.NewStateHTTPHandler(r, suc)
	cityHttp.NewCityHTTPHandler(r, cuc)
	canonicalHttp.NewCanonicalHTTPHandler(r, ccu)

	n := negroni.Classic()
	n.UseHandler(r)
	return n
}
