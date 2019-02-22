#!/bin/sh

set -eux

docker build -f "build/docker/${COMPONENT}/Dockerfile" -t "powerssl/${COMPONENT}" .
