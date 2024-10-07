include .env
LOCAL_BIN=$(CURDIR)/bin

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.3

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

install-migrate:
	GOBIN=$(LOCAL_BIN) go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate-create:
	migrate create -ext sql -dir migrations 'migrate_name'

migrate-up:
	migrate -path migrations -database '$(DB_URL)' up

migrate-down:
	migrate -path migrations -database '$(DB_URL)' down
