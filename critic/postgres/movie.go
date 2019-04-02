package postgres

import (
	"database/sql"
	"fmt"
	"time"

	backend "github.com/nickmro/movie-critic-backend"

	"github.com/nickmro/movie-critic-backend/critic"
)

// MovieDatabase represents a PostgreSQL movie database.
type MovieDatabase struct {
	*DB
	ErrorLogger backend.ErrorLogger
}

// MovieTx represents a movie update transaction.
type MovieTx struct {
	*Tx
	ErrorLogger backend.ErrorLogger
}

const (
	moviesTableName = "movies"
)

// BeginTx begins a movie transaction.
func (db *MovieDatabase) BeginTx() critic.MovieTx {
	return &MovieTx{
		Tx:          db.MustBegin(),
		ErrorLogger: db.ErrorLogger,
	}
}

// Query returns a list of movies.
func (db *MovieDatabase) Query(params map[critic.MovieQueryParam]interface{}) ([]*critic.Movie, error) {
	movies := []*critic.Movie{}

	rows, err := db.DB.Query(movieQuery(params))
	if err != nil {
		go LogError(db.ErrorLogger, err)
		return nil, ErrInternal
	}
	defer rows.Close()

	for rows.Next() {
		movie := &critic.Movie{}
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Year,
			&movie.CreatedAt,
			&movie.UpdatedAt,
			&movie.DeletedAt,
		)
		if err != nil {
			go LogError(db.ErrorLogger, err)
			return nil, ErrInternal
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

// Get returns a movie or an error.
func (db *MovieDatabase) Get(id int64) (*critic.Movie, error) {
	var movie critic.Movie

	query := `
		SELECT
			movies.id,
			movies.title,
			movies.year,
			movies.created_at,
			movies.updated_at,
			movies.deleted_at
		FROM
			movies
		WHERE
			movies.id = $1
			AND movies.deleted_at IS NULL
	`

	err := db.DB.QueryRow(query, id).Scan(
		&movie.ID,
		&movie.Title,
		&movie.Year,
		&movie.CreatedAt,
		&movie.UpdatedAt,
		&movie.DeletedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}

		go LogError(db.ErrorLogger, err)
		return nil, ErrInternal
	}

	return &movie, nil
}

// Create creates a movie.
func (tx *MovieTx) Create(movie *critic.Movie) error {
	query := `
		INSERT INTO movies (title, year, created_at)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	movie.CreatedAt = time.Now()

	err := tx.QueryRow(query, movie.Title, movie.Year, movie.CreatedAt).Scan(&movie.ID)
	if err != nil {
		go LogError(tx.ErrorLogger, err)
		return ErrInternal
	}

	return nil
}

// Update updates a movie.
func (tx *MovieTx) Update(movie *critic.Movie) error {
	query := `
		UPDATE movies
		SET title=$2, year=$3, updated_at=$4
		WHERE id = $1
	`

	now := time.Now()
	movie.UpdatedAt = &now

	_, err := tx.Exec(query, movie.ID, movie.Title, movie.Year, movie.UpdatedAt)
	if err != nil {
		go LogError(tx.ErrorLogger, err)
		return ErrInternal
	}

	return nil
}

// Delete deletes a movie.
func (tx *MovieTx) Delete(id int64) error {
	query := `
		UPDATE movies
		SET deleted_at=$2
		WHERE id = $1
	`

	_, err := tx.Exec(query, id, time.Now())
	if err != nil {
		go LogError(tx.ErrorLogger, err)
		return ErrInternal
	}

	return nil
}

func movieQuery(params map[critic.MovieQueryParam]interface{}) (query string, args []interface{}) {
	// select clause
	query = `
		SELECT
			movies.id,
			movies.title,
			movies.year,
			movies.created_at,
			movies.updated_at,
			movies.deleted_at
		FROM movies
	`
	args = []interface{}{}

	// where clauses
	if params != nil {
		wheres := []string{`movies.deleted_at IS NULL`}
		if before, ok := params[critic.MovieQueryParamBefore]; ok {
			args = append(args, before)
			wheres = append(wheres, fmt.Sprintf(`movies.id < $%d`, len(args)))
		}

		for i := range wheres {
			if i == 0 {
				query += fmt.Sprintf(` WHERE %s`, wheres[i])
			} else {
				query += fmt.Sprintf(` AND %s`, wheres[i])
			}
		}
	}

	// order by clause
	query += ` ORDER BY movies.id DESC`

	// limit clause
	if params != nil {
		if limit, ok := params[critic.MovieQueryParamLimit]; ok {
			args = append(args, limit)
			query += fmt.Sprintf(` LIMIT $%d`, len(args))
		}
	}

	return query, args
}
