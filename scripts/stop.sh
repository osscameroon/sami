#!/bin/bash

USAGE="./stop.sh <deploment_directory>
Example: ./dep.loy.sh /Users/user_name/deployments/services
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

echo "Stop deployments..."
cd $DEPLOY_SERVICE_DIR
DOCKER_COMPOSE_FILES="$(find . -name "docker-compose.*")"
echo "DOCKER_COMPOSE_FILES:"
echo "$DOCKER_COMPOSE_FILES"

for file in $DOCKER_COMPOSE_FILES; do
	set -x
	docker-compose -f "$file" down
done
