#!/usr/bin/env bash
set -e

mkdir -p engines

FORCE_FLAG=""
TARGET=""

for arg in "$@"; do
  if [ "$arg" = "clean" ]; then
    FORCE_FLAG="-a"
  else
    TARGET="$arg"
  fi
done

build_engine() {
  local dir="$1"
  local name
  name=$(basename "$dir")
  echo "Building $name${FORCE_FLAG:+ (clean)}..."
  go build $FORCE_FLAG -o "engines/$name" "./$dir"
}

if [ -n "$TARGET" ]; then
  if [ -d "cmd/$TARGET" ]; then
    build_engine "cmd/$TARGET"
  else
    echo "Error: Engine '$TARGET' not found in cmd/" >&2
    exit 1
  fi
else
  for dir in cmd/*/; do
    if [ -d "$dir" ]; then
      build_engine "$dir"
    fi
  done
fi

echo "Build complete. Binaries available in engines/"
