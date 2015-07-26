package version

import (
	"testing"
)

func TestFullVersion(t *testing.T) {
	v := FullVersion()

	if v != "0.0.1 (HEAD)" {
		t.Fatalf("unexpected version received")
	}
}
