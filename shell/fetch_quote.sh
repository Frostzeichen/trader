#! /bin/bash

source .env
REQUEST=POST
ENDPOINT="/openapi/convert/v1/get-quote"
TIMESTAMP=$(date +%s%N | cut -b1-13)
sourceCurrency=PHP
targetCurrency=$1
sourceAmount=60
QUERY="sourceCurrency=$sourceCurrency&targetCurrency=$targetCurrency&sourceAmount=$sourceAmount&timestamp=$TIMESTAMP"
SIGNATURE=$(echo -n $QUERY | openssl dgst -sha256 -hmac $SECRET)
SIGNATURE=${SIGNATURE#"SHA2-256(stdin)= "}
# echo "$BASE$ENDPOINT?$QUERY&signature=$SIGNATURE"
# echo "X-COINS-APIKEY: $KEY"
curl -X $REQUEST "$BASE$ENDPOINT?$QUERY&signature=$SIGNATURE" -H "X-COINS-APIKEY: $KEY" | jq
sleep 5 # Protects against error 10000003 (too many requests)
