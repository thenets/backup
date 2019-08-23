#!/bin/bash

set -e
set -x

# Vars
MYSQL_CONTAINER_NAME=tests_mysql-8_1
MYSQL_PASS=ImWatchingYou
POSTGRES_CONTAINER_NAME=tests_postgres-10_1
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# Clean temp/ dir if exist
rm -rf ${DIR}/temp

# Populate PostgreSQL database
mkdir -p ${DIR}/temp/
cp ${DIR}/samples/postgres/world-1.0.tar.gz ${DIR}/temp/
cd ${DIR}/temp
tar -zxvf world-1.0.tar.gz
cd dbsamples-0.1/world
docker cp world.sql ${POSTGRES_CONTAINER_NAME}:/tmp/
docker exec --user postgres ${POSTGRES_CONTAINER_NAME} psql -c "DROP DATABASE IF EXISTS world;"
docker exec --user postgres ${POSTGRES_CONTAINER_NAME} psql -c "CREATE DATABASE world;"
docker exec --user postgres ${POSTGRES_CONTAINER_NAME} psql -f /tmp/world.sql world

# Populate MySQL database
mkdir -p ${DIR}/temp/
cp ${DIR}/samples/mysql/sakila-db.tar.gz ${DIR}/temp/
cd ${DIR}/temp
tar -zxvf sakila-db.tar.gz
cd sakila-db
docker cp sakila-schema.sql ${MYSQL_CONTAINER_NAME}:/tmp/
docker cp sakila-data.sql ${MYSQL_CONTAINER_NAME}:/tmp/
docker exec ${MYSQL_CONTAINER_NAME} mysql -p${MYSQL_PASS} -e "DROP DATABASE IF EXISTS sakila;"
docker exec ${MYSQL_CONTAINER_NAME} mysql -p${MYSQL_PASS} -e "SOURCE /tmp/sakila-schema.sql;"
docker exec ${MYSQL_CONTAINER_NAME} mysql -p${MYSQL_PASS} -e "SOURCE /tmp/sakila-data.sql;"

# Cleanup
rm -rf ${DIR}/temp
