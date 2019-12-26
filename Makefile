.PHONY: build clean deploy

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/books/endpoints/create modules/books/endpoints/create.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/books/endpoints/read modules/books/endpoints/read.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/books/endpoints/update modules/books/endpoints/update.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/books/endpoints/delete modules/books/endpoints/delete.go
clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
