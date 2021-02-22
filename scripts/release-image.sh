#!/bin/bash

set -euxo pipefail

docker push "$TAG"
