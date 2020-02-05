.PHONY: build

build:
	go build .

test:
	$(eval RICHIGO := $(shell which richgo > /dev/null; echo $$?))
	@if [ $(RICHIGO) = 0 ]; then richgo test -v ./...; else go test -v ./...; fi
