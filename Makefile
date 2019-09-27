
## swagger: generate swagger scheme code
swagger:
	@swagger validate swagger.yml
	@swagger generate server
	@go mod tidy

## up: start server
up:
	@docker-compose up

## up: start server
build:
	@docker-compose up --build

## down: stop server
down:
	@docker-compose down

## test: run tests
test:
	@echo test

## help: Get makefile manual
help: Makefile
	@echo
	@echo Choose command to run:
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

.PHONY: swagger up down test help