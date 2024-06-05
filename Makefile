build:
	GO111MODULE=on go build -o bin/server ./cmd/server

dev:
	go run ./cmd/server

clean:
	go clean
	rm -f bin/server
