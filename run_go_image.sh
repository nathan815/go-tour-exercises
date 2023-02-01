#!/bin/sh

GO_PROGRAM="$1"
OUTPUT_FILE="$2"

echo "go run $GO_PROGRAM"

go run $GO_PROGRAM | sed -e "s/^IMAGE://" | base64 -d > $OUTPUT_FILE

echo "Saved image to $OUTPUT_FILE"
