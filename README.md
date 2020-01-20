# Hello web

This repo provides a sample Hello World `Go` web server designed to run as a Docker container.

The project uses `Go modules`.

## Dockerfile

`FROM scratch`

Included is a basic `Dockerfile` that uses Docker’s reserved, minimal image, `scratch` as a starting point.

`ADD ca-certificates.crt /etc/ssl/certs/`

Then, we add the public root certificates.

`ADD hello-web /`

Then, we add our compiled `Go` binary. Use the `Makefile` to compile this.

`ENTRYPOINT ["/hello-web"]`

Then, configure the container to run as an executable (running the compiled `Go` binary).

## docker-compose

A basic `docker-compose` file to build a `ezodude/hello-web` image and run it as a container.

## Makefile

Contains commands and instructions to,

- Build a `Go` linux binary.
- Clean up older binaries.
- Build docker images using docker-compose.
- Run docker images using docker-compose.
- Stop docker images using docker-compose.