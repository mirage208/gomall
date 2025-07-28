#!/usr/bin/env bash

for app in user payment product order cart; do
    # Tag RPC image
    docker tag gomall/${app}_rpc:latest crpi-uuzhidt8pbh6o5xf.cn-shanghai.personal.cr.aliyuncs.com/mirage208/mall:${app}_rpc
done
