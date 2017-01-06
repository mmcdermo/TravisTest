#!/bin/bash

$(aws ecr get-login --region us-west-1)

docker build -t signal_api_server .

docker tag signal_api_server:latest 289307572179.dkr.ecr.us-west-1.amazonaws.com/signal_api_server:latest

docker push 289307572179.dkr.ecr.us-west-1.amazonaws.com/signal_api_server:latest
