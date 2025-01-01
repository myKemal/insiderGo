APP_NAME := insider-go
DOCKER_IMAGE := $(APP_NAME):latest
CONTAINER_NAME := insiderGo
PLATFORMS := linux/amd64,linux/arm64
GOOS := darwin
GOARCH := amd64
PORT := 8090

.PHONY: all
all: build

.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(APP_NAME)

.PHONY: build
build:
	@echo "Building the Go binary for platform: $(GOOS)/$(GOARCH)..."
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(APP_NAME)

.PHONY: run
run:
	@echo "Running the application locally..."
	@PORT=$(PORT) TEMP_STORAGE=MEMORY ./$(APP_NAME)

.PHONY: test
test:
	@echo "Running tests..."
	go test ./... -v

.PHONY: docker-build
docker-build:
	@echo "Building Docker image for current architecture..."
	docker build -t $(DOCKER_IMAGE) --build-arg GOOS=$(GOOS) --build-arg GOARCH=$(GOARCH) .

.PHONY: docker-run
docker-run:
	@echo "Running Docker container..."
	@if [ $$(docker ps -q -f name=$(CONTAINER_NAME)) ]; then \
		echo "Container $(CONTAINER_NAME) is already running."; \
	else \
		docker run --name $(CONTAINER_NAME) -p $(PORT):$(PORT) -d $(DOCKER_IMAGE); \
		echo "Container $(CONTAINER_NAME) started."; \
	fi

.PHONY: docker-stop
docker-stop:
	@echo "Stopping Docker container..."
	@if [ $$(docker ps -q -f name=$(CONTAINER_NAME)) ]; then \
		docker stop $(CONTAINER_NAME); \
		docker rm $(CONTAINER_NAME); \
		echo "Container $(CONTAINER_NAME) stopped and removed."; \
	else \
		echo "No running container named $(CONTAINER_NAME)."; \
	fi


.PHONY: compose-up
compose-up:
	@echo "Starting the application with Docker Compose..."
	docker-compose up --build

.PHONY: compose-up-in-memory
compose-up-in-memory:
	@echo "Starting the application with Docker Compose to use InMemory..."
	docker-compose -f docker-compose.yml -f docker-compose.override.yml up --build

.PHONY: compose-down
compose-down:
	@echo "Stopping the application with Docker Compose..."
	docker-compose down

.PHONY: lint
lint:
	@echo "Linting the code..."
	@which golangci-lint > /dev/null || (echo "Installing golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	golangci-lint run


.PHONY: fmt
fmt:
	@echo "Formatting the code..."
	go fmt ./...

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build          Build the Go binary"
	@echo "  run            Run the application locally"
	@echo "  test           Run tests"
	@echo "  docker-build   Build Docker image for the current architecture"
	@echo "  docker-run     Run Docker container for the current architecture"
	@echo "  docker-stop    Stop Docker container for the current architecture"
	@echo "  compose-up     Start the application with Docker Compose"
	@echo "  compose-down   Stop the application with Docker Compose"
	@echo "  lint           Lint the code using golangci-lint"
	@echo "  fmt            Format the code"
	@echo "  clean          Clean build artifacts"
