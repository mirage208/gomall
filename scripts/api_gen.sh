#!/bin/bash

# Get the script directory and project root
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
if [[ "$(basename "$SCRIPT_DIR")" == "scripts" ]]; then
    PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
else
    PROJECT_ROOT="$SCRIPT_DIR"
fi

echo "Project root: $PROJECT_ROOT"

# Generate API code
for app in user product order payment cart; do
    api_file="$PROJECT_ROOT/app/$app/specs/api/${app}.api"
    if [[ -f "$api_file" ]]; then
        echo "Generating API code for: $api_file"

        goctl api go --api "$api_file" \
            --dir "$PROJECT_ROOT/app/$app/api" \
            --style gozero
    else
        echo "API file not found: $api_file"
    fi
done
