#!/bin/bash
# This script is used to run a Docker container for a Go API application.
# It mounts the source code directory into the container and exposes port 8080.

# Linux command to run the Docker container
docker run -it --rm -v "$(pwd)/src:/app/go-api/src" -v "$(pwd)/database:/app/go-api/database" -v "$(pwd)/test:/app/go-api/test" -p 8080:8080 go-api

# PowerShell command to run the Docker container
# docker run -it --rm -v "$(PWD)\src:/app/go-api/src" -v "$(PWD)/database:/app/go-api/database" -v "$(pwd)/test:/app/go-api/test" -p 8080:8080 go-api

# CMD command to run the Docker container
# docker run -it --rm -v "%cd%/src:/app/go-api/src" -v "%cd%/database:/app/go-api/database" -v "$(pwd)/test:/app/go-api/test" -p 8080:8080 go-api
