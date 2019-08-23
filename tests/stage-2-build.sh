#!/bin/bash

set -e
set -x

# Vars
IMAGE_TAG=thenets/backup
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# Build
docker build -t ${IMAGE_TAG} ${DIR}/..
