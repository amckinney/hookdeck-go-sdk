// This file was auto-generated by Fern from our API Definition.

package api

type CreateRequestBulkRetryRequest struct {
	// Filter properties for the events to be included in the bulk retry, use query parameters of [Requests](#requests)
	Query *CreateRequestBulkRetryRequestQuery `json:"query,omitempty"`
}
