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
		db = &MovieDatabase{DB: NewDB(os.Getenv("DB_URL"))}
	})

	AfterEach(func() {
		db.Close()
	})

	Describe("Create", func() {
		var tx *MovieTx

		BeforeEach(func() {
			tx = db.Begin().(*MovieTx)
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
			tx = db.Begin().(*MovieTx)
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
			tx = db.Begin().(*MovieTx)
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

		BeforeEach(func() {
			tx = db.Begin().(*MovieTx)
		})

		It("returns a list of movies", func() {
			tx.Create(&critic.Movie{
				Title: "Jaws",
				Year:  "1975",
			})

			tx.Create(&critic.Movie{
				Title: "Citizen Kane",
				Year:  "1941",
			})

			tx.Commit()

			movies, err := db.Query(
				critic.MovieQueryParams{},
				critic.MovieListOptions{},
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(movies).ToNot(BeNil())
			Expect(len(movies)).To(Equal(2))
		})

		Context("when a before ID is provided", func() {
			It("only returns movies before the given ID", func() {
				jaws := &critic.Movie{
					Title: "Jaws",
					Year:  "1975",
				}
				tx.Create(jaws)

				citizenKane := &critic.Movie{
					Title: "Citizen Kane",
					Year:  "1941",
				}
				tx.Create(citizenKane)

				movies, err := db.Query(
					critic.MovieQueryParams{
						BeforeID: citizenKane.ID,
					},
					critic.MovieListOptions{},
				)

				Expect(err).ToNot(HaveOccurred())
				Expect(movies).ToNot(BeNil())
				Expect(movies).To(HaveLen(1))
				Expect(movies[0].ID).To(Equal(jaws.ID))
			})
		})

		Context("when a limit is provided", func() {
			It("limits the number of movies returned", func() {
				jaws := &critic.Movie{
					Title: "Jaws",
					Year:  "1975",
				}
				tx.Create(jaws)

				citizenKane := &critic.Movie{
					Title: "Citizen Kane",
					Year:  "1941",
				}
				tx.Create(citizenKane)

				movies, err := db.Query(
					critic.MovieQueryParams{},
					critic.MovieListOptions{
						Limit: 1,
					},
				)

				Expect(err).ToNot(HaveOccurred())
				Expect(movies).ToNot(BeNil())
				Expect(movies).To(HaveLen(1))
				Expect(movies[0].ID).To(Equal(citizenKane.ID))
			})
		})
	})

	Describe("Get", func() {
		var tx *MovieTx

		BeforeEach(func() {
			tx = db.Begin().(*MovieTx)
		})

		It("returns a movie", func() {
			jaws := &critic.Movie{
				Title: "Jaws",
				Year:  "1975",
			}

			tx.Create(jaws)
			tx.Commit()

			movie, err := db.Get(jaws.ID)
			Expect(err).ToNot(HaveOccurred())
			Expect(movie).ToNot(BeNil())
			Expect(movie.ID).To(Equal(jaws.ID))
		})

		Context("when the movie does not exist", func() {
			It("returns a not found error", func() {
				tx.Create(&critic.Movie{
					Title: "Jaws",
					Year:  "1975",
				})
				tx.Commit()

				_, err := db.Get(0)
				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(ErrNotFound))
			})
		})
	})
})
