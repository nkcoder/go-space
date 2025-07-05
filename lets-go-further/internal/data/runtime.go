package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Runtime int32

// MarshalJSON Strictly speaking, when Go is encoding a particular type to JSON it looks to see if the type
// satisfies the json.Marshaler interface.
// If yes, then Go will call its MarshalJSON() method and use the []byte slice that it returns as the encoded JSON value.
// If no, then Go will fall back to trying to encode it to JSON based on its own internal set of rules.
func (r *Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	// It needs to be surrounded by double quotes in order bo be a valid JSON string
	quotedJsonValue := strconv.Quote(jsonValue)

	return []byte(quotedJsonValue), nil
}

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

// UnmarshalJSON Implement this method on the type so that it satisfies the json.Unmarshaler interface.
// Note: because UnmarshalJSON needs to modify the receiver, we must use a pointer receiver,
// otherwise we will only be modifying a copy which is then discarded when the method returns.
func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	unQuotedJsonValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	// the format is "180 mins"
	parts := strings.Split(unQuotedJsonValue, " ")
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(i)
	return nil
}
