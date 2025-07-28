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
    # Define ports for each service
    case $app in
    "user")
        api_port=8001
        rpc_port=9001
        ;;
    "product")
        api_port=8002
        rpc_port=9002
        ;;
    "order")
        api_port=8003
        rpc_port=9003
        ;;
    "payment")
        api_port=8004
        rpc_port=9004
        ;;
    "cart")
        api_port=8005
        rpc_port=9005
        ;;
    esac

    # Check if API service exists
    api_dir="$PROJECT_ROOT/app/$app/api"
    rpc_dir="$PROJECT_ROOT/app/$app/rpc"

    # Generate Dockerfile for API service
    if [[ -d "$api_dir" ]]; then
        echo "Generating Dockerfile for API service: $app (port: $api_port)"

        cd "$api_dir"

        # Delete existing Dockerfile
        if [[ -f "Dockerfile" ]]; then
            echo "Deleting existing Dockerfile in $api_dir"
            rm -f Dockerfile
        fi

        goctl docker \
            --go "${app}.go" \
            --exe "${app}_api" \
            --port $api_port

        if [[ $? -eq 0 ]]; then
            echo "✓ Dockerfile generated successfully for $app API service"
            mv Dockerfile "$PROJECT_ROOT/build/docker/${app}_api.Dockerfile"
        else
            echo "✗ Failed to generate Dockerfile for $app API service"
        fi
        cd "$PROJECT_ROOT"
    else
        echo "API directory not found: $api_dir"
    fi

    # Generate Dockerfile for RPC service
    if [[ -d "$rpc_dir" ]]; then
        echo "Generating Dockerfile for RPC service: $app (port: $rpc_port)"

        cd "$rpc_dir"

        # Delete existing Dockerfile
        if [[ -f "Dockerfile" ]]; then
            echo "Deleting existing Dockerfile in $rpc_dir"
            rm -f Dockerfile
        fi

        goctl docker \
            --go "${app}.go" \
            --exe "${app}_rpc" \
            --port $rpc_port

        if [[ $? -eq 0 ]]; then
            echo "✓ Dockerfile generated successfully for $app RPC service"
            mv Dockerfile "$PROJECT_ROOT/build/docker/${app}_rpc.Dockerfile"
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
