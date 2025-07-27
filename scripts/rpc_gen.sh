#!/bin/bash

# Get the script directory and project root
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
if [[ "$(basename "$SCRIPT_DIR")" == "scripts" ]]; then
    PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
else
    PROJECT_ROOT="$SCRIPT_DIR"
fi

echo "Project root: $PROJECT_ROOT"

# Generate gRPC code
for app in user product order payment cart; do
    proto_file="$PROJECT_ROOT/app/$app/specs/rpc/${app}.proto"
    if [[ -f "$proto_file" ]]; then
        echo "Generating gRPC code for: $proto_file"

        goctl rpc protoc "$proto_file" \
            --proto_path="$PROJECT_ROOT/app/$app/specs/rpc" \
            --go_out="$PROJECT_ROOT/app/$app/rpc/pb" \
            --go-grpc_out="$PROJECT_ROOT/app/$app/rpc/pb" \
            --zrpc_out="$PROJECT_ROOT/app/$app/rpc" \
            --style gozero
    else
        echo "Proto file not found: $proto_file"
    fi
done
