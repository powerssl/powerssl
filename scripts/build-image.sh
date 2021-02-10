#!/bin/bash

set -euo pipefail

case "$COMPONENT" in
	powerssl-apiserver|powerssl-controller|powerssl-signer)
		BUILD_ARG="COMPONENT=$COMPONENT"
		DIR=grpc-server
		;;
	powerssl-auth|powerssl-grpcgateway|powerssl-temporalserver|powerssl-webapp)
		BUILD_ARG="COMPONENT=$COMPONENT"
		DIR=web-server
		;;
	powerssl-integration-*)
		BUILD_ARG="INTEGRATION=${COMPONENT/powerssl-integration-}"
		DIR=integration
		;;
  powerssl-worker)
    BUILD_ARG="COMPONENT=$COMPONENT"
		DIR=worker
		;;
	*)
		DIR=cli
		;;
esac

set -x

docker build -f "build/docker/$DIR/Dockerfile${CIRCLECI:+.circleci}" -t "$TAG" ${BUILD_ARG:+--build-arg=$BUILD_ARG} .
