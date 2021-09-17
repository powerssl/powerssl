#!/bin/bash

set -euxo pipefail

go run github.com/golang-migrate/migrate/v4/cmd/migrate@v4.14.1 create -ext sql -dir db/migrations "$@"
