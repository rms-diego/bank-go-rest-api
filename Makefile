build:
	go build -o bin/server ./cmd/server

start:
	./bin/server

dev:
	air

clean:
	go clean
	rm -f bin/server tmp/server
