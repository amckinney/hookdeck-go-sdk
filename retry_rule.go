// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
)

type RetryRule struct {
	Strategy RetryStrategy `json:"strategy,omitempty"`
	// Time in MS between each retry
	Interval *int `json:"interval,omitempty"`
	// Maximum number of retries to attempt
	Count *int `json:"count,omitempty"`
	type_ string
}

func (r *RetryRule) Type() string {
	return r.type_
}

func (r *RetryRule) UnmarshalJSON(data []byte) error {
	type unmarshaler RetryRule
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RetryRule(value)
	r.type_ = "retry"
	return nil
}

func (r *RetryRule) MarshalJSON() ([]byte, error) {
	type embed RetryRule
	var marshaler = struct {
		embed
		Type string `json:"type,omitempty"`
	}{
		embed: embed(*r),
		Type:  "retry",
	}
	return json.Marshal(marshaler)
}
