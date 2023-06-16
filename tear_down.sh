#!/bin/bash

kubectl delete deployment server client
kubectl delete service server-service



for pod_name in $(kubectl get pods | grep -v NAME | awk '{print $1}') 
do
    kubectl delete pod $pod_name
done

for service_name in $(kubectl get services | grep client | awk '{print $1}') 
do
    kubectl delete service $service_name
done
