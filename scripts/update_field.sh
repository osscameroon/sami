#!/bin/bash

set -e

USAGE="./update_field.sh <deployment_folder> <service_name> <environment> <path_field> <path_value> <deployment_type> <docker-compose.yaml>
Example: ./update_field.sh /Users/user_name/deployments camerdevs stage ".services.api.image" "camerdevs-api@sha1234" compose docker-compose.yml

The deploment_directory should not end with a \"/\""

if [[ $# -ne 7 ]]; then
	echo "Wrong usage: $USAGE"
	exit 1
fi

# parse paramater
DEPLOY_DIR=$1
SERVICE_NAME=$2
ENVIRONMENT=$3
PATCH_FIELD=$4
PATCH_VALUE=$5
#deployment type (compose | kubernetes | nomad | jobs)
TYPE=$6
#the name of the file we want to patch
FILE_NAME=$7


SERVICE_DIR="services"
FILE_FULL_PATH=$DEPLOY_DIR/$SERVICE_DIR/$SERVICE_NAME/$TYPE/$ENVIRONMENT/$FILE_NAME

echo "DEPLOY_DIR: $DEPLOY_DIR"
echo "SERVICE_NAME: $SERVICE_NAME"
echo "ENVIRONMENT: $ENVIRONMENT"
echo "PATCH_FIELD: $PATCH_FIELD"
echo "TYPE: $TYPE"
echo "FILE_NAME: $FILE_NAME"
echo "PATCH_VALUE: $PATCH_VALUE"
echo "FILE_FULL_PATH: $FILE_FULL_PATH"

 yq eval -i "$PATCH_FIELD = \"$PATCH_VALUE\"" $FILE_FULL_PATH
