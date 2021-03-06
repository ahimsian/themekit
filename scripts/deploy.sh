#!/bin/bash
set -e

main() {
  if [ -z "$DEPLOY_VERSION" ]; then
    echo "---------- No tag to deploy ---------- "
    return
  fi

  echo "---------- Building $DEPLOY_VERSION ----------"
  make all

  echo "---------- Deploying $DEPLOY_VERSION ----------"
  go build -o ./build/release ./cmd/tkrelease
  if [ -z "$FORCE_DEPLOY" ]; then
    ./build/release -k="$AWS_KEY" -s="$AWS_SECRET" $DEPLOY_VERSION
  else
    ./build/release -f -k="$AWS_KEY" -s="$AWS_SECRET" $DEPLOY_VERSION
  fi

  echo "---------- Generating Shasum for homebrew ----------"
  make gen_sha
}

main "$@"
