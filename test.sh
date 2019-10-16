#!/bin/bash

DATABASES="pikachu bulbasaur charmander squirtle"

DATABASES_TO_IGNORE="squirtle pikachu"

for DB in ${DATABASES}; do
    IGNORE=false

    for DATABASE_TO_IGNORE in ${DATABASES_TO_IGNORE}; do
        if [[ "${DB}" == "${DATABASE_TO_IGNORE}" ]]; then
            IGNORE=true
        fi
    done

    if [[ "${IGNORE}" == "false" ]]; then
        continue
    fi
    
    echo ${DB}
done
