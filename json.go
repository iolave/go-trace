package trace

import "encoding/json"

// MarshalJSON returns the JSON encoding of the trace.
//
// If the trace is nil, it will return an empty JSON object.
func (t Trace) MarshalJSON() ([]byte, error) {
	var tmap map[string]string = t

	if tmap == nil {
		return []byte("{}"), nil
	}

	return json.Marshal(tmap)
}
