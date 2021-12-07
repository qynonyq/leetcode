TEST?=$$(go list ./...)
GOFMT_FILES?=$$(find . -name '*.go')

default: fmt

fmt:
	gofmt -w -s $(GOFMT_FILES)

test:
	go test $(TEST) || exit 1
	echo $(TEST) | \
		xargs -tn4 go test $(TESTARGS)-race -cover -timeout=30s -parallel=16

lint:
	golangci-lint run ./...

.PHONY: fmt test lint
