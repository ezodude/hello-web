SHELL=/bin/bash

# Clean, build image, and run container locally.
.PHONY: up
up: clean_build_image up_images clean

# Stop container and clean locally.
.PHONY: down
down: down_images clean

.PHONY: up_images
up_images:
	docker-compose -f ./docker-compose.yml up -d

.PHONY: down_images
down_images:
	docker-compose -f ./docker-compose.yml down -v --remove-orphans

# Build GO binaries for dev. Assumes local GO env is setup.

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hello-web .

.PHONY: clean
clean:
	@if [ -e "hello-web" ]; \
		then \
			rm hello-web; \
	fi

# Build hello_web docker image for dev, i.e local machine.

.PHONY: build_image
build_image:
	docker-compose -f ./docker-compose.yml build --force-rm hello-web

.PHONY: clean_build_image
clean_build_image: clean build build_image
	