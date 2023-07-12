// This file was auto-generated by Fern from our API Definition.

package api

// UpdateDestinationRequest is an in-lined request used by the UpdateDestination endpoint.
type UpdateDestinationRequest struct {
	// Name for the destination <span style="white-space: nowrap">`<= 155 characters`</span>
	Name string `json:"name"`
	// Endpoint of the destination
	Url *string `json:"url,omitempty"`
	// Path for the CLI destination
	CliPath *string `json:"cli_path,omitempty"`
	// Period to rate limit attempts
	RateLimitPeriod *UpdateDestinationRequestRateLimitPeriod `json:"rate_limit_period,omitempty"`
	// Limit event attempts to receive per period
	RateLimit *int `json:"rate_limit,omitempty"`
	// Date the destination was archived
	ArchivedAt             *string `json:"archived_at,omitempty"`
	PathForwardingDisabled *bool   `json:"path_forwarding_disabled,omitempty"`
}
