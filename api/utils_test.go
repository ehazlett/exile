package api

import (
	"testing"

	"github.com/cloudflare/cfssl/signer"
)

var (
	goodCSR = &signer.SignRequest{
		Hosts:   nil,
		Request: "-----BEGIN CERTIFICATE REQUEST-----\n12345\n",
		Profile: "default",
		Label:   "primary",
	}

	badCSR = &signer.SignRequest{
		Request: "-----BEGIN CERTIFICATE REQUEST-----\n12345\n",
	}
)

func TestIsValidCSR(t *testing.T) {
	valid := isCSRValid(goodCSR)
	if !valid {
		t.Fatal("expected valid CSR")
	}
}

func TestIsValidCSRBadCSR(t *testing.T) {
	valid := isCSRValid(badCSR)
	if valid {
		t.Fatal("expected invalid CSR")
	}
}
