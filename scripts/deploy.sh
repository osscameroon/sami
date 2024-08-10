#!/bin/bash

set -e

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

echo "DEPLOY_SERVICE_DIR: $DEPLOY_SERVICE_DIR"

ls $DEPLOY_SERVICE_DIR
ls $DEPLOY_CONF_DIR

#update the repository
cd $DEPLOY_SERVICE_DIR
git config --global user.name "osscameroon-bot"

if [ ! -z "$GITHUB_TOKEN" ]; then
	#get repo url takes a `  Fetch URL: git@github.com:elhmn/infra.git`
	#and returns `@github.com/elhmn/infra.git`
	REPO_PATH=$(git remote show origin | grep Fetch | sed -E 's#^.*(git@|https://)([^/]+)/(.*)$#@\2:\3#;s#:#/#g')
	git pull --rebase https://$GITHUB_TOKEN$REPO_PATH main
else
	#get repo url takes a `  Fetch URL: git@github.com:elhmn/infra.git`
	#and returns `@github.com:elhmn/infra.git`
	REPO_PATH=$(git remote show origin | grep Fetch | sed -E 's#^.*(git@|https://)([^/]+)/(.*)$#@\2/\3#')
	git pull --rebase git$REPO_PATH
fi
cd -

echo "Start deployment..."
cd $DEPLOY_SERVICE_DIR
DOCKER_COMPOSE_FILES="$(find . -regex ".*-compose\.\(yml\|yaml\)")"
echo "DOCKER_COMPOSE_FILES:"
echo "$DOCKER_COMPOSE_FILES"

DOCKER_STACK_FILES="$(find . -regex ".*-stack\.\(yml\|yaml\)")"
echo "DOCKER_STACK:"
echo "$DOCKER_STACK_FILES"

# make sure that we add the user to the docker group
sudo usermod -aG docker $USER

for file in $DOCKER_COMPOSE_FILES; do
	name=$(echo "$file" | cut -d "/" -f 2)
	env=$(echo "$file" | cut -d "/" -f 4)
	name="${env}_${name}"
	sudo docker compose -f "$file" -p "$name" up -d
done

for file in $DOCKER_STACK_FILES; do
	echo file: $file
	name=$(echo "$file" | cut -d "/" -f 2)
	env=$(echo "$file" | cut -d "/" -f 4)
	name="${env}_${name}"
	sudo docker stack deploy  --with-registry-auth -c "$file" "$name" --detach=true
done
