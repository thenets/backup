#!/bin/bash

#set -e

##
# Variables
##

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
cd ${DIR}

for CONFIG_FILE in conf.d/*.ini; do
    [ -e "${CONFIG_FILE}" ] || continue

    SERVER_NAME=$(echo ${CONFIG_FILE} | rev | cut -d '/' -f 1 | rev | cut -d '.' -f 1)

    # Get ENV from config file
    source ${CONFIG_FILE}

    # Show selected server
    echo -e "\e[1;36m[START]\e[0;36m Compressing backup from \e[1m${SERVER_NAME}\e[0m \e[39m"

    # Get current time
    CURRENT_DATE=$(date +"%y-%m-%d")

    # Compress files
    ORIGIN_DIR=${TARGET_DIR}/${SERVER_NAME}/latest/
    OUT_DIR=${TARGET_DIR}/${SERVER_NAME}/older/${CURRENT_DATE}

    # Compress each subdir
    if [[ "${COMPRESS_SUBDIR}" == "1" ]]; then
        SUBDIRS_TO_BACKUP=$(find $ORIGIN_DIR -maxdepth 1 ! -path $ORIGIN_DIR -type d)

        for D_PATH in $SUBDIRS_TO_BACKUP; do
            D_NAME=${D_PATH/$ORIGIN_DIR/}
            D_NAME="${D_NAME}"
            OUT_FILE=$OUT_DIR"/./"$D_NAME"_"$CURRENT_DATE".tar.gz"

            mkdir -p ${OUT_DIR}

            # Check if backup already exists for the same $CURRENT_DATE
            if [ -f ${OUT_FILE} ]; then
                echo -e "\e[33m[SKIP ]\e[39m Backup $OUT_FILE already exist."
            else
                # Create backup
                printf '\e[36m%s \e[m%s' "[.....]" "Compressing to ${OUT_FILE}..."
                sh -c "cd ${ORIGIN_DIR} && tar -zcf ${OUT_FILE} ${D_NAME} 1>/dev/null"
                echo -e " \e[36mok\e[39m"
            fi
        done

    # Compress entire dir
    else
        OUT_FILE=${OUT_DIR}/./${SERVER_NAME}_${CURRENT_DATE}.tar.gz

        mkdir -p ${OUT_DIR}

        # Check if backup already exists for the same $CURRENT_DATE
        if [ -f ${OUT_FILE} ]; then
            echo -e "\e[33m[SKIP]\e[39m Backup $OUT_FILE already exist."
        else
            # Create backup
            printf '\e[36m%s \e[m%s' "[.....]" "Compressing to ${OUT_FILE}..."
            sh -c "cd ${ORIGIN_DIR} && tar -zcf ${OUT_FILE} . 1>/dev/null"
            echo -e "\e[36mok\e[39m"
        fi
    fi

done