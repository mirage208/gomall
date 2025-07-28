#!/usr/bin/env bash

for app in user payment product order cart; do
    # Build RPC image
    docker build -f app/$app/rpc/Dockerfile -t gomall/${app}_rpc:latest .
done
