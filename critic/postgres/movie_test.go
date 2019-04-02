package postgres_test

import (
	"os"

	"github.com/nickmro/movie-critic-backend/critic"

	. "github.com/nickmro/movie-critic-backend/critic/postgres"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MovieDatabase", func() {
	var db *MovieDatabase

	BeforeEach(func() {
		db = &MovieDatabase{
			DB:          NewDB(os.Getenv("DB_URL")),
			ErrorLogger: &Logger{},
		}
	})

	AfterEach(func() {
		db.Close()
	})

	Describe("Create", func() {
		var tx *MovieTx

		BeforeEach(func() {
			tx = db.BeginTx().(*MovieTx)
		})

		AfterEach(func() {
			tx.Rollback()
		})

		It("creates a movie", func() {
			movie := &critic.Movie{
				Title: "Jaws",
				Year:  "1975",
			}

			err := tx.Create(movie)
			Expect(err).ToNot(HaveOccurred())
			Expect(movie.ID).ToNot(BeZero())
		})
	})

	Describe("Update", func() {
		var tx *MovieTx

		BeforeEach(func() {
			tx = db.BeginTx().(*MovieTx)
		})

		AfterEach(func() {
			tx.Rollback()
		})

		It("updates a movie", func() {
			movie := &critic.Movie{
				Title: "Jaws",
				Year:  "1976",
			}

			tx.Create(movie)

			movie.Year = "1975"

			err := tx.Update(movie)
			Expect(err).ToNot(HaveOccurred())
			Expect(movie.UpdatedAt).ToNot(BeNil())
		})
	})

	Describe("Delete", func() {
		var tx *MovieTx

		BeforeEach(func() {
			tx = db.BeginTx().(*MovieTx)
		})

		AfterEach(func() {
			tx.Rollback()
		})

		It("deletes a movie", func() {
			movie := &critic.Movie{
				Title: "Jaws",
				Year:  "1975",
			}

			tx.Create(movie)
			err := tx.Delete(movie.ID)
			Expect(err).ToNot(HaveOccurred())

			_, err = db.Get(movie.ID)
			Expect(err).To(Equal(ErrNotFound))
		})
	})

	Describe("Query", func() {
		var tx *MovieTx
		var records = []*critic.Movie{
			&critic.Movie{
				Title: "Jaws",
				Year:  "1975",
			},
			&critic.Movie{
				Title: "Citizen Kane",
				Year:  "1941",
			},
		}

		BeforeEach(func() {
			tx = db.BeginTx().(*MovieTx)
			for _, record := range records {
				tx.Create(record)
			}
			tx.Commit()
		})

		AfterEach(func() {
			tx = db.BeginTx().(*MovieTx)
			for _, record := range records {
				tx.Delete(record.ID)
			}
			tx.Commit()
		})

		It("returns a list of movies", func() {
			movies, err := db.Query(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(movies).ToNot(BeNil())
			Expect(len(movies)).To(Equal(2))
		})

		Context("when a before ID is provided", func() {
			It("only returns movies before the given ID", func() {
				movies, err := db.Query(map[critic.MovieQueryParam]interface{}{
					critic.MovieQueryParamBefore: records[1].ID,
				})

				Expect(err).ToNot(HaveOccurred())
				Expect(movies).ToNot(BeNil())
				Expect(movies).To(HaveLen(1))
				Expect(movies[0].ID).To(Equal(records[0].ID))
			})
		})

		Context("when a limit is provided", func() {
			It("limits the number of movies returned", func() {
				movies, err := db.Query(map[critic.MovieQueryParam]interface{}{
					critic.MovieQueryParamLimit: 1,
				})

				Expect(err).ToNot(HaveOccurred())
				Expect(movies).ToNot(BeNil())
				Expect(movies).To(HaveLen(1))
				Expect(movies[0].ID).To(Equal(records[1].ID))
			})
		})
	})

	Describe("Get", func() {
		var tx *MovieTx
		var record = &critic.Movie{
			Title: "Jaws",
			Year:  "1975",
		}

		BeforeEach(func() {
			tx = db.BeginTx().(*MovieTx)
			tx.Create(record)
			tx.Commit()
		})

		AfterEach(func() {
			tx = db.BeginTx().(*MovieTx)
			tx.Delete(record.ID)
			tx.Commit()
		})

		It("returns a movie", func() {
			movie, err := db.Get(record.ID)
			Expect(err).ToNot(HaveOccurred())
			Expect(movie).ToNot(BeNil())
			Expect(movie.ID).To(Equal(record.ID))
		})

		Context("when the movie does not exist", func() {
			It("returns a not found error", func() {
				_, err := db.Get(0)
				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(ErrNotFound))
			})
		})
	})
})
