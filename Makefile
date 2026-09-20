build:
	go build -o bin/api ./cmd/api

run:
	go run ./cmd/api

start:
	./bin/api

migrate:
	go run ./cmd/migrate up

migrate-down:
	go run ./cmd/migrate down

clean:
	rm -f bin/api