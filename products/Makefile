.PHONY: dev
dev: docker-up goose-up
	@air

.PHONY: docker-up
docker-up:
	@docker compose up -d

.PHONY: docker-down
docker-down:
	@docker compose down

.PHONY: goose-status
goose-status:
	@make docker-up
	@go build -o ./tmp/goose ./cmd/goose/main.go
	@./tmp/goose status
	@rm -rf ./tmp

.PHONY: goose-up
goose-up:
	@make docker-up
	@go build -o ./tmp/goose ./cmd/goose/main.go
	@./tmp/goose up
	@rm -rf ./tmp

.PHONY: goose-down
goose-down:
	@make docker-up
	@go build -o ./tmp/goose ./cmd/goose/main.go
	@./tmp/goose down
	@rm -rf ./tmp
