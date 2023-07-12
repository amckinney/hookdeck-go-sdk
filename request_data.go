// This file was auto-generated by Fern from our API Definition.

package api

type RequestData struct {
	Path           string                  `json:"path"`
	Query          *string                 `json:"query,omitempty"`
	ParsedQuery    *RequestDataParsedQuery `json:"parsed_query,omitempty"`
	Headers        *RequestDataHeaders     `json:"headers,omitempty"`
	Body           *RequestDataBody        `json:"body,omitempty"`
	IsLargePayload *bool                   `json:"is_large_payload,omitempty"`
}
