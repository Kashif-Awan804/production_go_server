build:
	go build -o bin/api ./cmd/api

run:
	go run ./cmd/api

start:
	./bin/api

clean:
	rm -f bin/api