all: clean build

.PHONY: run
run:
	@go run cmd/main.go examples/get.yml

.PHONY: clean
clean:
	rm -rf dist

.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	go build -i -v -ldflags="-X main.version=$(git describe --always --dirty)" -o dist/relay cmd/main.go

release:
	@./scripts/release.sh
