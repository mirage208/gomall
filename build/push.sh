#!/usr/bin/env bash

for app in user payment product order cart; do

    # Push API image
    docker push crpi-uuzhidt8pbh6o5xf.cn-shanghai.personal.cr.aliyuncs.com/mirage208/mall:${app}-api
    if [[ $? -ne 0 ]]; then
        echo "Failed to push ${app}-api image"
        exit 1
    fi

    # Push RPC image
    docker push crpi-uuzhidt8pbh6o5xf.cn-shanghai.personal.cr.aliyuncs.com/mirage208/mall:${app}-rpc
    if [[ $? -ne 0 ]]; then
        echo "Failed to push ${app}-rpc image"
        exit 1
    fi

done
