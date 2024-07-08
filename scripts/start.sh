#!/bin/bash

docker compose down
docker rm -f $(docker ps -a -q)
docker volume rm $(docker volume ls -q)
docker build . -t app-img

docker secret inspect cert.pem &> /dev/null
if [ $? -ne 0 ]; then
    docker secret create cert.pem ./nginx/cert.pem
else
    echo "The secret 'cert.pem' allready exist. Skip."
fi

docker secret inspect key.pem &> /dev/null
if [ $? -ne 0 ]; then
    docker secret create key.pem ./nginx/key.pem
else
    echo "The secret 'cert.pem' allready exist. Skip."
fi

docker secret inspect ca.crt &> /dev/null
if [ $? -ne 0 ]; then
    docker secret create ca.crt ./nginx/certs/ca.crt
else
    echo "The secret 'cert.pem' allready exist. Skip."
fi

docker compose up -d 

# docker compose up -d --scale app-image=0

