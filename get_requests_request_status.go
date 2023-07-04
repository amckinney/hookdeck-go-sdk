// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type GetRequestsRequestStatus uint8

const (
	GetRequestsRequestStatusAccepted GetRequestsRequestStatus = iota + 1
	GetRequestsRequestStatusRejected
)

func (g GetRequestsRequestStatus) String() string {
	switch g {
	default:
		return strconv.Itoa(int(g))
	case GetRequestsRequestStatusAccepted:
		return "accepted"
	case GetRequestsRequestStatusRejected:
		return "rejected"
	}
}

func (g GetRequestsRequestStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", g.String())), nil
}

func (g *GetRequestsRequestStatus) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "accepted":
		value := GetRequestsRequestStatusAccepted
		*g = value
	case "rejected":
		value := GetRequestsRequestStatusRejected
		*g = value
	}
	return nil
}
