#!/bin/bash


# Get the script directory and project root
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
if [[ "$(basename "$SCRIPT_DIR")" == "scripts" ]]; then
    PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
else
    PROJECT_ROOT="$SCRIPT_DIR"
fi
cd "$PROJECT_ROOT"
echo "Project root: $PROJECT_ROOT"

# Generate Kubernetes manifests for each service
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

    rpc_yml="$PROJECT_ROOT/build/ci/${app}-rpc.yaml"
    api_yml="$PROJECT_ROOT/build/ci/${app}-api.yaml"

    # Delete existing Kubernetes manifests
    if [[ "$api_yml" ]]; then
        echo "Deleting existing Kubernetes manifests in $api_dir"
        rm -f "$api_yml"
    fi
    if [[ -f "$rpc_yml" ]]; then
        echo "Deleting existing Kubernetes manifests in $rpc_dir"
        rm -f "$rpc_yml"
    fi

    # Generate Kubernetes manifests for API service
    echo "Generating Kubernetes manifests for API service: $app (port: $api_port)"
    goctl kube deploy \
        -replicas 1 -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 \
        -name ${app}-api -namespace go-mall -image crpi-uuzhidt8pbh6o5xf.cn-shanghai.personal.cr.aliyuncs.com/mirage208/mall:${app}-api \
        -o "$api_yml" -port ${api_port} --serviceAccount find-endpoints

    if [[ $? -eq 0 ]]; then
        echo "✓ Kubernetes manifests generated successfully for $app API service"
    else
        echo "✗ Failed to generate Kubernetes manifests for $app API service"
    fi

    # Generate Kubernetes manifests for RPC service
    goctl kube deploy \
        -replicas 1 -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 \
        -name ${app}-rpc -namespace go-mall -image crpi-uuzhidt8pbh6o5xf.cn-shanghai.personal.cr.aliyuncs.com/mirage208/mall:${app}-rpc \
        -o "$rpc_yml" -port ${rpc_port} --serviceAccount find-endpoints

    if [[ $? -eq 0 ]]; then
        echo "✓ Kubernetes manifests generated successfully for $app RPC service"
    else
        echo "✗ Failed to generate Kubernetes manifests for $app RPC service"
    fi
    cd "$PROJECT_ROOT"

    echo "---"
done

echo "Kubernetes manifests generation completed!"
