include .envrc

run:
	go run ./cmd/api

db:
	docker-compose up

up:
	migrate -path=./migrations -database="postgresql://postgres:postgres@localhost:5432/greenlight?sslmode=disable" up


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #
audit: vendor
	@echo 'Formatting code'
	go fmt ./...
	@echo 'Vetting code'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests'
	go test -race -vet=off ./...

vendor:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

# ==================================================================================== #
# BUILD 
# ==================================================================================== #
build:
	@echo 'Building..'
	go build -ldflags='-s' -o=./bin/api ./cmd/api
	GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=./bin/linux_amd64/api ./cmd/api
