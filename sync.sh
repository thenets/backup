#!/bin/bash

#set -e

##
# Variables
##

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
cd ${DIR}

for CONFIG_FILE in conf.d/*.ini; do
    [ -e "${CONFIG_FILE}" ] || continue

    # Get ENV from config file
    source ${CONFIG_FILE}

    # Show selected server
    echo -e "\e[1;36m[START]\e[0;36m Starting backup for \e[1;36m${SERVER_NAME}\e[0;39m"

    ##
    # Sync
    #
    # Don't edit the code below.
    ##

    echo -e "\e[36m[.....]\e[39m Sync data from \e[0m${SERVER_DIR}\e[0;39m"
    mkdir -p ${TARGET_DIR}/${SERVER_NAME}/latest/

    # Update latest backup version
    rsync -av --timeout=300 --delete ${SERVER_DIR} ${TARGET_DIR}/${SERVER_NAME}/latest/ 1>/dev/null
    echo -e "\e[36m[DONE ]\e[39m Saved into ${TARGET_DIR}/${SERVER_NAME}/latest/"

done