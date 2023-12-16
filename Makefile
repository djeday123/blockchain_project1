build:
	go1.18 build -o ./bin/blockchain1

run: build
	./bin/blockchain1

test:
	go1.18 test  ./...

test2:
	go1.18 test  -v ./...
