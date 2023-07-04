// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
)

type AlertRule struct {
	Strategy AlertStrategy `json:"strategy,omitempty"`
	type_    string
}

func (a *AlertRule) Type() string {
	return a.type_
}

func (a *AlertRule) UnmarshalJSON(data []byte) error {
	type unmarshaler AlertRule
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AlertRule(value)
	a.type_ = "alert"
	return nil
}

func (a *AlertRule) MarshalJSON() ([]byte, error) {
	type embed AlertRule
	var marshaler = struct {
		embed
		Type string `json:"type,omitempty"`
	}{
		embed: embed(*a),
		Type:  "alert",
	}
	return json.Marshal(marshaler)
}
