#!/bin/bash
./tear_down.sh
./build_docker_images.sh
kubectl apply -f server/service.yaml
kubectl apply -f server/deployment.yaml

python test_runner/main.py