#!/bin/bash

##
# Variables
##

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
cd ${DIR}

GOT_ERROR=0

for CONFIG_FILE in conf.d/*.ini; do
    [ -e "${CONFIG_FILE}" ] || continue

    # Get ENV from config file
    source ${CONFIG_FILE}

    ##
    # Test connection
    #
    # Check if current user has access to the target server.
    ##
    SERVER_CONN=$(echo ${SERVER_DIR} | sed 's/\:\(.*\)//g' | sed 's/\(.*\)=//g')
    SERVER_IP=$(cut -d'@' -f2 <<<${SERVER_CONN})

    # Add or re-add key to the know_hosts
    ssh-keygen -f "~/.ssh/known_hosts" -R ${SERVER_IP} 1>/dev/null 2>/dev/null || true
    ssh-keyscan -H ${SERVER_IP} >> ~/.ssh/known_hosts 2>/dev/null || true

    TEST_OUT=$(ssh -T ${SERVER_CONN} echo ok 2>&1)

    if [[ "${TEST_OUT}" == "ok" ]]; then
        echo -e "\e[1;32m[SUCCESS]\e[0;32m Access granted to the \e[1;32m${SERVER_NAME}\e[0;32m (${SERVER_CONN})\e[0;39m"
    else
        echo -e "\e[1;31m[ERROR  ]\e[0;31m Access denied to the \e[0;31m${SERVER_NAME}\e[0;39m"
        echo -e "\e[1;31m${SERVER_CONN}:\e[0;31m ${TEST_OUT}\e[39m"
        GOT_ERROR=1
    fi

done


if [[ ${GOT_ERROR} == "1" ]]; then
    printf "\nError during test connection to the servers!\n" >&2
    exit 1
fi