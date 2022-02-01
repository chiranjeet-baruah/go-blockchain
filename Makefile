install:
	go mod tidy 

test:
	go vet ./...

binary-build:
	echo "compiling with `go version`"
	CGO_ENABLED=0 GOOS=linux go build -o ./go-blockchain
	echo "built ./go-blockchain"

build: test binary-build