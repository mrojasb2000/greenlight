package data

import "time"

// Annotate the Movie struct with struct tags to control how the keys appear in the
// JSON-encoded output.
type Movie struct {
	ID       int64     `json:"id"`             // Unique integer ID for the movie
	CreateAt time.Time `json:"-"`              // Timestamp for when the movie is added to our database
	Title    string    `json:"title"`          // Movie title
	Year     int32     `json:"year,omitempty"` // Movie release year
	// Use the Runtime type instead of int32. Note that the omitempty directive will
	// still work on this: if the Runtime field has the underlying value 0, then it will
	// be considered empty and omitted -- and the MarshalJSON() method we just made
	// won.t be called at all.
	//Runtime  int32     `json:"runtime,omitempty,string"` // Movie runtime (in minutes)
	Runtime Runtime  `json:"runtime, omitempty"`
	Genres  []string `json:"genres,omitempty"` // Slice of geners for the movie (romace, comedy, etc.)
	Version int32    `json:"version"`          // The version number start 1 and will be incremented each time the movie information is updated
}
