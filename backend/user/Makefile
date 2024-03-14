.PHONY: install test build serve clean pack deploy ship
TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

export TAG

install:
	go get .

test: install
	go test ./...

build: install
	go build -ldflags "-X main.version=$(TAG)" -o ./bin/go-user-service .

serve: build
	./bin/go-user-service serve grpc

clean:
	rm -f ./bin/go-user-service

dev:
	GOOS=linux make build
	docker build -t n3kizz/go-user-service .
	docker push n3kizz/go-user-service