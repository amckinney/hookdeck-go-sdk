// This file was auto-generated by Fern from our API Definition.

package api

type ToggleWebhookNotifications struct {
	Enabled  bool           `json:"enabled,omitempty"`
	Topics   *[]TopicsValue `json:"topics,omitempty"`
	SourceId string         `json:"source_id,omitempty"`
}
