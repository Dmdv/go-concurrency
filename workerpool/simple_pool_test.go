package workerpool

import (
	"errors"
	"testing"
)

func TestWorkerPool_NewPool(t *testing.T) {
	if _, err := NewSimplePool(0, 0); !errors.Is(err, ErrNoWorkers) {
		t.Fatalf("expected error when creating pool with 0 workers, got: %v", err)
	}
	if _, err := NewSimplePool(-1, 0); !errors.Is(err, ErrNoWorkers) {
		t.Fatalf("expected error when creating pool with -1 workers, got: %v", err)
	}
	if _, err := NewSimplePool(1, -1); !errors.Is(err, ErrNegativeChannelSize) {
		t.Fatalf("expected error when creating pool with -1 channel size, got: %v", err)
	}

	p, err := NewSimplePool(5, 0)
	if err != nil {
		t.Fatalf("expected no error creating pool, got: %v", err)
	}
	if p == nil {
		t.Fatal("NewSimplePool returned nil Pool for valid input")
	}
}
