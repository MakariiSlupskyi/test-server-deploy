.PHONY: dev
dev:
	@echo "Running dev server with Air..."
	air -c .air.toml

.PHONY: run
run:
	@echo "Running Go app..."
	@go run cmd/api/main.go