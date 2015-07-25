package api

import (
	"net/http"

	"github.com/ehazlett/exile/version"
)

func (a *API) index(w http.ResponseWriter, r *http.Request) {
	info := "exile " + version.FullVersion()
	w.Write([]byte(info))
}
