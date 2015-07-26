package config

import (
	"encoding/json"
	"testing"
)

var (
	baseConfig = `
{
    "roots": {
        "primary": {
            "key": "/path/to/key",
            "certificate": "/path/to/cert",
            "config": "/path/to/config"
        }
    }
}
    `
)

func TestBaseConfig(t *testing.T) {
	var cfg Config
	if err := json.Unmarshal([]byte(baseConfig), &cfg); err != nil {
		t.Fatal(err)
	}

	if len(cfg.Roots) != 1 {
		t.Fatalf("expected 1 root; received %d", len(cfg.Roots))
	}

	c, ok := cfg.Roots["primary"]
	if c == nil || !ok {
		t.Fatalf("expected root %q in config", "primary")
	}

	keyPath := "/path/to/key"
	if c.Key != keyPath {
		t.Fatalf("expected key path %q; received", c.Key)
	}

}
