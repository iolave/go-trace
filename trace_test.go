package trace_test

import (
	"testing"

	"github.com/iolave/go-trace"
)

func TestTrace_Set(t *testing.T) {
	t.Run("should set a key-value pair in the trace", func(t *testing.T) {
		trace := trace.Trace{}

		k := "key"
		v := "value"

		err := trace.Set(k, v)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		got := trace[k]
		if got != v {
			t.Fatalf("expected %s, got %s", v, got)
		}
	})

	t.Run("should return an error if the key is not in snake case", func(t *testing.T) {
		trace := trace.Trace{}

		k := "Key"
		v := "value"

		err := trace.Set(k, v)
		if err == nil {
			t.Fatalf("expected an error, got none")
		}
	})
}

func TestTrace_Get(t *testing.T) {
	t.Run("should get a key-value pair from the trace", func(t *testing.T) {
		k := "key"
		v := "value"

		trace := trace.Trace{
			k: v,
		}

		got := trace.Get(k)
		if got != v {
			t.Fatalf("expected %s, got %s", v, got)
		}
	})
}
