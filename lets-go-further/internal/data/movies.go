package data

import "time"

// Movie The fields are exported which is necessary for them to be visible to `encoding/json` package.
// Any fields that aren't exported won't be included when encoding a struct to JSON.
//
// Movie Use the - (hyphen) directive to hide fields, use omitempty to hide fields only if the field value is empty:
// - `false`, `0`, `""`
// - empty array, slice or map
// - nil pointer or nil interface value
type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}
