package trace

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/iolave/go-errors"
)

// HTTP_HEADER_PREFIX is the prefix that trace headers
// must have.
const HTTP_HEADER_PREFIX = "X-Trace-"

// SetHTTPHeaders sets trace data into the given [http.Header] h.
// It does by converting the trace keys underscores to dashes
// and appending at the beginning of each header key the
// [trace.HTTP_HEADER_PREFIX] value.
//
// If h is nil, it will do nothing.
func (t Trace) SetHTTPHeaders(h http.Header) {
	if h == nil {
		return
	}

	for k, v := range t {
		k = strings.ReplaceAll(k, "_", "-")
		h.Set(fmt.Sprintf("%s%s", HTTP_HEADER_PREFIX, k), v)
	}
}

// GetFromHTTPRequest builds the trace struct from an
// [http.Request] req. It does by finding headers
// with the [trace.HTTP_HEADER_PREFIX] then removeing
// the prefix and converting the key to snake case.
//
// It will return an error if req is nil or if
// req.Header is nil. trace will always be non-nil.
//
// Errors are of type [github.com/iolave/go-errors.GenericError]
// and have the name [trace.ERR_NAME]. Errors can be casted
// to this type.
//
//	_, err := trace.GetFromHTTPRequest(nil)
//	if err != nil {
//		e := err.(*errors.GenericError)
//	}
func GetFromHTTPRequest(req *http.Request) (trace Trace, err error) {
	trace = Trace{}

	if req == nil {
		return trace, errors.NewWithName(
			"trace_error",
			"failed to load trace from http request: req is nil",
		)
	}

	if req.Header == nil {
		return trace, errors.NewWithName(
			"trace_error",
			"failed to load trace from http request: req.Header is nil",
		)
	}

	for k := range req.Header {
		if !strings.HasPrefix(k, HTTP_HEADER_PREFIX) {
			continue
		}

		newK := strings.ReplaceAll(k, HTTP_HEADER_PREFIX, "")
		newK = strings.ReplaceAll(newK, "-", "_")
		newK = strings.ToLower(newK)
		trace.Set(newK, req.Header.Get(k))
	}

	return trace, nil
}
