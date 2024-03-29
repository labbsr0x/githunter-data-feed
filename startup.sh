#!/bin/bash
set -e
set -x

echo "Starting GitHunter Data Feed..."
githunter \
    --log-level=$GH_LOG_LEVEL \
    --graphql-github-url=$GH_GRAPHQL_GITHUB_URL \
    --graphql-gitlab-url=$GH_GRAPHQL_GITLAB_URL \
    --api-gitlab-url=$GH_API_GITLAB_URL \
    --server-port=$GH_SERVER_PORT