#!/bin/sh
set -e

docker stop orderservice postgres18 2>/dev/null || true
docker rm orderservice postgres18 2>/dev/null || true

set -a
. ./debug.env
set +a

# todo
# docker build
docker build -t orderservice .
echo "built"

# docker run db

  # --network host and no port
docker run -d \
  --name postgres18 \
  --env-file ./debug.env \
  --network orders-net \
  -v pgdata:/var/lib/postgresql/18/docker \
  -p 5432:5432 \
  postgres:18

echo "run db"
# docker run orderservice
sleep 10

#  -p 8080:8080 \
docker run -d \
  --name orderservice \
  --env-file ./debug.env \
  --network orders-net \
  -p 3000:3000 \
  orderservice:latest
echo "run orderservice"
