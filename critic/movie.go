package critic

import "time"

// Movie represents a movie.
type Movie struct {
	ID        int64      // The movie ID
	Title     string     // The movie title
	Year      string     // The movie release year
	CreatedAt time.Time  // The time of creation
	UpdatedAt *time.Time // The time of last update
	DeletedAt *time.Time // The time of deletion
}
