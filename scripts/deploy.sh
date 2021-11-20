#!/bin/bash

USAGE="./deploy.sh <deploment_directory>
Example: ./deploy.sh /Users/user_name/deployments
The deploment_directory should not end with a \"/\""

DEPLOY_DIR=$1

if [ -z "$DEPLOY_DIR" ]; then
	echo "Wrong usage: $USAGE"
	exit 0
fi

SERVICE_DIR="services"
CONF_DIR="conf"
DEPLOY_SERVICE_DIR=$DEPLOY_DIR/$SERVICE_DIR
DEPLOY_CONF_DIR=$DEPLOY_DIR/$CONF_DIR

echo "DEPLOY_CONF_DIR: $DEPLOY_CONF_DIR"
echo "DEPLOY_SERVICE_DIR: $DEPLOY_SERVICE_DIR"

ls $DEPLOY_SERVICE_DIR
ls $DEPLOY_CONF_DIR

echo "Start deployment..."
cd $DEPLOY_SERVICE_DIR
DOCKER_COMPOSE_FILES="$(find . -regex ".*-compose\.\(yml\|yaml\)")"
echo "DOCKER_COMPOSE_FILES:"
echo "$DOCKER_COMPOSE_FILES"

DOCKER_STACK_FILES="$(find . -regex ".*-stack\.\(yml\|yaml\)")"
echo "DOCKER_STACK:"
echo "$DOCKER_STACK_FILES"

for file in $DOCKER_COMPOSE_FILES; do
	docker-compose -f "$file" up -d
done

for file in $DOCKER_STACK_FILES; do
	docker stack deploy -c "$file" $(echo "$file" | cut -d "/" -f 2)
done
