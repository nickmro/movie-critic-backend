# Movie Critic Backend

Movie Critic Backend is a server for a movie critic application. It is written with Go with a PostgreSQL database.

## About Movie Critic

Movie critic is an application that allows users to concisely review and rate movies. Users may follow other users.

## Structure

```
.
+-- cmd          // application commands
+-- critic       // models and components
    +-- postgres // postgreSQL database functions
```

## Setup

### Go

Install Go by following the instructions here: [https://golang.org/doc/install](https://golang.org/doc/install)

### PostgreSQL

Install and run postgreSQL:

```bash
brew install postgresql@11
brew services start postgresql@11
```

## Test

To run the test suite:

```bash
make test DB_NAME=movie_critic_test
```
