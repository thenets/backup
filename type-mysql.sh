#!/bin/bash

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
cd ${DIR}

# Create cache dir
CACHE_DIR=${DIR}/cache
mkdir -p ${CACHE_DIR}


for CONFIG_FILE in conf.d/*.ini; do
    [ -e "${CONFIG_FILE}" ] || continue

    SERVER_NAME=$(echo ${CONFIG_FILE} | rev | cut -d '/' -f 1 | rev | cut -d '.' -f 1)
    BACKUP_TYPE=$(echo ${CONFIG_FILE} | rev | cut -d '/' -f 1 | rev | cut -d '.' -f 2)
    source ${CONFIG_FILE}
    
    ##
    # TYPE: MYSQL
    ##
    if [ "${BACKUP_TYPE}" == "mysql" ]; then

        # Get MySQL Params
        MYSQL_PARAMS="-u ${MYSQL_USER} -h ${MYSQL_HOST} -P ${MYSQL_PORT} -p${MYSQL_PASS}"

        # Get list of all database names
        ERROR_MYSQL=$(mktemp)
        DATABASE_NAMES=$(mysql ${MYSQL_PARAMS} -r -e "show databases;" 2>${ERROR_MYSQL} | awk -F"\t" '{if (NR!=1) print $1}')
        if [ "$(cat ${ERROR_MYSQL})" != "mysql: [Warning] Using a password on the command line interface can be insecure." ]; then
            echo -e "\e[31m[ERROR ] $(cat ${ERROR_MYSQL})\e[39m"
            rm -f ${ERROR_MYSQL}
            exit 1
        fi
        rm -f ${ERROR_MYSQL}

        # Cleanup latest backup dir
        mkdir -p ${TARGET_DIR}/${SERVER_NAME}/latest/
        rm -rf ${TARGET_DIR}/${SERVER_NAME}/latest/
    
        echo -e "\e[1;36m[START] [MYSQL]\e[0;36m Starting backup for \e[1;36m${SERVER_NAME}\e[0;39m"
        echo -e "\e[1;36m[INFO ] \e[0;39m${MYSQL_HOST}:${MYSQL_PORT}\e[0;39m"
        
        for DATABASE_NAME in ${DATABASE_NAMES}; do
            # Ignore native tables
            if [ "${DATABASE_NAME}" == "information_schema" ]; then continue; fi
            if [ "${DATABASE_NAME}" == "mysql" ]; then continue; fi
            if [ "${DATABASE_NAME}" == "performance_schema" ]; then continue; fi
            if [ "${DATABASE_NAME}" == "sys" ]; then continue; fi

            printf '\e[36m%s \e[m%s\n' "[.....]" "Dumping '${DATABASE_NAME}'..."

            OUTPUT_DUMP_DIR=${TARGET_DIR}/${SERVER_NAME}/latest/${DATABASE_NAME}
            mkdir -p ${OUTPUT_DUMP_DIR}
            
            # Run mysql cmd
            ERROR_MYSQL=$(mktemp)
            mysqldump ${MYSQL_PARAMS} ${DATABASE_NAME} 2>${ERROR_MYSQL} > ${OUTPUT_DUMP_DIR}/${DATABASE_NAME}.sql
            if [ "$(cat ${ERROR_MYSQL})" != "mysqldump: [Warning] Using a password on the command line interface can be insecure." ]; then
                echo -e "\e[31m[ERROR ] $(cat ${ERROR_MYSQL})\e[39m"
                rm -f ${ERROR_MYSQL}
                exit 1
            fi
            rm -f ${ERROR_MYSQL}

        done

        echo -e "\e[36m[DONE ]\e[39m Saved into ${TARGET_DIR}/${SERVER_NAME}/latest/\n"

    fi

done
