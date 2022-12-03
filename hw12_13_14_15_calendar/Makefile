BIN := "./bin/calendar"
DOCKER_IMG="calendar:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

SWAGGER_UI_VERSION:=v4.15.5

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/calendar

run: build
	$(BIN) -config ./configs/calendar.yaml

build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG) \
		-f build/Dockerfile .

run-img: build-img
	docker run $(DOCKER_IMG)

version: build
	$(BIN) version

test:
	go test -race ./internal/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.50.1

lint: install-lint-deps
	golangci-lint run ./...

install-migrate-deps:
	go install github.com/pressly/goose/v3/cmd/goose@latest

migrate: install-lint-deps
	goose -dir ./migrations up

install-generate-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

generate-proto: install-generate-deps
	mkdir -p pkg/calendarpb
	mkdir -p third_party/OpenAPI
	protoc api/calendar/calendar.proto \
		-I api \
		-I third_party/googleapis \
		-I third_party/grpc-gateway \
		--go_out pkg/calendarpb \
		--go-grpc_out pkg/calendarpb \
		--grpc-gateway_out pkg/calendarpb \
		--grpc-gateway_opt paths=import \
		--grpc-gateway_opt logtostderr=true \
		--openapiv2_out third_party/OpenAPI \
		--openapiv2_opt disable_default_responses=true \
		--openapiv2_opt logtostderr=true
	mv pkg/calendarpb/github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/pkg/calendarpb/* pkg/calendarpb/
	rm -rf pkg/calendarpb/github.com/


generate-swagger-ui:
	SWAGGER_UI_VERSION=$(SWAGGER_UI_VERSION) ./scripts/generate-swagger-ui.sh

generate: generate-proto generate-swagger-ui

clean:
	rm -rf ./bin
	rm -rf ./pkg/calendarpb
	rm -rf ./third_party/OpenAPI

.PHONY: build run build-img run-img version test lint migrate generate clean
