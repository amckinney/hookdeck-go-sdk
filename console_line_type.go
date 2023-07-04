// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type ConsoleLineType uint8

const (
	ConsoleLineTypeError ConsoleLineType = iota + 1
	ConsoleLineTypeLog
	ConsoleLineTypeWarn
	ConsoleLineTypeInfo
	ConsoleLineTypeDebug
)

func (c ConsoleLineType) String() string {
	switch c {
	default:
		return strconv.Itoa(int(c))
	case ConsoleLineTypeError:
		return "error"
	case ConsoleLineTypeLog:
		return "log"
	case ConsoleLineTypeWarn:
		return "warn"
	case ConsoleLineTypeInfo:
		return "info"
	case ConsoleLineTypeDebug:
		return "debug"
	}
}

func (c ConsoleLineType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", c.String())), nil
}

func (c *ConsoleLineType) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "error":
		value := ConsoleLineTypeError
		*c = value
	case "log":
		value := ConsoleLineTypeLog
		*c = value
	case "warn":
		value := ConsoleLineTypeWarn
		*c = value
	case "info":
		value := ConsoleLineTypeInfo
		*c = value
	case "debug":
		value := ConsoleLineTypeDebug
		*c = value
	}
	return nil
}
