package api

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/cloudflare/cfssl/signer"
	"github.com/ehazlett/exile/version"
	"github.com/gorilla/mux"
)

type API struct {
	addr    string
	signers map[string]signer.Signer
}

func NewAPI(listen string, signers map[string]signer.Signer) (*API, error) {
	a := &API{
		addr:    listen,
		signers: signers,
	}

	return a, nil
}

func (a *API) setupRoutes() (*mux.Router, error) {
	r := mux.NewRouter()
	r.HandleFunc("/", a.index).Methods("GET")
	r.HandleFunc("/sign", a.sign).Methods("POST")

	if err := r.Walk(func(rt *mux.Route, rtr *mux.Router, ancestors []*mux.Route) error {
		u, err := rt.URL()
		if err != nil {
			return err
		}

		log.Debugf("setup route: path=%v", u)
		return nil
	}); err != nil {
		return nil, err
	}

	return r, nil
}

func (a *API) Run() error {
	router, err := a.setupRoutes()
	if err != nil {
		return err
	}

	http.Handle("/", router)

	log.Infof("exile %s: addr=%s", version.FullVersion(), a.addr)
	return http.ListenAndServe(a.addr, nil)
}
