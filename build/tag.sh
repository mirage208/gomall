#!/usr/bin/env bash

for app in user payment product order cart; do

    # Tag API image
    docker tag gomall/${app}-api:latest crpi-uuzhidt8pbh6o5xf.cn-shanghai.personal.cr.aliyuncs.com/mirage208/mall:${app}-api
    if [[ $? -ne 0 ]]; then
        echo "Failed to tag ${app}-api image"
        exit 1
    fi

    # Tag RPC image
    docker tag gomall/${app}-rpc:latest crpi-uuzhidt8pbh6o5xf.cn-shanghai.personal.cr.aliyuncs.com/mirage208/mall:${app}-rpc
    if [[ $? -ne 0 ]]; then
        echo "Failed to tag ${app}-rpc image"
        exit 1
    fi

done
