#!/bin/bash

./stage-0-start-servers.sh
sleep 60
./stage-1-populate.sh
./stage-2-run-tests.sh
