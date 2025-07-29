#!/bin/bash

# Get the script directory and project root
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
if [[ "$(basename "$SCRIPT_DIR")" == "scripts" ]]; then
    PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
else
    PROJECT_ROOT="$SCRIPT_DIR"
fi

echo "Project root: $PROJECT_ROOT"
echo "Updating port configurations for all services..."

# Update configurations for each service
for app in user product order payment cart; do
    # Define ports for each service
    case $app in
        "user")
            api_port=8001
            rpc_port=9001
            api_metrics_port=6001
            rpc_metrics_port=6011
            ;;
        "product")
            api_port=8002
            rpc_port=9002
            api_metrics_port=6002
            rpc_metrics_port=6012
            ;;
        "order")
            api_port=8003
            rpc_port=9003
            api_metrics_port=6003
            rpc_metrics_port=6013
            ;;
        "payment")
            api_port=8004
            rpc_port=9004
            api_metrics_port=6004
            rpc_metrics_port=6014
            ;;
        "cart")
            api_port=8005
            rpc_port=9005
            api_metrics_port=6005
            rpc_metrics_port=6015
            ;;
    esac
    
    echo "---"
    echo "Updating $app service ports:"
    echo "  API: $api_port (metrics: $api_metrics_port)"
    echo "  RPC: $rpc_port (metrics: $rpc_metrics_port)"
    
    # Update API configuration
    api_config="$PROJECT_ROOT/app/$app/api/etc/${app}.yaml"
    if [[ -f "$api_config" ]]; then
        # Update API port
        sed -i '' "s/^Port: .*/Port: $api_port/" "$api_config"
        
        # Update or add DevServer configuration
        if grep -q "DevServer:" "$api_config"; then
            # Update existing DevServer port
            sed -i '' "/DevServer:/,/^[^ ]/ s/Port: .*/Port: $api_metrics_port/" "$api_config"
        else
            # Add DevServer configuration
            echo "" >> "$api_config"
            echo "# API Metrics配置" >> "$api_config"
            echo "DevServer:" >> "$api_config"
            echo "  Enabled: true" >> "$api_config"
            echo "  Port: $api_metrics_port" >> "$api_config"
            echo "  MetricsPath: /metrics" >> "$api_config"
            echo "  EnableMetrics: true" >> "$api_config"
        fi
        
        echo "  ✓ Updated API config: $api_config"
    else
        echo "  ✗ API config not found: $api_config"
    fi
    
    # Update RPC configuration
    rpc_config="$PROJECT_ROOT/app/$app/rpc/etc/${app}.yaml"
    if [[ -f "$rpc_config" ]]; then
        # Update RPC port
        sed -i '' "s/^ListenOn: .*/ListenOn: 0.0.0.0:$rpc_port/" "$rpc_config"
        
        # Update or add DevServer configuration for RPC
        if grep -q "DevServer:" "$rpc_config"; then
            # Update existing DevServer port
            sed -i '' "/DevServer:/,/^[^ ]/ s/Port: .*/Port: $rpc_metrics_port/" "$rpc_config"
        else
            # Add DevServer configuration
            echo "" >> "$rpc_config"
            echo "# RPC Metrics配置" >> "$rpc_config"
            echo "DevServer:" >> "$rpc_config"
            echo "  Enabled: true" >> "$rpc_config"
            echo "  Port: $rpc_metrics_port" >> "$rpc_config"
            echo "  MetricsPath: /metrics" >> "$rpc_config"
            echo "  EnableMetrics: true" >> "$rpc_config"
        fi
        
        echo "  ✓ Updated RPC config: $rpc_config"
    else
        echo "  ✗ RPC config not found: $rpc_config"
    fi
done

echo "---"
echo "Port configuration update completed!"
echo ""
echo "Port Summary:"
echo "============="
echo "Service   | API  | RPC  | API Metrics | RPC Metrics"
echo "----------|------|------|-------------|------------"
echo "user      | 8001 | 9001 | 6001        | 6011"
echo "product   | 8002 | 9002 | 6002        | 6012"
echo "order     | 8003 | 9003 | 6003        | 6013"
echo "payment   | 8004 | 9004 | 6004        | 6014"
echo "cart      | 8005 | 9005 | 6005        | 6015"
