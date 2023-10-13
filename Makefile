build:
	go build -o ./bin/blockchain1

run: build
	./bin/blockchain1

test:
	go test -v ./...
