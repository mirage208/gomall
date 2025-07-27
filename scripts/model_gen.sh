#!/bin/bash

# Get the script directory and project root
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
if [[ "$(basename "$SCRIPT_DIR")" == "scripts" ]]; then
    PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
else
    PROJECT_ROOT="$SCRIPT_DIR"
fi

echo "Project root: $PROJECT_ROOT"

# Generate models code
for app in user product order payment cart; do
    for sql_file in "$PROJECT_ROOT"/app/$app/specs/sql/*.sql; do
        if [[ -f "$sql_file" ]]; then
            echo "Generating model for: $sql_file"

            goctl model mysql ddl --src "$sql_file" \
                --dir "$PROJECT_ROOT/app/$app/model" \
                --style gozero --strict
        else
            echo "SQL file not found: $sql_file"
        fi
    done
done
