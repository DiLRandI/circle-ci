#!/bin/bash

# Check if CIRCLECI_TOKEN is set
if [ -z "$CIRCLECI_TOKEN" ]; then
    echo "CIRCLECI_TOKEN is not set. Please set it in your environment."
    exit 1
fi

repo_url=$(git config --get remote.origin.url)
repo_name=$(echo "$repo_url" | sed -E -n 's/.*[\/:]([^\/]*\/[^\/]*).git$/\1/p')

# CircleCI API endpoint to get pending pipelines
api_endpoint="https://circleci.com/api/v2/pipeline?org-slug=$repo_name"

# Make the API request to get pending pipelines
response=$(curl -s -H "Circle-Token: $CIRCLECI_TOKEN" "$api_endpoint?status=pending")

# Check if the request was successful
if [ $? -eq 0 ]; then
    # Parse the response to extract relevant information
    pipeline_ids=$(echo "$response" | jq -r '.items[].id')

    if [ -n "$pipeline_ids" ]; then
        echo "Pending Pipelines:"

        for pipeline_id in $pipeline_ids; do
            echo "Pipeline ID: $pipeline_id"
        done
    else
        echo "No pending pipelines found."
    fi
else
    echo "Failed to retrieve pending pipelines."
    exit 1
fi
