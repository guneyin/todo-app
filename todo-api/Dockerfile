FROM golang:1.23-bookworm as builder

WORKDIR /app

COPY . .

RUN make init
RUN go build -v -o service .

FROM debian:bookworm-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/service /app/service

WORKDIR /app

CMD ["/app/service"]
