#!/bin/bash
echo "Beginning Amazon ECS Deployment Script"
$(aws ecr get-login --region us-west-1)
echo "Acquired ECR credentials"
docker build -t signal_api_server .
echo "Built docker image"
docker tag signal_api_server:latest 289307572179.dkr.ecr.us-west-1.amazonaws.com/signal_api_server:latest
echo "Tagged"
docker push 289307572179.dkr.ecr.us-west-1.amazonaws.com/signal_api_server:latest
echo "Pushed docker image to ECR"
