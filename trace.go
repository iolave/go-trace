// The trace package provides a Trace struct for
// storing and retrieving trace data.
//
// It also provides useful functions for setting
// and retrieving the trace to and from the
// context.
//
// There's also helpers for setting and retrieving
// the trace from http requests.
package trace

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/iolave/go-errors"
)

const (
	ERR_NAME = "trace_error"
)

// Trace is a map used to store trace
// data within an execution context.
//
//	i.e. An http request or an async event.
//
// Built-in properties:
//
//	A unique identifier for the trace.
type Trace map[string]string

// Get returns the value of the given key.
func (t Trace) Get(k string) string {
	return t[k]
}

// Set sets the value v fot the given key k. Keys must be
// in snake case, otherwise it will return an error.
//
// Errors are of type [github.com/iolave/go-errors.GenericError]
// and have the name [trace.ERR_NAME]. Errors can be casted
// to this type.
//
//	err := trace.Set("mykey", "this-will-error")
//	if err != nil {
//		e := err.(*errors.GenericError)
//	}
func (t Trace) Set(k string, v string) error {
	if got := strcase.ToSnake(k); got != k {
		return errors.NewWithName(
			ERR_NAME,
			fmt.Sprintf("trace key must be in snake case, (got %s, expected %s)", k, got),
		)

	}

	t[k] = v
	return nil
}
