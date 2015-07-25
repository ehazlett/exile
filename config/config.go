package config

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
	"github.com/cloudflare/cfssl/config"
	"github.com/cloudflare/cfssl/helpers"
	"github.com/cloudflare/cfssl/helpers/derhelpers"
	"github.com/cloudflare/cfssl/signer"
	"github.com/cloudflare/cfssl/signer/local"
)

type StoreConfig struct {
	Key         string `json:"key,omitempty"`
	Certificate string `json:"certificate,omitempty"`
}

type Config struct {
	Roots  map[string]*StoreConfig `json:"roots,omitempty"`
	Config *config.Signing         `json:"signing,omitempty"`
}

var (
	ErrUnsupportedPrivateKey = errors.New("unsupported private key type")
)

// LoadConfig returns a Config from the specified path
func LoadConfig(configPath string) (*Config, error) {
	var cfg Config
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// LoadSigners returns a map with a label as the key
// and the signer as the value
func LoadSigners(cfg *Config) (map[string]signer.Signer, error) {
	signers := make(map[string]signer.Signer)

	for l, c := range cfg.Roots {
		priv, err := LoadPrivateKey(c.Key)
		if err != nil {
			return nil, err
		}

		certData, err := ioutil.ReadFile(c.Certificate)
		if err != nil {
			return nil, err
		}
		cert, err := helpers.ParseCertificatePEM(certData)
		if err != nil {
			return nil, err
		}

		switch p := priv.(type) {
		case *rsa.PrivateKey, *ecdsa.PrivateKey:
			s, err := local.NewSigner(p, cert, signer.DefaultSigAlgo(priv), nil)
			if err != nil {
				return nil, err
			}
			s.SetPolicy(cfg.Config)
			signers[l] = s
		default:
			return nil, ErrUnsupportedPrivateKey
		}
	}

	return signers, nil
}

// LoadPrivateKey loads a crypto signer from the specified path
func LoadPrivateKey(keyPath string) (crypto.Signer, error) {
	// used from github.com/cloudflare/cfssl/cmd/multirootca/ca.go
	in, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}

	log.Debug("attempting to load PEM-encoded private key")
	priv, err := helpers.ParsePrivateKeyPEM(in)
	if err != nil {
		log.Debug("file is not a PEM-encoded private key")
		log.Debug("attempting to load DER-encoded private key")
		priv, err = derhelpers.ParsePrivateKeyDER(in)
		if err != nil {
			return nil, err
		}
	}

	log.Debug("loaded private key")
	return priv, nil
}
