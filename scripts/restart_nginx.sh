#!/bin/bash

set -e

nginx -t
service nginx restart
