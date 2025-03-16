BINARY_NAME=rating_service

.PHONY: build

PACKAGE=github.com/guneyin/rate-my-service/services/rating

init: clean download tidy vet

clean:
	go clean
	rm -f ${BINARY_NAME}

download:
	go mod download

tidy:
	go mod tidy

vet:
	go vet ./...

run:
	go run .

lint:
	golangci-lint run --fix

test:
	go test ./...

coverage:
	go test ./... -covermode=count -coverprofile=coverage.out fmt
	go tool cover -func=coverage.out -o=coverage.out
	gobadge -filename=coverage.out