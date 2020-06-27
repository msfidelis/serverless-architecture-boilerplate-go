.PHONY: build clean deploy

build:
	export GO111MODULE="on"
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/books/endpoints/create modules/books/endpoints/create.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/books/endpoints/read modules/books/endpoints/read.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/books/endpoints/detail modules/books/endpoints/detail.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/books/endpoints/update modules/books/endpoints/update.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/books/endpoints/delete modules/books/endpoints/delete.go

	env GOOS=linux go build -ldflags="-s -w" -o bin/books/functions/worker modules/books/functions/worker.go
clean:
	rm -rf ./bin
test:
	go test
deploy: clean build
	serverless deploy --verbose --force
