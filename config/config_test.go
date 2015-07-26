package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testConfig = `
{
    "roots": {
        "primary": {
            "key": "../test/certs/primary.key",
            "certificate": "../test/certs/primary.pem"
        },
        "secondary": {
            "key": "../test/certs/secondary.key",
            "certificate": "../test/certs/secondary.pem"
        }
    },
    "signing": {
        "default": {
            "expiry": "8760h"
        },
        "profiles": {
            "client": {
                    "usages": [
                            "signing",
                            "key encipherment",
                            "client auth"
                    ],
                    "expiry": "8760h"
            },
            "node": {
                    "usages": [
                            "signing",
                            "key encipherment",
                            "server auth",
                            "client auth"
                    ],
                    "expiry": "8760h"
            },
            "intermediate": {
                    "usages": [
                            "signing",
                            "key encipherment",
                            "cert sign",
                            "crl sign"
                    ],
                    "is_ca": true,
                    "expiry": "8760h"
            }
        }
    }
}
`
)

func getTestConfigPath() (string, error) {
	f, err := ioutil.TempFile("", "exile-test")
	if err != nil {
		return "", err
	}
	defer f.Close()

	f.Write([]byte(testConfig))

	return f.Name(), nil
}

func TestLoadConfig(t *testing.T) {
	configPath, err := getTestConfigPath()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(configPath)

	cfg, err := LoadConfig(configPath)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(cfg.Roots), fmt.Sprintf("expected 2 roots; received %d", len(cfg.Roots)))

	expectedRoots := []string{
		"primary",
		"secondary",
	}

	for _, r := range expectedRoots {
		if _, ok := cfg.Roots[r]; !ok {
			t.Fatalf("expected root %s in config", r)
		}
	}
}

func TestLoadSigners(t *testing.T) {
	configPath, err := getTestConfigPath()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(configPath)

	cfg, err := LoadConfig(configPath)
	if err != nil {
		t.Fatal(err)
	}

	// override config paths
	signers, err := LoadSigners(cfg)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(signers), fmt.Sprintf("expected 2 signers; received %d", len(signers)))
}

func TestLoadPrivateKey(t *testing.T) {
	configPath, err := getTestConfigPath()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(configPath)

	cfg, err := LoadConfig(configPath)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEqual(t, len(cfg.Roots), 0, "expected at least 1 root")
	c, ok := cfg.Roots["primary"]
	if !ok {
		t.Fatal("expected primary root")
	}

	s, err := LoadPrivateKey(c.Key)
	if err != nil {
		t.Fatal(err)
	}

	if s == nil {
		t.Fatal("expected signer; received nil")
	}
}
