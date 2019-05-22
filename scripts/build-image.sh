#!/bin/sh

set -eux

BUILD_ARG=
CIRCLECI=${CIRCLECI:-false}
DIR=cli
DOCKERFILE=Dockerfile

if [ "${CIRCLECI}" = true ]; then
  DOCKERFILE=Dockerfile.circleci
fi

case $COMPONENT in
  powerssl-apiserver|powerssl-controller|powerssl-signer)      
    BUILD_ARG="COMPONENT=${COMPONENT}"
    DIR=grpc-server
    ;;
  powerssl-auth|powerssl-webapp)      
    BUILD_ARG="COMPONENT=${COMPONENT}"
    DIR=web-server
    ;;
  powerssl-integration-*)      
    BUILD_ARG="INTEGRATION=$(echo "${COMPONENT}" | sed 's/powerssl-integration-//')"
    DIR=integration
    ;;
esac

docker build -f "build/docker/${DIR}/${DOCKERFILE}" -t "${TAG}" --build-arg="${BUILD_ARG}" .
