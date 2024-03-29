#!/bin/bash

##
# Test connection
#
# Check if current user has access to the target server.
##

# Default
SSH_PORT=22

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
cd ${DIR}

GOT_ERROR=0

for CONFIG_FILE in conf.d/*.ini; do
    [ -e "${CONFIG_FILE}" ] || continue

    SERVER_NAME=$(echo ${CONFIG_FILE} | rev | cut -d '/' -f 1 | rev | cut -d '.' -f 1)
    BACKUP_TYPE=$(echo ${CONFIG_FILE} | rev | cut -d '/' -f 1 | rev | cut -d '.' -f 2)
    source ${CONFIG_FILE}

    ##
    # TYPE: DIR
    ##
    if [ "${BACKUP_TYPE}" == "dir" ]; then
        SERVER_CONN=$(echo ${SERVER_DIR} | sed 's/\:\(.*\)//g' | sed 's/\(.*\)=//g')
        SERVER_IP=$(cut -d'@' -f2 <<<${SERVER_CONN})

        # Add or re-add key to the know_hosts
        ssh-keygen -f "~/.ssh/known_hosts" -R ${SERVER_IP} 1>/dev/null 2>/dev/null || true
        ssh-keyscan -p ${SSH_PORT} -H ${SERVER_IP} >> ~/.ssh/known_hosts 2>/dev/null || true

        # HACK first one makes new "warning" be ignored
        TEST_OUT=$(ssh -p ${SSH_PORT} -T ${SERVER_CONN} echo ok 2>&1)
        TEST_OUT=$(ssh -p ${SSH_PORT} -T ${SERVER_CONN} echo ok 2>&1)

        if [[ "${TEST_OUT}" == "ok" ]]; then
            echo -e "\e[1;32m[SUCCESS] [DIR]\e[0;32m Access granted to the \e[1;32m${SERVER_NAME}\e[0;32m (${SERVER_CONN})\e[0;39m"
        else
            echo -e "\e[1;31m[ERROR  ]\e[0;31m Access denied to the \e[0;31m${SERVER_NAME}\e[0;39m"
            echo -e "\e[1;31m${SERVER_CONN}:\e[0;31m ${TEST_OUT}\e[39m"
            GOT_ERROR=1
        fi
    fi

    
    ##
    # TYPE: MYSQL
    ##
    if [ "${BACKUP_TYPE}" == "mysql" ]; then
        MYSQL_PARAMS="-u ${MYSQL_USER} -h ${MYSQL_HOST} -P ${MYSQL_PORT} -p${MYSQL_PASS}"
        ERROR_MYSQL=$(mktemp)
        DATABASE_NAMES=$(mysql ${MYSQL_PARAMS} -r -e "show databases;" 2>${ERROR_MYSQL} | awk -F"\t" '{if (NR!=1) print $1}')
        if [ "$(cat ${ERROR_MYSQL})" != "mysql: [Warning] Using a password on the command line interface can be insecure." ]; then
            echo -e "\e[1;31m[ERROR  ]\e[0;31m Access denied to the \e[0;31m${SERVER_NAME}\e[0;39m"
            echo -e "\e[31m[.......] $(cat ${ERROR_MYSQL})\e[39m"
            rm -f ${ERROR_MYSQL}
        else
            echo -e "\e[1;32m[SUCCESS] [MYSQL]\e[0;32m Access granted to the \e[1;32m${SERVER_NAME}\e[0;32m (${MYSQL_USER}@${MYSQL_HOST})\e[0;39m"
        fi
        rm -f ${ERROR_MYSQL}
    fi

    
    ##
    # TYPE: POSTGRES
    ##
    if [ "${BACKUP_TYPE}" == "postgres" ]; then
        export PGHOST=${POSTGRES_HOST}
        export PGPORT=${POSTGRES_PORT}
        export PGUSER=${POSTGRES_USER}
        export PGPASSWORD=${POSTGRES_PASS}
        
        ERROR_POSTGRES=$(mktemp)
        DATABASE_NAMES=$(psql -c "\l" 2>${ERROR_POSTGRES} )
        
        if [ "$(cat ${ERROR_POSTGRES})" != "" ]; then
            echo -e "\e[1;31m[ERROR  ]\e[0;31m Access denied to the \e[0;31m${SERVER_NAME}\e[0;39m"
            echo -e "\e[31m[.......] $(cat ${ERROR_POSTGRES})\e[39m"
            rm -f ${ERROR_POSTGRES}
        else
            echo -e "\e[1;32m[SUCCESS] [POSTGRES]\e[0;32m Access granted to the \e[1;32m${SERVER_NAME}\e[0;32m (${POSTGRES_USER}@${POSTGRES_HOST})\e[0;39m"
        fi

        rm -f ${ERROR_POSTGRES}
    fi

    # Default
    SSH_PORT=22


done


if [[ ${GOT_ERROR} == "1" ]]; then
    printf "\nError during test connection to the servers!\n" >&2
    exit 1
fi
