.PHONY: install test build serve clean pack deploy ship
TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

export TAG

install:
	go get .

test: install
	go test ./...

build: install
	go build -ldflags "-X main.version=$(TAG)" -o ./bin/go-room-service .

serve: build
	./bin/go-room-service serve grpc

clean:
	rm -f ./bin/go-room-service

dev:
	GOOS=linux make build
	docker build -t n3kizz/go-room-service .
	docker push n3kizz/go-room-service