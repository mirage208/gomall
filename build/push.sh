#!/usr/bin/env bash

for app in user payment product order cart; do
    # Push RPC image
    docker push crpi-uuzhidt8pbh6o5xf.cn-shanghai.personal.cr.aliyuncs.com/mirage208/mall:${app}_rpc
done
