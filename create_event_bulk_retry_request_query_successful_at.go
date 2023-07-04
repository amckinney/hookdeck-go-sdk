// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
)

// Filter by `successful_at` date using a date operator
type CreateEventBulkRetryRequestQuerySuccessfulAt struct {
	typeName                                                                                 string
	String                                                                                   string
	CreateEventBulkRetryRequestQuerySuccessfulAtCreateEventBulkRetryRequestQuerySuccessfulAt *CreateEventBulkRetryRequestQuerySuccessfulAtCreateEventBulkRetryRequestQuerySuccessfulAt
}

func NewCreateEventBulkRetryRequestQuerySuccessfulAtFromString(value string) *CreateEventBulkRetryRequestQuerySuccessfulAt {
	return &CreateEventBulkRetryRequestQuerySuccessfulAt{typeName: "string", String: value}
}

func NewCreateEventBulkRetryRequestQuerySuccessfulAtFromCreateEventBulkRetryRequestQuerySuccessfulAtCreateEventBulkRetryRequestQuerySuccessfulAt(value *CreateEventBulkRetryRequestQuerySuccessfulAtCreateEventBulkRetryRequestQuerySuccessfulAt) *CreateEventBulkRetryRequestQuerySuccessfulAt {
	return &CreateEventBulkRetryRequestQuerySuccessfulAt{typeName: "createEventBulkRetryRequestQuerySuccessfulAtCreateEventBulkRetryRequestQuerySuccessfulAt", CreateEventBulkRetryRequestQuerySuccessfulAtCreateEventBulkRetryRequestQuerySuccessfulAt: value}
}

func (c *CreateEventBulkRetryRequestQuerySuccessfulAt) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		c.typeName = "string"
		c.String = valueString
		return nil
	}
	valueCreateEventBulkRetryRequestQuerySuccessfulAtCreateEventBulkRetryRequestQuerySuccessfulAt := new(CreateEventBulkRetryRequestQuerySuccessfulAtCreateEventBulkRetryRequestQuerySuccessfulAt)
	if err := json.Unmarshal(data, &valueCreateEventBulkRetryRequestQuerySuccessfulAtCreateEventBulkRetryRequestQuerySuccessfulAt); err == nil {
		c.typeName = "createEventBulkRetryRequestQuerySuccessfulAtCreateEventBulkRetryRequestQuerySuccessfulAt"
		c.CreateEventBulkRetryRequestQuerySuccessfulAtCreateEventBulkRetryRequestQuerySuccessfulAt = valueCreateEventBulkRetryRequestQuerySuccessfulAtCreateEventBulkRetryRequestQuerySuccessfulAt
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c CreateEventBulkRetryRequestQuerySuccessfulAt) MarshalJSON() ([]byte, error) {
	switch c.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", c.typeName, c)
	case "string":
		return json.Marshal(c.String)
	case "createEventBulkRetryRequestQuerySuccessfulAtCreateEventBulkRetryRequestQuerySuccessfulAt":
		return json.Marshal(c.CreateEventBulkRetryRequestQuerySuccessfulAtCreateEventBulkRetryRequestQuerySuccessfulAt)
	}
}

type CreateEventBulkRetryRequestQuerySuccessfulAtVisitor interface {
	VisitString(string) error
	VisitCreateEventBulkRetryRequestQuerySuccessfulAtCreateEventBulkRetryRequestQuerySuccessfulAt(*CreateEventBulkRetryRequestQuerySuccessfulAtCreateEventBulkRetryRequestQuerySuccessfulAt) error
}

func (c *CreateEventBulkRetryRequestQuerySuccessfulAt) Accept(v CreateEventBulkRetryRequestQuerySuccessfulAtVisitor) error {
	switch c.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", c.typeName, c)
	case "string":
		return v.VisitString(c.String)
	case "createEventBulkRetryRequestQuerySuccessfulAtCreateEventBulkRetryRequestQuerySuccessfulAt":
		return v.VisitCreateEventBulkRetryRequestQuerySuccessfulAtCreateEventBulkRetryRequestQuerySuccessfulAt(c.CreateEventBulkRetryRequestQuerySuccessfulAtCreateEventBulkRetryRequestQuerySuccessfulAt)
	}
}
