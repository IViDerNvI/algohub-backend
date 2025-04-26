# 第一阶段：构建应用
FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux

RUN go build -o ./algohub ./cmd/apiserver && chmod +x ./algohub

FROM alpine:latest  

RUN apk add --no-cache ca-certificates

WORKDIR /root

RUN mkdir -p conf/cert

COPY --from=builder /app/algohub ./algohub

# COPY --from=builder /app/conf/cert/public.crt ./conf/cert/public.crt
# COPY --from=builder /app/conf/cert/private.crt ./conf/cert/private.key

ENV ALGOHUB_MYSQL_HOSTNAME="localhost"
ENV ALGOHUB_MINIO_ENDPOINT="localhost:9000"
ENV ALGOHUB_JUDGE_RPC_ENDPOINT="localhost:50052"
ENV ALGOHUB_REDIS_HOSTNAME="localhost:6379"
ENV ALGOHUB_JUDGE_RPC_ENDPOINT="localhost:50051"

EXPOSE 8080
EXPOSE 8443

CMD ["./algohub", "apiserver"]