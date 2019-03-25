package critic

import "time"

// Review prepresents a movie review.
type Review struct {
	ID        string     // The rating ID
	UserID    int64      // The rating user ID
	MovieID   int64      // The movie ID
	Score     float64    // The rating score
	Text      string     // The rating text
	CreatedAt time.Time  // The creation time
	UpdatedAt *time.Time // The time of last update
	DeletedAt *time.Time // The time of deletion
}
