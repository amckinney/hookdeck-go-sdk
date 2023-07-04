// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
)

// Filter by `last_attempt_at` date using a date operator
type CreateEventBulkRetryRequestQueryLastAttemptAt struct {
	typeName                                                                                   string
	String                                                                                     string
	CreateEventBulkRetryRequestQueryLastAttemptAtCreateEventBulkRetryRequestQueryLastAttemptAt *CreateEventBulkRetryRequestQueryLastAttemptAtCreateEventBulkRetryRequestQueryLastAttemptAt
}

func NewCreateEventBulkRetryRequestQueryLastAttemptAtFromString(value string) *CreateEventBulkRetryRequestQueryLastAttemptAt {
	return &CreateEventBulkRetryRequestQueryLastAttemptAt{typeName: "string", String: value}
}

func NewCreateEventBulkRetryRequestQueryLastAttemptAtFromCreateEventBulkRetryRequestQueryLastAttemptAtCreateEventBulkRetryRequestQueryLastAttemptAt(value *CreateEventBulkRetryRequestQueryLastAttemptAtCreateEventBulkRetryRequestQueryLastAttemptAt) *CreateEventBulkRetryRequestQueryLastAttemptAt {
	return &CreateEventBulkRetryRequestQueryLastAttemptAt{typeName: "createEventBulkRetryRequestQueryLastAttemptAtCreateEventBulkRetryRequestQueryLastAttemptAt", CreateEventBulkRetryRequestQueryLastAttemptAtCreateEventBulkRetryRequestQueryLastAttemptAt: value}
}

func (c *CreateEventBulkRetryRequestQueryLastAttemptAt) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		c.typeName = "string"
		c.String = valueString
		return nil
	}
	valueCreateEventBulkRetryRequestQueryLastAttemptAtCreateEventBulkRetryRequestQueryLastAttemptAt := new(CreateEventBulkRetryRequestQueryLastAttemptAtCreateEventBulkRetryRequestQueryLastAttemptAt)
	if err := json.Unmarshal(data, &valueCreateEventBulkRetryRequestQueryLastAttemptAtCreateEventBulkRetryRequestQueryLastAttemptAt); err == nil {
		c.typeName = "createEventBulkRetryRequestQueryLastAttemptAtCreateEventBulkRetryRequestQueryLastAttemptAt"
		c.CreateEventBulkRetryRequestQueryLastAttemptAtCreateEventBulkRetryRequestQueryLastAttemptAt = valueCreateEventBulkRetryRequestQueryLastAttemptAtCreateEventBulkRetryRequestQueryLastAttemptAt
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c CreateEventBulkRetryRequestQueryLastAttemptAt) MarshalJSON() ([]byte, error) {
	switch c.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", c.typeName, c)
	case "string":
		return json.Marshal(c.String)
	case "createEventBulkRetryRequestQueryLastAttemptAtCreateEventBulkRetryRequestQueryLastAttemptAt":
		return json.Marshal(c.CreateEventBulkRetryRequestQueryLastAttemptAtCreateEventBulkRetryRequestQueryLastAttemptAt)
	}
}

type CreateEventBulkRetryRequestQueryLastAttemptAtVisitor interface {
	VisitString(string) error
	VisitCreateEventBulkRetryRequestQueryLastAttemptAtCreateEventBulkRetryRequestQueryLastAttemptAt(*CreateEventBulkRetryRequestQueryLastAttemptAtCreateEventBulkRetryRequestQueryLastAttemptAt) error
}

func (c *CreateEventBulkRetryRequestQueryLastAttemptAt) Accept(v CreateEventBulkRetryRequestQueryLastAttemptAtVisitor) error {
	switch c.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", c.typeName, c)
	case "string":
		return v.VisitString(c.String)
	case "createEventBulkRetryRequestQueryLastAttemptAtCreateEventBulkRetryRequestQueryLastAttemptAt":
		return v.VisitCreateEventBulkRetryRequestQueryLastAttemptAtCreateEventBulkRetryRequestQueryLastAttemptAt(c.CreateEventBulkRetryRequestQueryLastAttemptAtCreateEventBulkRetryRequestQueryLastAttemptAt)
	}
}
