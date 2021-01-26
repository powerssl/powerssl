#!/bin/bash

set -euxo pipefail

gobin -m -run github.com/golang-migrate/migrate/v4/cmd/migrate create -ext sql -dir db/migrations $@
