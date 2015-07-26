package api

import (
	"testing"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func getTestAPI() (*API, error) {
	log.SetLevel(log.ErrorLevel)
	return NewAPI("", nil)
}

func TestNewAPI(t *testing.T) {
	a, err := getTestAPI()
	if err != nil {
		t.Fatal(err)
	}

	if a == nil {
		t.Fatalf("expected to get API; received nil")
	}
}

func TestSetupRoutes(t *testing.T) {
	a, err := getTestAPI()
	if err != nil {
		t.Fatal(err)
	}

	expectedRoutes := map[string]bool{
		"/":     true,
		"/sign": true,
	}

	r, err := a.setupRoutes()
	if err != nil {
		t.Fatal(err)
	}

	if err := r.Walk(func(rt *mux.Route, rtr *mux.Router, ancestors []*mux.Route) error {
		u, err := rt.URL()
		if err != nil {
			t.Fatal(err)
		}

		log.Debugf("setup route: path=%v", u)
		if _, ok := expectedRoutes[u.Path]; !ok {
			t.Fatalf("unexpected route: path=%s", u.Path)
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}
}
