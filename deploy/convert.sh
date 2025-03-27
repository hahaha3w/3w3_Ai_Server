#!/bin/zsh
kompose convert -f services-compose.yaml -o service-manifests/
kubectl apply -f ./service-manifests


