#!/bin/bash


sudo docker build -t aaspcodes/client:latest client/

sudo docker build -t aaspcodes/server:latest server/

docker push aaspcodes/client:latest
docker push aaspcodes/server:latest


# sudo docker run aaspcodes/client
# sudo docker run aaspcodes/server