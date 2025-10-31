# Stage 1: Build
FROM golang:1.25 AS builder

#working directory:
WORKDIR /app

#copy project files
COPY . /app
#run the scripts
RUN chmod +x /app/scripts/build-application.sh && \
    bash /app/scripts/build-application.sh


# Stage 2: Runtime
FROM alpine:latest
WORKDIR /app
#copy compiled binary
COPY --from=builder /app/ordersystem /app/ordersystem
#publish it
EXPOSE 3000
CMD ["/app/ordersystem"]
