# tests the repository packages
test:
	-dropdb $(DB_NAME)
	createdb $(DB_NAME)
	migrate -path=./critic/postgres/migrations -database="postgres://localhost/$(DB_NAME)?sslmode=disable" up
	DB_URL=postgres://localhost/$(DB_NAME)?sslmode=disable go test ./...

# makes a new migration file
migration:
	migrate create -ext "sql" -dir "./critic/postgres/migrations" $(NAME) 

# runs the migration files
migrate:
	migrate -path=./critic/postgres/migrations -database="$(DB_URL)" up
