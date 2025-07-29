#!/usr/bin/env bash

# Get the script directory and project root
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
if [[ "$(basename "$SCRIPT_DIR")" == "build" ]]; then
    PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
else
    PROJECT_ROOT="$SCRIPT_DIR"
fi
echo "Project root: $PROJECT_ROOT"
cd "$PROJECT_ROOT"

# Build RPC images
for app in user payment product order cart; do

    api_dockerfile="$PROJECT_ROOT/build/docker/${app}_api.Dockerfile"
    rpc_dockerfile="$PROJECT_ROOT/build/docker/${app}_rpc.Dockerfile"

    if [[ -f "$api_dockerfile" ]]; then
        echo "Building API image for ${app}"
        docker build -f "$api_dockerfile" -t gomall/${app}-api:latest .
        if [[ $? -ne 0 ]]; then
            echo "Failed to build ${app}-api image"
            exit 1
        fi
    else
        echo "Dockerfile for ${app} API not found. Skipping build."
    fi

    if [[ -f "$rpc_dockerfile" ]]; then
        echo "Building RPC image for ${app}"
        docker build -f "$rpc_dockerfile" -t gomall/${app}-rpc:latest .
        if [[ $? -ne 0 ]]; then
            echo "Failed to build ${app}-rpc image"
            exit 1
        fi
    else
        echo "Dockerfile for ${app} RPC not found. Skipping build."
        echo $rpc_dockerfile
    fi

done
