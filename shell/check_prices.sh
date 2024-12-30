#! /bin/bash

source .env

REQUEST=GET
ENDPOINT="/openapi/quote/v1/avgPrice"
TIMESTAMP=$(date +%s%N | cut -b1-13)
PAIR="$1PHP"
QUERY="symbol=$PAIR"
curl -X $REQUEST "$BASE$ENDPOINT?$QUERY"
