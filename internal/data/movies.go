package data

import (
	"time"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitzero"`
	// Use the Runtime type instead of int32. Note that the omitzero directive will
	// still work on this: if the Runtime field has the underlying value 0, then it will
	// be considered zero and omitted -- and the MarshalJSON() method we just made
	// won't be called at all.
	Runtime Runtime  `json:"runtime,omitzero"`
	Genres  []string `json:"genres,omitzero"`
	Version int32    `json:"version"`
}
