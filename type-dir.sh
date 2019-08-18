#!/bin/bash

#set -e

##
# Variables
##

# Default values
SSH_PORT=22

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
cd ${DIR}

for CONFIG_FILE in conf.d/*.ini; do
    [ -e "${CONFIG_FILE}" ] || continue

    SERVER_NAME=$(echo ${CONFIG_FILE} | rev | cut -d '/' -f 1 | rev | cut -d '.' -f 1)
    BACKUP_TYPE=$(echo ${CONFIG_FILE} | rev | cut -d '/' -f 1 | rev | cut -d '.' -f 2)
    source ${CONFIG_FILE}

    ##
    # TYPE: DIR
    ##
    if [ "${BACKUP_TYPE}" == "dir" ]; then
        # Show selected server
        echo -e "\e[1;36m[START] [DIR]\e[0;36m Starting backup for \e[1;36m${SERVER_NAME}\e[0;39m"

        echo -e "\e[36m[.....]\e[39m Sync data from \e[0m${SERVER_DIR}\e[0;39m"
        mkdir -p ${TARGET_DIR}/${SERVER_NAME}/latest/

        # Update latest backup version
        rsync -av --timeout=300 --delete -e "ssh -p ${SSH_PORT}" ${SERVER_DIR} ${TARGET_DIR}/${SERVER_NAME}/latest/ 1>/dev/null
        echo -e "\e[36m[DONE ]\e[39m Saved into ${TARGET_DIR}/${SERVER_NAME}/latest/\n"
    fi

    # Default values
    SSH_PORT=22

done
