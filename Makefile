GO111MODULE=on
VERSION=$(shell git describe --tags --match 'v*' --always --abbrev=0)
SRC=.

LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION}"

build:
	@echo " > Building $@..."
	@mkdir -p "bin"
	@cd cmd/forum && GO111MODULE=$(GO111MODULE) go build -v ${LDFLAGS} -o forum
	@mv cmd/forum/forum bin/
	@echo " > Building $@...done."

test:
	@echo " > Testing $@..."
	@go test ./...
	@echo " > Testing $@...done."

generate:
	@echo " > Generating $@..."
	@go generate ./...
	@echo " > Generating $@...done."

new-migration:
	@echo " > New migration $@..."
	@migrate create -ext sql ${FILENAME}
	@mv *.up.sql internal/repository/postgres/migrations/
	@mv *.down.sql internal/repository/postgres/migrations/
	@echo " > New migration $@...done."

lint:
	@echo " > Linter $@..."
	@golangci-lint run --config .golangci.yml
	@echo " > Linter $@...done."

run:
	@echo " > Running $@..."
	@docker-compose -f docker/docker-compose.yml up -d
	@echo " > Running $@...done."

run-with-rebuild:
	@echo " > Running with rebuild $@..."
	@docker-compose -f docker/docker-compose.yml up --build -d
	@echo " > Running with rebuild $@...done."

stop:
	@echo " > Stopping $@..."
	@docker-compose -f docker/docker-compose.yml down
	@echo " > Stopping $@...done."
