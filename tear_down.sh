#!/bin/bash

kubectl delete deployment server client
kubectl delete service server-service



for pod_name in $(kubectl get pods | grep -v NAME | awk '{print $1}') 
do
    kubectl delete pod $pod_name
done
# kubectl get services | grep -v NAME | awk '{print $1}' | xargs kubectl delete service