.PHONY: build
build:
	go build .

.PHONY: test
test:
	$(eval RICHIGO := $(shell which richgo > /dev/null; echo $$?))
	@if [ $(RICHIGO) = 0 ]; then richgo test -v ./...; else go test -v ./...; fi

.PHONY: coverage
coverage:
	go test -coverprofile coverage.out -covermode atomic .
	go tool cover -html=coverage.out

.PHONY: lint
lint:
	golint -set_exit_status ./...
