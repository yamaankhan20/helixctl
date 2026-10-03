package domain_test

import (
	"testing"

	"github.com/yamaankhan20/helixctl/internal/domain"
)

func TestResources_Fits(t *testing.T) {
	r := domain.Resources{CPUMillis: 1000, MemoryBytes: 1024}
	fits := domain.Resources{CPUMillis: 500, MemoryBytes: 512}
	tooBig := domain.Resources{CPUMillis: 1500, MemoryBytes: 512}

	if !r.Fits(fits) {
		t.Errorf("expected %v to fit %v", r, fits)
	}
	if r.Fits(tooBig) {
		t.Errorf("expected %v not to fit %v", r, tooBig)
	}
}
