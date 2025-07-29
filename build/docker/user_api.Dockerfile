FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0


RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .

RUN go build -ldflags="-s -w" -o /app/user_api app/user/api/user.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ=Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/user_api /app/user_api
COPY app/user/api/etc /app/etc

EXPOSE 8001

CMD ["./user_api", "-f", "etc/user.yaml"]
