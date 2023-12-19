#!/bin/bash
docker image prune -f
docker build -t backgo .
docker stop backgo
docker rm backgo

docker run -d --name backgo -p 3005:3005 --env-file="./.env.deploi" backgo
docker logs -f backgo