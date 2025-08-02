.PHONY: dev-build

dev-build:
	@echo "Creating volume directories if they do not exist..."
	@mkdir -p volumes/quantdrey-mongo volumes/quantdrey-redis volumes/quantdrey-nats volumes/quantdrey-postgres
	@echo "Starting docker compose..."
	docker compose -f docker-compose-dev.yml up -d
