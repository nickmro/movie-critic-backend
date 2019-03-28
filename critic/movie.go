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

// MovieQueryParams are the parameters used to query movies.
type MovieQueryParams struct {
	BeforeID int64 // The maximum id to return
}

// MovieListOptions are the options used to list movies.
type MovieListOptions struct {
	Limit uint64
}

// MovieRepository defines the operations that may be performed on a movie repository.
type MovieRepository interface {
	Begin() MovieTx
	Query(params MovieQueryParams, options MovieListOptions) ([]*Movie, error)
	Get(id int64) (*Movie, error)
}

// MovieTx defines the operations that may be used to update movies.
type MovieTx interface {
	Create(movie *Movie) error
	Update(movie *Movie) error
	Delete(id int64) error
	Commit() error
	Rollback() error
}
