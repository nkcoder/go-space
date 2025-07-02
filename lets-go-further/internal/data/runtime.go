package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

// MarshalJSON Strictly speaking, when Go is encoding a particular type to JSON it looks to see if the type
// satisfies the json.Marshaler interface.
// If yes, then Go will call its MarshalJSON() method and use the []byte slice that it returns as the encoded JSON value.
// If no, then Go will fall back to trying to encode it to JSON based on its own internal set of rules.
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	// It needs to be surrounded by double quotes in order bo be a valid JSON string
	quotedJsonValue := strconv.Quote(jsonValue)

	return []byte(quotedJsonValue), nil
}
