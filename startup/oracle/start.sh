#!/bin/bash

set -m
./start-db.sh &
if [[ "$INSERT_TABLE" -eq 1 ]]; then
    echo "--- INSERT TABLES ---";
    ./insert-tables.sh;
fi
fg %1
sleep infinity