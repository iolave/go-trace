package trace_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/iolave/go-trace"
)

func TestTrace_SetHTTPHeaders(t *testing.T) {
	t.Run("should do nothing if the header param is nil", func(t *testing.T) {
		trace := trace.Trace{}
		trace.SetHTTPHeaders(nil)
	})

	t.Run("should set the trace data in the header", func(t *testing.T) {
		k := "key"
		v := "value"

		tr := trace.Trace{
			k: v,
		}

		h := http.Header{}
		tr.SetHTTPHeaders(h)

		got := h.Get(fmt.Sprintf("%s%s", trace.HTTP_HEADER_PREFIX, k))
		if got != v {
			t.Fatalf("expected %s, got %s", "value", got)
		}
	})
}

func TestGetFromHTTPRequest(t *testing.T) {
	t.Run("slould return an error if the request is nil", func(t *testing.T) {
		_, err := trace.GetFromHTTPRequest(nil)
		if err == nil {
			t.Fatalf("expected an error, got none")
		}
	})

	t.Run("should return an error if the request header is nil", func(t *testing.T) {
		req := &http.Request{}

		_, err := trace.GetFromHTTPRequest(req)
		if err == nil {
			t.Fatalf("expected an error, got none")
		}
	})

	t.Run("should return a trace with data", func(t *testing.T) {
		k := "key"
		v := "value"

		h := http.Header{}
		h.Set(fmt.Sprintf("%s%s", trace.HTTP_HEADER_PREFIX, k), v)
		h.Set("another", v)

		req := &http.Request{
			Header: h,
		}

		tr, err := trace.GetFromHTTPRequest(req)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		fmt.Println(tr)

		got := tr.Get(k)
		if got != v {
			t.Fatalf("expected %s, got %s", v, got)
		}
	})
}
