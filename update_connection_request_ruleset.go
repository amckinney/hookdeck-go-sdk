// This file was auto-generated by Fern from our API Definition.

package api

// Ruleset input object
type UpdateConnectionRequestRuleset struct {
	// Name for the ruleset <span style="white-space: nowrap">`<= 155 characters`</span>
	Name string `json:"name,omitempty"`
	// Array of rules to apply
	Rules         *[]*Rule `json:"rules,omitempty"`
	IsTeamDefault *bool    `json:"is_team_default,omitempty"`
}
