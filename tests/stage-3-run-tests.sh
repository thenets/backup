#!/bin/bash

set -e
set -x

# Vars
IMAGE_TAG=thenets/backup
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# Run test
docker run --rm \
    --name thenets-backup \
    --network tests_thenets \
    -v ${DIR}/conf.d:/app/conf.d \
    -v ${DIR}/docker-sshd/client.pub:/root/.ssh/id_rsa.pub \
    -v ${DIR}/docker-sshd/client:/root/.ssh/id_rsa \
    ${IMAGE_TAG}
