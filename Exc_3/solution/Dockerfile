# ==========================
# Stage 1: Builder
# ==========================
FROM golang:1.25 AS builder

WORKDIR /app
COPY . /app
RUN chmod +x /app/scripts/build-application.sh && \
    bash /app/scripts/build-application.sh

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/ordersystem /app/ordersystem
EXPOSE 3000
ENTRYPOINT ["/app/ordersystem"]
