#!/bin/bash

#set -e

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
cd ${DIR}

./test_connection.sh
echo ""
./type-postgres.sh
./type-mysql.sh
./type-dir.sh

./compress.sh
echo ""
./remove_old.sh

echo ""
echo "Free space"
df -h
