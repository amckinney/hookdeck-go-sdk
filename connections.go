// This file was auto-generated by Fern from our API Definition.

package api

import (
	time "time"
)

type CreateConnectionRequest struct {
	// A unique name of the connection for the source
	Name *string `json:"name,omitempty"`
	// Description for the connection
	Description *string `json:"description,omitempty"`
	// ID of a destination to bind to the connection
	DestinationId *string `json:"destination_id,omitempty"`
	// ID of a source to bind to the connection
	SourceId *string `json:"source_id,omitempty"`
	// Destination input object
	Destination *CreateConnectionRequestDestination `json:"destination,omitempty"`
	// Source input object
	Source *CreateConnectionRequestSource `json:"source,omitempty"`
	Rules  []*Rule                        `json:"rules,omitempty"`
}

type GetConnectionsRequest struct {
	Id            *string                       `json:"-"`
	Name          *string                       `json:"-"`
	DestinationId *string                       `json:"-"`
	SourceId      *string                       `json:"-"`
	Archived      *bool                         `json:"-"`
	ArchivedAt    *time.Time                    `json:"-"`
	FullName      *string                       `json:"-"`
	PausedAt      *time.Time                    `json:"-"`
	OrderBy       *GetConnectionsRequestOrderBy `json:"-"`
	Dir           *GetConnectionsRequestDir     `json:"-"`
	Limit         *int                          `json:"-"`
	Next          *string                       `json:"-"`
	Prev          *string                       `json:"-"`
}

type UpdateConnectionRequest struct {
	// <span style="white-space: nowrap">`<= 155 characters`</span>
	Name *string `json:"name,omitempty"`
	// Description for the connection
	Description *string `json:"description,omitempty"`
	Rules       []*Rule `json:"rules,omitempty"`
}

type UpsertConnectionRequest struct {
	// A unique name of the connection for the source
	Name *string `json:"name,omitempty"`
	// Description for the connection
	Description *string `json:"description,omitempty"`
	// ID of a destination to bind to the connection
	DestinationId *string `json:"destination_id,omitempty"`
	// ID of a source to bind to the connection
	SourceId *string `json:"source_id,omitempty"`
	// Destination input object
	Destination *UpsertConnectionRequestDestination `json:"destination,omitempty"`
	// Source input object
	Source *UpsertConnectionRequestSource `json:"source,omitempty"`
	Rules  []*Rule                        `json:"rules,omitempty"`
}
