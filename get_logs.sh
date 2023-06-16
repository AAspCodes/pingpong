#!/bin/bash
echo -e "\n\nserver logs:"
kubectl logs $(kubectl get pods | grep server | awk '{print $1}')

for client_name in $(kubectl get pods | grep client | awk '{print $1}')
do
    echo -e "\n\n$client_name logs:"
    kubectl logs $client_name
done

echo -e "\n\n"