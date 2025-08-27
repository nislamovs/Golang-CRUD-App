#!/usr/bin/env bash

#https://vsupalov.com/cleaning-up-after-docker/

#docker system prune


docker-compose down -v --rmi all --remove-orphans
docker stop $(docker ps -a -q)
docker rm -f $(docker ps -a -q)
docker rmi -f $(docker images -q)

docker rm -v $(docker ps -aq -f 'status=exited')
docker rmi $(docker images -aq -f 'dangling=true')

docker volume rm $(docker volume ls -q -f 'dangling=true')