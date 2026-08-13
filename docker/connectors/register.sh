#!/bin/sh

until curl -s http://connect:8083/connectors > /dev/null; do
    sleep 2
done

curl -X POST \
    -H "Content-Type: application/json" \
    --data @/connectors/warehouse-connector.json \
    http://connect:8083/connectors