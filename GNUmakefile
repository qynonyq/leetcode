TEST?=$$(go list ./...)
GOFMT_FILES?=$$(find . -name '*.go')

default: build

build: fmtcheck
	go install

test: fmtcheck
	go test $(TEST) || exit 1
	echo $(TEST) | \
		xargs -tn4 go test $(TESTARGS)-race -cover -timeout=30s -parallel=16

vet:
	@echo "go vet ."
	@go vet $$(go list ./...) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

errcheck:
	@sh -c "'$(CURDIR)/scripts/errcheck.sh'"

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.1
	golangci-lint run ./...

.PHONY: build test vet fmt fmtcheck errcheck lint
