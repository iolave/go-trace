package trace

import (
	"context"
	"reflect"
)

// key used to store and retrieve
// the trace from a context.
const CTX_KEY = "trace"

// GetFromContext returns the trace from a context using
// the key [trace.CTX_KEY].
//
// If the trace value is nil or is not of type [trace.Trace],
// it will return a new empty trace. Returned trace will always
// be non-nil.
func GetFromContext(ctx context.Context) Trace {
	trace := ctx.Value(CTX_KEY)

	if trace == nil {
		return Trace{}
	}

	if reflect.TypeOf(trace).Kind() == reflect.Pointer {
		if trace, ok := trace.(*Trace); ok {
			return *trace
		}
	}

	if trace, ok := trace.(Trace); ok {
		return trace
	}

	return Trace{}
}

// SetInContext sets the trace in a context using
// the key [trace.CTX_KEY] and returns the a new context.
func (t Trace) SetInContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, CTX_KEY, t)
}
