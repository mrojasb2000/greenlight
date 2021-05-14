package data

import "time"

type Movie struct {
	ID       int64     // Unique integer ID for the movie
	CreateAt time.Time // Timestamp for when the movie is added to our database
	Title    string    // Movie title
	Year     int32     // Movie release year
	Runtime  int32     // Movie runtime (in minutes)
	Genres   []string  // Slice of geners for the movie (romace, comedy, etc.)
	Version  int32     // The version number start 1 and will be incremented each time the movie information is updated
}
