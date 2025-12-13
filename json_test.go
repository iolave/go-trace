package trace_test

import (
	"fmt"
	"testing"

	"github.com/iolave/go-trace"
)

func TestTrace_JSON(t *testing.T) {
	t.Run("should return an empty json object if the trace is nil", func(t *testing.T) {
		trace := make(trace.Trace)
		trace = nil

		expected := "{}"

		got, err := trace.MarshalJSON()
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if string(got) != expected {
			t.Fatalf("expected %s, got %s", expected, got)
		}
	})

	t.Run("should return a json object with the trace data", func(t *testing.T) {
		k := "key"
		v := "value"

		trace := trace.Trace{
			k: v,
		}

		expected := fmt.Sprintf(`{"%s":"%s"}`, k, v)

		got, err := trace.MarshalJSON()
		if err == nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if string(got) != expected {
			t.Fatalf("expected %s, got %s", expected, got)
		}
	})
}
