FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0


RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .

RUN go build -ldflags="-s -w" -o /app/cart_rpc app/cart/rpc/cart.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ=Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/cart_rpc /app/cart_rpc
COPY app/cart/rpc/etc /app/etc

EXPOSE 9005

CMD ["./cart_rpc", "-f", "etc/cart.yaml"]
