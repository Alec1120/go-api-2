#!/bin/bash
# This script is used to build a Docker image for a Go API application.
# It uses the Dockerfile located in the docker directory and tags the image as go-api.

docker build -f docker/Dockerfile -t go-api .