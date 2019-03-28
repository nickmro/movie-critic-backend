package postgres

import "github.com/nickmro/movie-critic-backend/critic"

// MovieDatabase represents a PostgreSQL movie database.
type MovieDatabase struct {
	*DB
}

// MovieTx represents a movie update transaction.
type MovieTx struct {
	*Tx
}

const (
	moviesTableName = "movies"
)

// Begin begins a movie transaction.
func (db *MovieDatabase) Begin() critic.MovieTx {
	return &MovieTx{db.MustBegin()}
}

// Query returns a list of movies.
func (db *MovieDatabase) Query(params critic.MovieQueryParams, options critic.MovieListOptions) ([]*critic.Movie, error) {
	return nil, nil
}

// Get returns a movie or an error.
func (db *MovieDatabase) Get(id int64) (*critic.Movie, error) {
	return nil, nil
}

// Create creates a movie.
func (tx *MovieTx) Create(movie *critic.Movie) error {
	return nil
}

// Update updates a movie.
func (tx *MovieTx) Update(movie *critic.Movie) error {
	return nil
}

// Delete deletes a movie.
func (tx *MovieTx) Delete(id int64) error {
	return nil
}
