#!/bin/bash

##
# Variables
##

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
cd ${DIR}

GOT_ERROR=0

for CONFIG_FILE in conf.d/*.ini; do
    [ -e "${CONFIG_FILE}" ] || continue

    SERVER_NAME=$(echo ${CONFIG_FILE} | rev | cut -d '/' -f 1 | rev | cut -d '.' -f 1)
    
    # Get ENV from config file
    source ${CONFIG_FILE}

    # Show selected server
    echo -e "\e[1;36m[START]\e[0;36m Removing old backup from \e[1;36m${SERVER_NAME}\e[0;39m"
    echo -e "        Max age: ${DELETE_OLDER_THAN_X_DAYS} days"

    ##
    # Remove old
    #
    # Remove files older than X days.
    ##

    # Compress files
    BACKUP_DIR=${TARGET_DIR}/${SERVER_NAME}/older/

    # Compress each subdir
    REVISION_DIR_LIST=$(find ${BACKUP_DIR} -maxdepth 1 ! -path ${BACKUP_DIR} -type d)

    for REVISION_DIR in ${REVISION_DIR_LIST}; do
        REVISION_DATE=$(echo ${REVISION_DIR} | sed 's/\(.*\)older\///g')

        datetime="${REVISION_DATE}"
        timeago="${DELETE_OLDER_THAN_X_DAYS} days ago"

        dtSec=$(date --date "$datetime" +'%s')
        taSec=$(date --date "$timeago" +'%s')

        # DEBUG
        #echo "INFO: dtSec=$dtSec, taSec=$taSec" >&2
        #[ $dtSec -lt $taSec ] && echo remove
        #! [ $dtSec -lt $taSec ] && echo not remove

        if [ $dtSec -lt $taSec ]; then
            printf "\e[33m%s\e[39m" "[.....] Removing ${REVISION_DIR}..."
            rm -rf ${REVISION_DIR}
            echo -e "\e[33m ok\e[39m"
        fi

    done

done