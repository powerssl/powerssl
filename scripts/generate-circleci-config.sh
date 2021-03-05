#!/bin/bash

set -eo pipefail

if ! command -v circleci-config-merge > /dev/null 2>&1; then
  echo "Please install circleci-config-merge" >&2
  echo "https://github.com/suzuki-shunsuke/circleci-config-merge#install" >&2
  exit 1
fi

cd "$(dirname "$0")/.."
cat .circleci/config-header-comment.txt > .circleci/config.yml
list_file=.circleci/config-list.txt
{ if [ -f "$list_file" ]; then grep -v -E "^#" "$list_file"; fi; git ls-files -o --exclude-standard | grep -E "^[^/]*/.circleci/.*\\.ya?ml$" || :; git ls-files | grep -E "^[^/]*/.circleci/.*\\.ya?ml$"; } | xargs circleci-config-merge merge >> .circleci/config.yml

if [[ -z "$CIRCLECI" ]]; then
  if command -v circleci > /dev/null 2>&1; then
    circleci config validate
  fi
fi
