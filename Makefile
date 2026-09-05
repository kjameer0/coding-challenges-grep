run:
	go run .

watch:
	find . -name '*.go' | entr -r make run

build:
	go build .

test_all: build
	go test -tags=integration ./...

test_unit:
	go test ./...

test_integration:
	go test -v -tags=integration
