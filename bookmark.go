// This file was auto-generated by Fern from our API Definition.

package api

type Bookmark struct {
	// ID of the bookmark
	Id string `json:"id,omitempty"`
	// ID of the workspace
	TeamId string `json:"team_id,omitempty"`
	// ID of the associated connection
	WebhookId string `json:"webhook_id,omitempty"`
	// ID of the bookmarked event data
	EventDataId string `json:"event_data_id,omitempty"`
	// Descriptive name of the bookmark
	Label string `json:"label,omitempty"`
	// Alternate alias for the bookmark
	Alias *string         `json:"alias,omitempty"`
	Data  *ShortEventData `json:"data,omitempty"`
	// Date the bookmark was last manually triggered
	LastUsedAt *string `json:"last_used_at,omitempty"`
	// Date the bookmark was last updated
	UpdatedAt string `json:"updated_at,omitempty"`
	// Date the bookmark was created
	CreatedAt string `json:"created_at,omitempty"`
}
