BINARY=tender

build: dep test
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' ./...
	CGO_ENABLED=0 GOOS=linux go build -o ${BINARY} -a -ldflags '-extldflags "-static"' ./cmd/${BINARY}

dep:
	dep ensure

test: dep
	go test ./...
