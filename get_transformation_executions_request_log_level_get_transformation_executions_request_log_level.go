// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type GetTransformationExecutionsRequestLogLevelGetTransformationExecutionsRequestLogLevel uint8

const (
	GetTransformationExecutionsRequestLogLevelGetTransformationExecutionsRequestLogLevelDebug GetTransformationExecutionsRequestLogLevelGetTransformationExecutionsRequestLogLevel = iota + 1
	GetTransformationExecutionsRequestLogLevelGetTransformationExecutionsRequestLogLevelInfo
	GetTransformationExecutionsRequestLogLevelGetTransformationExecutionsRequestLogLevelWarn
	GetTransformationExecutionsRequestLogLevelGetTransformationExecutionsRequestLogLevelError
	GetTransformationExecutionsRequestLogLevelGetTransformationExecutionsRequestLogLevelFatal
)

func (g GetTransformationExecutionsRequestLogLevelGetTransformationExecutionsRequestLogLevel) String() string {
	switch g {
	default:
		return strconv.Itoa(int(g))
	case GetTransformationExecutionsRequestLogLevelGetTransformationExecutionsRequestLogLevelDebug:
		return "debug"
	case GetTransformationExecutionsRequestLogLevelGetTransformationExecutionsRequestLogLevelInfo:
		return "info"
	case GetTransformationExecutionsRequestLogLevelGetTransformationExecutionsRequestLogLevelWarn:
		return "warn"
	case GetTransformationExecutionsRequestLogLevelGetTransformationExecutionsRequestLogLevelError:
		return "error"
	case GetTransformationExecutionsRequestLogLevelGetTransformationExecutionsRequestLogLevelFatal:
		return "fatal"
	}
}

func (g GetTransformationExecutionsRequestLogLevelGetTransformationExecutionsRequestLogLevel) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", g.String())), nil
}

func (g *GetTransformationExecutionsRequestLogLevelGetTransformationExecutionsRequestLogLevel) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "debug":
		value := GetTransformationExecutionsRequestLogLevelGetTransformationExecutionsRequestLogLevelDebug
		*g = value
	case "info":
		value := GetTransformationExecutionsRequestLogLevelGetTransformationExecutionsRequestLogLevelInfo
		*g = value
	case "warn":
		value := GetTransformationExecutionsRequestLogLevelGetTransformationExecutionsRequestLogLevelWarn
		*g = value
	case "error":
		value := GetTransformationExecutionsRequestLogLevelGetTransformationExecutionsRequestLogLevelError
		*g = value
	case "fatal":
		value := GetTransformationExecutionsRequestLogLevelGetTransformationExecutionsRequestLogLevelFatal
		*g = value
	}
	return nil
}
