package trace_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/iolave/go-trace"
)

func TestTrace_SetInContext(t *testing.T) {
	t.Run("should set the trace in the context", func(t *testing.T) {
		ctx := context.Background()

		tr := trace.Trace{}
		ctx = tr.SetInContext(ctx)

		got := ctx.Value(trace.CTX_KEY)
		if got == nil {
			t.Fatalf("expected a trace, got none")
		}

		if reflect.TypeOf(got) != reflect.TypeOf(tr) {
			t.Fatalf("expected %s, got %s", reflect.TypeOf(tr), reflect.TypeOf(got))
		}
	})
}

func TestGetFromContext(t *testing.T) {
	t.Run("should return a new trace from the context", func(t *testing.T) {
		ctx := context.Background()

		got := trace.GetFromContext(ctx)
		if got == nil {
			t.Fatalf("expected a trace, got none")
		}
	})

	t.Run("should return trace when conext has a pointer to a trace", func(t *testing.T) {
		tr := &trace.Trace{}
		ctx := context.WithValue(context.Background(), trace.CTX_KEY, tr)

		got := trace.GetFromContext(ctx)
		if got == nil {
			t.Fatalf("expected a trace, got none")
		}
	})

	t.Run("should return the trace struct when the context has a trace struct", func(t *testing.T) {
		tr := trace.Trace{}
		ctx := context.WithValue(context.Background(), trace.CTX_KEY, tr)

		got := trace.GetFromContext(ctx)
		if got == nil {
			t.Fatalf("expected a trace, got none")
		}
	})

	t.Run("should return a new trace when the value of the context is not a trace", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), trace.CTX_KEY, "not a trace")

		got := trace.GetFromContext(ctx)
		if got == nil {
			t.Fatalf("expected a trace, got none")
		}
	})
}
