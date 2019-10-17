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
    # TYPE: POSTGRES
    ##
    if [ "${BACKUP_TYPE}" == "postgres" ]; then

        # Get PostgreSQL Params
        POSTGRES_PARAMS=""
        export PGHOST=${POSTGRES_HOST}
        export PGPORT=${POSTGRES_PORT}
        export PGUSER=${POSTGRES_USER}
        export PGPASSWORD=${POSTGRES_PASS}
        export DATABASES_TO_IGNORE=${POSTGRES_DATABASES_TO_IGNORE}

        echo ${POSTGRES_PARAMS}

        # Get list of all database names
        ERROR_POSTGRES=$(mktemp)
        DATABASE_NAMES=$(psql ${POSTGRES_PARAMS} -c "\l" 2>${ERROR_POSTGRES} | \
            awk -vcol=Name '!/template|Name|\(|\)|---/{if (NR!=1) print $1}' | \
            sed "s/ //g" | sed "s/|//g" )
        cat ${ERROR_POSTGRES}
        rm -f ${ERROR_POSTGRES}


        # Cleanup latest backup dir
        mkdir -p ${TARGET_DIR}/${SERVER_NAME}/latest/
        rm -rf ${TARGET_DIR}/${SERVER_NAME}/latest/
    
        echo -e "\e[1;36m[START] [POSTGRES]\e[0;36m Starting backup for \e[1;36m${SERVER_NAME}\e[0;39m"
        echo -e "\e[1;36m[INFO ] \e[0;39m${POSTGRES_HOST}:${POSTGRES_PORT}\e[0;39m"
        
        for DATABASE_NAME in ${DATABASE_NAMES}; do

            # Skip if database in "ignore list"
            IGNORE=false
            for DATABASE_TO_IGNORE in ${DATABASES_TO_IGNORE}; do
                if [[ "${DATABASE_NAME}" == "${DATABASE_TO_IGNORE}" ]]; then
                    IGNORE=true
                fi
            done
            if [[ "${IGNORE}" != "false" ]]; then
                continue
            fi
            
            printf '\e[36m%s \e[m%s\n' "[.....]" "Dumping '${DATABASE_NAME}'..."
            
            OUTPUT_DUMP_DIR=${TARGET_DIR}/${SERVER_NAME}/latest/${DATABASE_NAME}
            mkdir -p ${OUTPUT_DUMP_DIR}

            ERROR_POSTGRES=$(mktemp)
            pg_dump --format plain ${DATABASE_NAME} 2>${ERROR_POSTGRES} > ${OUTPUT_DUMP_DIR}/${DATABASE_NAME}.sql
            cat ${ERROR_POSTGRES}
            rm -f ${ERROR_POSTGRES}
            
        done
        
        echo -e "\e[36m[DONE ]\e[39m Saved into ${TARGET_DIR}/${SERVER_NAME}/latest/\n"
    fi

done
