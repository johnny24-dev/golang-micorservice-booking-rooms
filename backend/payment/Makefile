.PHONY: install test build serve clean pack deploy ship
TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

export TAG

install:
	go get .

test: install
	go test ./...

build: install
	go build -ldflags "-X main.version=$(TAG)" -o ./bin/go-payment-service .

serve: build
	./bin/go-payment-service serve grpc

clean:
	rm -f ./bin/go-payment-service

dev:
	GOOS=linux make build
	docker build -t localhost:5000/go-payment-service .
	docker push localhost:5000/go-payment-service