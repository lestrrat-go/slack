#!/bin/sh

if [[ -z "$SLACK_TOKEN" ]]; then
    echo "You must specify SLACK_TOKEN environment variable with a valid"
    echo "token from Slack"
    exit 1
fi

if [[ -z "$SLACK_LISTEN" ]]; then
    SLACK_LISTEN=:8080
fi

exec /slaproxy -token "$SLACK_TOKEN" -listen "$SLACK_LISTEN"