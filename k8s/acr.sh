#!/bin/bash

#build image
az acr build  --registry allahthedev --image search-nearby-places:1.0

#show tags
az acr repository show-tags --name allahthedev --repository search-nearby-places

#create azure ad principal # this will be required for creating aks cluster
az ad sp create-for-rbac --name serach-nearby-places-principal

#create cluster # but this didn't work

AKS_RESOURCE_GROUP=
AKS_CLUSTER_NAME=
ACR_RESOURCE_GROUP=
ACR_NAME=

# Get the id of the service principal configured for AKS
CLIENT_ID=$(az aks show --resource-group $AKS_RESOURCE_GROUP --name $AKS_CLUSTER_NAME --query "servicePrincipalProfile.clientId" --output tsv)

# Get the ACR registry resource id
ACR_ID=$(az acr show --name $ACR_NAME --resource-group $ACR_RESOURCE_GROUP --query "id" --output tsv)

# Create role assignment
az role assignment create --assignee $CLIENT_ID --role acrpull --scope $ACR_ID

#get kubrnetes server credential to make it accessible for client
az aks get-credentials --resource-group search-nearby-places --name search-nearby-places

#do the deployment
kubectl apply -f deployment.yaml

#create a loadbalancer service
kubectl apply -f loadbalancer.yaml


