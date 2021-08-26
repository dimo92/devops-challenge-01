#!/bin/bash

ACTION=$1
DEPLOY="u"
UNDEPLOY="d"
RESTART="r"
DOCKER_COMPOSE_FILE_NAME="docker-compose-db-mysql.yml"

echo "Action: $ACTION"

if [ "$ACTION" == "$DEPLOY" ]; then
   echo "Starting to deploy Docker-compose..."
   docker-compose -f "$DOCKER_COMPOSE_FILE_NAME" up -d
   #docker-compose -f docker-compose-$parameterA.yml -p $parameterA up -d
   echo "Docker-compose is running!"

elif [ "$ACTION" == "$UNDEPLOY" ]; then
   echo "Starting to undeploy Docker-compose..."
   docker-compose -f "$DOCKER_COMPOSE_FILE_NAME" down -v && docker-compose rm -fsv
   #docker-compose -f docker-compose-$parameterA.yml -p $parameterA down
   echo "Docker-compose is down!"
elif [ "$ACTION" == "$RESTART" ]; then
   echo "Starting to restart Docker-compose..."
   docker-compose -f "$DOCKER_COMPOSE_FILE_NAME" down -v && docker-compose rm -fsv
   docker-compose -f "$DOCKER_COMPOSE_FILE_NAME" up -d 
   #docker-compose -f docker-compose-$parameterA.yml -p $parameterA down
   echo "Docker-compose is restared!"
else
   echo "Invalid Action!"
fi