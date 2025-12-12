# Go Trace Package

A simple Go package for propagating trace data through execution contexts. It provides a `Trace` type with helpers for context integration, HTTP headers, and JSON serialization.

## Installation

```bash
go get github.com/iolave/go-trace
```

## Usage

### Creating a Trace

The `Trace` type is a `map[string]string`. The `Set` method enforces that keys are in `snake_case`.

```go
package main

import (
	"fmt"
	"github.com/iolave/go-trace"
)

func main() {
	// Create a new trace
	t := make(trace.Trace)
    // or:
    //  t := trace.Trace{}

	// Set some data
	err := t.Set("request_id", "xyz-123")
	if err != nil {
		panic(err)
	}

	// Set enforces snake_case keys
	err = t.Set("invalidKey", "some-value")
	if err != nil {
		// Error: trace key must be in snake case, (got invalidKey, expected invalid_key)
		fmt.Println("Error:", err)
	}

	fmt.Println(t.Get("request_id")) // xyz-123
}
```

### Context Integration

You can easily set and get a trace from a `context.Context`.

```go
package main

import (
	"context"
	"fmt"
	"github.com/iolave/go-trace"
)

func main() {
    ctx := context.Background()

    // Get the trace from the context
	t := trace.GetFromContext(ctx)

    // Add more data to the trace
    t.Set("service_name", "my-service")

	// Set the trace in a context
	ctx = t.SetInContext(context.Background())
    
    // Later in a downstream function
    t = trace.GetFromContext(ctx)
    fmt.Println(t.Get("service_name")) // my-service
}
```

### HTTP Request Integration

The package provides helpers to inject and extract trace data from HTTP headers, prefixed with `X-Trace-`.

```go
package main

import (
	"fmt"
	"github.com/iolave/go-trace"
	"net/http"
	"net/http/httptest"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get trace from incoming request headers
		t, err := trace.GetFromHTTPRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Printf("Trace from request: %v\n", t)

		// Add more data to the trace
		t.Set("service_name", "my-service")

		// Set trace headers for an outgoing response (or a new client request)
		t.SetHTTPHeaders(w.Header())
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Hello, client")
	})

	// Simulate an incoming request from an upstream service
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	req.Header.Set("X-Trace-Request-Id", "xyz-abc-123")
	req.Header.Set("X-Trace-Source", "api-gateway")


	w := httptest.NewRecorder()
	handler(w, req)

	fmt.Printf("Response Headers: %v\n", w.Header())
}
```

### JSON Serialization

The `Trace` type implements the `json.Marshaler` interface.

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/iolave/go-trace"
)

func main() {
	t := make(trace.Trace)
	t.Set("request_id", "xyz-123")
	t.Set("user_id", "42")

	jsonData, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}

	// {"request_id":"xyz-123","user_id":"42"}
	fmt.Println(string(jsonData))
}
```
