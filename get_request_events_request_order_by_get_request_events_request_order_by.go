// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type GetRequestEventsRequestOrderByGetRequestEventsRequestOrderBy uint8

const (
	GetRequestEventsRequestOrderByGetRequestEventsRequestOrderByLastAttemptAt GetRequestEventsRequestOrderByGetRequestEventsRequestOrderBy = iota + 1
	GetRequestEventsRequestOrderByGetRequestEventsRequestOrderByNextAttemptAt
	GetRequestEventsRequestOrderByGetRequestEventsRequestOrderByCreatedAt
)

func (g GetRequestEventsRequestOrderByGetRequestEventsRequestOrderBy) String() string {
	switch g {
	default:
		return strconv.Itoa(int(g))
	case GetRequestEventsRequestOrderByGetRequestEventsRequestOrderByLastAttemptAt:
		return "last_attempt_at"
	case GetRequestEventsRequestOrderByGetRequestEventsRequestOrderByNextAttemptAt:
		return "next_attempt_at"
	case GetRequestEventsRequestOrderByGetRequestEventsRequestOrderByCreatedAt:
		return "created_at"
	}
}

func (g GetRequestEventsRequestOrderByGetRequestEventsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", g.String())), nil
}

func (g *GetRequestEventsRequestOrderByGetRequestEventsRequestOrderBy) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "last_attempt_at":
		value := GetRequestEventsRequestOrderByGetRequestEventsRequestOrderByLastAttemptAt
		*g = value
	case "next_attempt_at":
		value := GetRequestEventsRequestOrderByGetRequestEventsRequestOrderByNextAttemptAt
		*g = value
	case "created_at":
		value := GetRequestEventsRequestOrderByGetRequestEventsRequestOrderByCreatedAt
		*g = value
	}
	return nil
}
