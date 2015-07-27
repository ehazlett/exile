package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiPostSignNoContent(t *testing.T) {
	a, err := getTestAPI()
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(a.sign))
	defer ts.Close()

	res, err := http.Post(ts.URL, "application/json", nil)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 400, res.StatusCode, "expected response code 400")
}

func TestApiPostSignBadCSR(t *testing.T) {
	a, err := getTestAPI()
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(a.sign))
	defer ts.Close()

	csr := []byte(`{"request": "12345"}`)
	buf := bytes.NewBuffer(csr)

	res, err := http.Post(ts.URL, "application/json", buf)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 400, res.StatusCode, "expected response code 400")
}
