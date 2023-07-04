// This file was auto-generated by Fern from our API Definition.

package api

type IssueTrigger struct {
	// ID of the issue trigger
	Id string `json:"id,omitempty"`
	// ID of the workspace
	TeamId *string `json:"team_id,omitempty"`
	// Optional unique name to use as reference when using the API
	Name     *string                `json:"name,omitempty"`
	Type     IssueType              `json:"type,omitempty"`
	Configs  *IssueTriggerReference `json:"configs,omitempty"`
	Channels *IssueTriggerChannels  `json:"channels,omitempty"`
	// ISO timestamp for when the issue trigger was disabled
	DisabledAt *string `json:"disabled_at,omitempty"`
	// ISO timestamp for when the issue trigger was last updated
	UpdatedAt string `json:"updated_at,omitempty"`
	// ISO timestamp for when the issue trigger was created
	CreatedAt string `json:"created_at,omitempty"`
	// ISO timestamp for when the issue trigger was deleted
	DeletedAt *string `json:"deleted_at,omitempty"`
}
