#!/bin/bash
set -e

command -v jq >/dev/null 2>&1 || { echo "jq required but not installed." >&2; exit 1; }

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

export DIGITOX_ACCESS_TOKEN=$($DIR/get-token.sh | jq -r '.access_token')
echo $DIGITOX_ACCESS_TOKEN
