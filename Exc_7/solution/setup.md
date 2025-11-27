# todo note commands

### Install docker on all nodes
```sh
curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh`
```

### Initialize docker swarm on the controller
```sh
docker swarm init --advertise-addr 192.168.2.150
```

### Join the worker nodes
```sh
docker swarm join --token <YOUR_TOKEN_HERE> 192.168.2.150:2377
```

### Check if the cluster has 3 nodes
```sh
docker node ps
```

### Log into github package registry
- Generate an access token in github and execute:
```sh
docker login ghcr.io -u S2410929034
```

### Build images for ARM
```
docker build -f Dockerfile.ordersystem --platform linux/arm64 -t ghcr.io/s2410929034/sdb-ais-exercise/ordersystem:1.0-arm64 .

docker push ghcr.io/s2410929034/sdb-ais-exercise/ordersystem:1.0-arm64

docker build -f Dockerfile.frontend --platform linux/arm64 -t ghcr.io/s2410929034/sdb-ais-exercise/frontend:1.0-arm64 .

docker push ghcr.io/s2410929034/sdb-ais-exercise/frontend:1.0-arm64
```

### Copy all files to target machine
- Can be done via winscp or ssh
- Copy the docker directory
- Also copy docker-compose.yml

### Deploy the containers with docker swarm
```sh
docker stack deploy -c docker-compose.yml sdb-stack
```

### See your containers
```sh
docker service ls
```

### Forward traefik on port 80 to your local machine
```sh
ssh -L 80:127.0.0.1:80 root@cp01
```

### To delete your deployment use
```sh
docker stack rm sdb-stack
```


--------------------------------
Worked with Jonas Welt & Karim Ortner