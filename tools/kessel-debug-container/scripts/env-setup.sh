#!/bin/bash

CDAPPCONFIG="/cdapp/cdappconfig.json"
echo "Setting up Kafka Connection Info..."

export KAFKA_AUTH_CONFIG=/tmp/config.props

if [[ $(cat $CDAPPCONFIG | jq '.kafka.brokers | length') -gt 1 ]]; then
    export BOOTSTRAP_SERVERS=$(cat $CDAPPCONFIG | jq -r '.kafka.brokers[] | "\(.hostname):\(.port),"')
else
    export BOOTSTRAP_SERVERS=$(cat $CDAPPCONFIG | jq -r '.kafka.brokers[] | "\(.hostname):\(.port)"')
fi

if [[ $(cat $CDAPPCONFIG | jq '.kafka.brokers[] | has("sasl")') == "true" ]]; then
    export KAFKA_USERNAME=$(cat $CDAPPCONFIG | jq -r '.kafka.brokers[0].sasl.username')
    export KAFKA_PW=$(cat $CDAPPCONFIG | jq -r '.kafka.brokers[0].sasl.password')
fi

if [[ ! -z "$KAFKA_USERNAME" ]] && [[ ! -z "$KAFKA_PW" ]]; then
    echo "Setting up Kafka auth config..."
    cat <<EOF > $KAFKA_AUTH_CONFIG
sasl.mechanism=SCRAM-SHA-512
security.protocol=SASL_SSL
sasl.jaas.config=org.apache.kafka.common.security.scram.ScramLoginModule required username="$KAFKA_USER" password="$KAFKA_PW";
EOF
fi

echo "Kafka Setup Complete!"
