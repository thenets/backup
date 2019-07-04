#!/bin/bash

set -e

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
cd ${DIR}

./test_connection.sh
echo ""
./sync.sh
echo ""
./compress.sh
echo ""
./remove_old.sh
