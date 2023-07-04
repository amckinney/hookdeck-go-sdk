// This file was auto-generated by Fern from our API Definition.

package api

type UpsertConnectionRequest struct {
	// A unique name of the connection for the source <span style="white-space: nowrap">`<= 155 characters`</span>
	Name string `json:"name,omitempty"`
	// ID of a destination to bind to the connection
	DestinationId *string `json:"destination_id,omitempty"`
	// ID of a source to bind to the connection
	SourceId *string `json:"source_id,omitempty"`
	// Destination input object
	Destination *UpsertConnectionRequestDestination `json:"destination,omitempty"`
	// Source input object
	Source *UpsertConnectionRequestSource `json:"source,omitempty"`
	// Ruleset input object
	Ruleset *UpsertConnectionRequestRuleset `json:"ruleset,omitempty"`
	// ID of a rule to bind to the connection. Default to the Workspace default ruleset
	RulesetId *string `json:"ruleset_id,omitempty"`
	// Array of rules to apply
	Rules *[]*Rule `json:"rules,omitempty"`
}
