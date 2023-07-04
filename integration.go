// This file was auto-generated by Fern from our API Definition.

package api

type Integration struct {
	// ID of the integration
	Id string `json:"id,omitempty"`
	// ID of the workspace
	TeamId string `json:"team_id,omitempty"`
	// Label of the integration
	Label    string              `json:"label,omitempty"`
	Provider IntegrationProvider `json:"provider,omitempty"`
	// List of features to enable (see features list below)
	Features []IntegrationFeature `json:"features,omitempty"`
	// Decrypted Key/Value object of the associated configuration for that provider
	Configs *IntegrationConfigs `json:"configs,omitempty"`
	// List of source IDs the integration is attached to
	Sources []string `json:"sources,omitempty"`
	// Date the integration was last updated
	UpdatedAt string `json:"updated_at,omitempty"`
	// Date the integration was created
	CreatedAt string `json:"created_at,omitempty"`
}
