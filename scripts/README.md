# `/scripts`

执行各种构建、安装、分析等操作的脚本。

这些脚本保持了根级别的 Makefile 变得小而简单

- `model_gen.sh`: 生成model代码
- `rpc_gen.sh`: 生成rpc代码
- `api_gen.sh`: 生成api代码
- `dockerfile_gen.sh`: 生成Dockerfile文件
- `update_ports.sh`: 更新所有服务的端口配置

## 端口分配方案

| 服务    | API端口 | RPC端口 | API Metrics | RPC Metrics |
| ------- | ------- | ------- | ----------- | ----------- |
| user    | 8001    | 9001    | 6001        | 6011        |
| product | 8002    | 9002    | 6002        | 6012        |
| order   | 8003    | 9003    | 6003        | 6013        |
| payment | 8004    | 9004    | 6004        | 6014        |
| cart    | 8005    | 9005    | 6005        | 6015        |

**说明:**
- API端口: HTTP REST API服务
- RPC端口: gRPC内部通信服务  
- API Metrics: API服务的监控指标端口
- RPC Metrics: RPC服务的监控指标端口