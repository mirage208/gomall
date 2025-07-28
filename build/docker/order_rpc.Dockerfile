FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0


RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .

RUN go build -ldflags="-s -w" -o /app/order_rpc app/order/rpc/order.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ=Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/order_rpc /app/order_rpc
COPY app/order/rpc/etc /app/etc

EXPOSE 9003

CMD ["./order_rpc", "-f", "etc/order.yaml"]
