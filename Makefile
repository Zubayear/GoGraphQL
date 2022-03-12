GOPATH:=$(shell go env GOPATH)

.PHONY: gql
gql:
	@go run github.com/99designs/gqlgen generate

.PHONY: run
run:
	@go run server.go

.PHONY: migrate
migrate:
	@go generate ./ent

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: di
di:
	@wire di/wire.go