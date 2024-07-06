#!/bin/bash

docker compose down
docker rm -f $(docker ps -a -q)
docker volume rm $(docker volume ls -q)
docker compose build app-image 
docker compose up -d --scale app-image=0
