#!/bin/bash

./stage-0-start-servers.sh
sleep 60
./stage-1-populate.sh
./stage-2-build.sh
./stage-3-run-tests.sh
