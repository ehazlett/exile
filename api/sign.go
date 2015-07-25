package api

import (
	"encoding/json"
	"net/http"

	"github.com/cloudflare/cfssl/signer"
)

func (a *API) sign(w http.ResponseWriter, r *http.Request) {
	var signRequest signer.SignRequest
	if err := json.NewDecoder(r.Body).Decode(&signRequest); err != nil {
		http.Error(w, "invalid signing request", http.StatusBadRequest)
		return
	}

	signer, ok := a.signers[signRequest.Label]
	if !ok {
		http.Error(w, "unable to find signer with specified label", http.StatusBadRequest)
		return
	}

	cert, err := signer.Sign(signRequest)
	if err != nil {
		http.Error(w, "error signing request", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(string(cert)))
}
