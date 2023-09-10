#!/bin/bash

# Check if a single argument (parameter) is provided
if [ $# -ne 1 ]; then
    echo "Usage: $0 <parameter>"
    exit 1
fi

parameter="$1"

# Get the current Git branch
current_branch=$(git symbolic-ref --short HEAD 2>/dev/null)

if [ -z "$current_branch" ]; then
    echo "Error: Not in a Git repository or unable to determine the current branch."
    exit 1
fi

# Check if the current directory is a Git repository
if [ -n "$current_branch" ]; then
    # Get the Git repository URL (HTTPS or SSH)
    repo_url=$(git config --get remote.origin.url)

    if [ -n "$repo_url" ]; then
        # Extract the user|org/repositoryName part
        repo_name=$(echo "$repo_url" | sed -E -n 's/.*[\/:]([^\/]*\/[^\/]*).git$/\1/p')

        if [ -n "$repo_name" ]; then
            echo "You are in a Git repository."
            echo "Repository Name: $repo_name"
            echo "Current Branch: $current_branch"

            # Check if CIRCLECI_TOKEN is set
            if [ -n "$CIRCLECI_TOKEN" ]; then
                echo "CIRCLECI_TOKEN is set."

                # Make the curl request
                curl --request POST \
                    --url "https://circleci.com/api/v2/project/github/$repo_name/pipeline" \
                    --header "Circle-Token: $CIRCLECI_TOKEN" \
                    --data "{\"branch\":\"$current_branch\"}"
            else
                echo "CIRCLECI_TOKEN is not set. Please set it in your environment."
            fi
        else
            echo "Unable to determine the repository name."
        fi
    else
        echo "Unable to determine the repository URL."
    fi
else
    echo "You are not in a Git repository."
fi
