API_BIN=api_bin

## up: stops docker-compose (if running), builds all projects and starts docker compose
up: build_api
	@echo "Stopping docker images (if running...)"
	docker compose down
	@echo "Building (when required) and starting docker images..."
	docker compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker compose down
	@echo "Done!"

## test: run tests
test:
	@echo "Running tests..."
	cd ../api-service && go test ./cmd/api/
	@echo "Done!"

## build_api: builds the API service binary
build_api:
	@echo "Building API service binary..."
	cd ../api-service && env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${API_BIN} ./cmd/api
	@echo "Done!"
