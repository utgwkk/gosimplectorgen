#!/bin/bash
set -euo pipefail

FILE_PATH=$(jq -r '.tool_input.file_path // .tool_response.filePath // empty')

if [[ -z "$FILE_PATH" ]]; then
  exit 0
fi

if [[ "$FILE_PATH" == *.go ]]; then
  gofmt -w "$FILE_PATH"
fi
