APP_NAME=myapi
IMAGE_NAME=go-gin-api

# Build Go binary
build:
	go build -o $(APP_NAME)

run: build
	./$(APP_NAME)
# Build Docker image
docker-build:
	docker build -t $(IMAGE_NAME) .

# Run Docker container
docker-run: docker-build
	docker run -p 8080:8080 --rm --name $(IMAGE_NAME)-container $(IMAGE_NAME)

# Clean up binary
clean:
	rm -f $(APP_NAME)
