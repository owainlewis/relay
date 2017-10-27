.PHONY: clean
clean:
	rm -rf dist

.PHONY: build
build:
	go build -i -v -ldflags="-X main.version=$(git describe --always --long --dirty)" -o dist/relay main.go

release:
	@./scripts/release.sh
