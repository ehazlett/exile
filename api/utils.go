package api

import (
	"github.com/cloudflare/cfssl/signer"
)

func isCSRValid(r *signer.SignRequest) bool {
	if r.Request == "" {
		return false
	}

	if r.Profile == "" {
		return false
	}

	if r.Label == "" {
		return false
	}

	return true
}
