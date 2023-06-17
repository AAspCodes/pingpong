#!/bin/bash


sudo docker build -t aaspcodes/client:latest . -f ./client/Dockerfile

sudo docker build -t aaspcodes/server:latest . -f ./server/Dockerfile

docker push aaspcodes/client:latest
docker push aaspcodes/server:latest


# sudo docker run aaspcodes/client
# sudo docker run aaspcodes/server