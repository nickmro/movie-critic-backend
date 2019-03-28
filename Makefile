test:
	-dropdb $(DB_NAME)
	createdb $(DB_NAME)
	DB_URL=postgres://localhost/$(DB_NAME)?sslmode=disable go test ./...
