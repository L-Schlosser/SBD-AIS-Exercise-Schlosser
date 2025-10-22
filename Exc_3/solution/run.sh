#!/bin/sh

docker stop orderservice postgres18 2>/dev/null || true
docker rm orderservice postgres18 2>/dev/null || true

# todo
# docker build
docker build -t orderservice .
echo "built"

# docker run db
docker run -d \
  --name postgres18 \
  --env-file ./debug.env \
  --network orders-net \
  -v pgdata:/var/lib/postgresql/18/docker \
  -p 5432:5432 \
  postgres:18

echo "run db"

sleep 10

# docker run orderservice:
docker run -d \
  --name orderservice \
  --env-file ./debug.env \
  --network orders-net \
  -p 3000:3000 \
  orderservice:latest
echo "run orderservice"

echo "http://localhost:3000/openapi/index.html"
