#!/bin/bash

# Get the script directory and project root
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
if [[ "$(basename "$SCRIPT_DIR")" == "scripts" ]]; then
    PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
else
    PROJECT_ROOT="$SCRIPT_DIR"
fi

echo "Project root: $PROJECT_ROOT"

# Generate Dockerfile for each service
for app in user product order payment cart; do
    # Check if API service exists
    api_dir="$PROJECT_ROOT/app/$app/api"
    rpc_dir="$PROJECT_ROOT/app/$app/rpc"

    # Generate Dockerfile for API service
    if [[ -d "$api_dir" ]]; then
        echo "Generating Dockerfile for API service: $app"

        cd "$api_dir"
        goctl docker \
            --go "${app}.go" \
            --exe "${app}_api"

        if [[ $? -eq 0 ]]; then
            echo "✓ Dockerfile generated successfully for $app API service"
        else
            echo "✗ Failed to generate Dockerfile for $app API service"
        fi
        cd "$PROJECT_ROOT"
    else
        echo "API directory not found: $api_dir"
    fi

    # Generate Dockerfile for RPC service
    if [[ -d "$rpc_dir" ]]; then
        echo "Generating Dockerfile for RPC service: $app"

        cd "$rpc_dir"
        goctl docker \
            --go "${app}.go" \
            --exe "${app}_rpc"

        if [[ $? -eq 0 ]]; then
            echo "✓ Dockerfile generated successfully for $app RPC service"
        else
            echo "✗ Failed to generate Dockerfile for $app RPC service"
        fi
        cd "$PROJECT_ROOT"
    else
        echo "RPC directory not found: $rpc_dir"
    fi

    echo "---"
done

echo "Dockerfile generation completed!"
