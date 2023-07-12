// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
)

// Transformation issue
type TransformationIssueWithData struct {
	// Issue ID
	Id string `json:"id"`
	// ID of the workspace
	TeamId string      `json:"team_id"`
	Status IssueStatus `json:"status,omitempty"`
	// ISO timestamp for when the issue was last opened
	OpenedAt string `json:"opened_at"`
	// ISO timestamp for when the issue was first opened
	FirstSeenAt string `json:"first_seen_at"`
	// ISO timestamp for when the issue last occured
	LastSeenAt string `json:"last_seen_at"`
	// ID of the team member who last updated the issue status
	LastUpdatedBy *string `json:"last_updated_by,omitempty"`
	// ISO timestamp for when the issue was dismissed
	DismissedAt    *string `json:"dismissed_at,omitempty"`
	AutoResolvedAt *string `json:"auto_resolved_at,omitempty"`
	MergedWith     *string `json:"merged_with,omitempty"`
	// ISO timestamp for when the issue was last updated
	UpdatedAt string `json:"updated_at"`
	// ISO timestamp for when the issue was created
	CreatedAt       string                              `json:"created_at"`
	AggregationKeys *TransformationIssueAggregationKeys `json:"aggregation_keys,omitempty"`
	Reference       *TransformationIssueReference       `json:"reference,omitempty"`
	Data            *TransformationIssueData            `json:"data,omitempty"`
	type_           string
}

func (t *TransformationIssueWithData) Type() string {
	return t.type_
}

func (t *TransformationIssueWithData) UnmarshalJSON(data []byte) error {
	type unmarshaler TransformationIssueWithData
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TransformationIssueWithData(value)
	t.type_ = "transformation"
	return nil
}

func (t *TransformationIssueWithData) MarshalJSON() ([]byte, error) {
	type embed TransformationIssueWithData
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*t),
		Type:  "transformation",
	}
	return json.Marshal(marshaler)
}
