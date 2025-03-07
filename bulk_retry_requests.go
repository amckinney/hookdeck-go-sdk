// This file was auto-generated by Fern from our API Definition.

package api

import (
	time "time"
)

type CreateRequestBulkRetryRequest struct {
	// Filter properties for the events to be included in the bulk retry, use query parameters of [Requests](#requests)
	Query *CreateRequestBulkRetryRequestQuery `json:"query,omitempty"`
}

type GetRequestBulkRetriesRequest struct {
	CancelledAt       *time.Time                       `json:"-"`
	CompletedAt       *time.Time                       `json:"-"`
	CreatedAt         *time.Time                       `json:"-"`
	Id                *string                          `json:"-"`
	InProgress        *bool                            `json:"-"`
	QueryPartialMatch *bool                            `json:"-"`
	OrderBy           *string                          `json:"-"`
	Dir               *GetRequestBulkRetriesRequestDir `json:"-"`
	Limit             *int                             `json:"-"`
	Next              *string                          `json:"-"`
	Prev              *string                          `json:"-"`
}
