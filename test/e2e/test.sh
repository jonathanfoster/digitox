#!/bin/bash
set -e

command -v jq >/dev/null 2>&1 || { echo "jq required but not installed." >&2; exit 1; }

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

###########################
#    Access Token Test    #
###########################

echo "Getting access token"
export DIGITOX_ACCESS_TOKEN=$($DIR/oauth/get-token.sh | jq -r '.access_token')
echo "Access token: $DIGITOX_ACCESS_TOKEN"

######################
#    Device Tests    #
######################

echo "Creating test device"
DEVICE=$($DIR/device/create-device.sh $(uuidgen) $(uuidgen))
echo $DEVICE | jq
DEVICE_NAME=$(echo $DEVICE | jq -r '.name')
DEVICE_PASSWORD=$(echo $DEVICE | jq -r '.password')
echo ""

echo "Listing test devices"
$DIR/device/list-devices.sh | jq
echo ""

echo "Finding test device $DEVICE_NAME"
$DIR/device/find-device.sh $DEVICE_NAME | jq
echo ""

echo "Updating test device $DEVICE_NAME"
DEVICE=$($DIR/device/update-device.sh ${DEVICE_NAME} $(uuidgen))
echo $DEVICE | jq
DEVICE_PASSWORD=$(echo $DEVICE | jq -r '.password')
# Sleep for 1 second to allow Squid restart to complete
sleep 1
echo ""

echo "Asserting blocked domain allowed before test session created"
$DIR/proxy/proxy-allow-blocked.sh $DEVICE_NAME $DEVICE_PASSWORD
echo ""

#########################
#    Blocklist Tests    #
#########################

echo "Creating test blocklist"
BLOCKLIST=$($DIR/blocklist/create-blocklist.sh $(uuidgen) news.ycombinator.com)
echo $BLOCKLIST | jq
BLOCKLIST_ID=$(echo $BLOCKLIST | jq -r '.id')
echo ""

echo "Listing test blocklists"
$DIR/blocklist/list-blocklists.sh | jq
echo ""

echo "Finding test blocklist $BLOCKLIST_ID"
$DIR/blocklist/find-blocklist.sh $BLOCKLIST_ID | jq
echo ""

echo "Updating test blocklist $BLOCKLIST_ID"
$DIR/blocklist/update-blocklist.sh $BLOCKLIST_ID $(uuidgen) news.ycombinator.com | jq
echo ""

#######################
#    Session Tests    #
#######################

echo "Creating test session"
SESSION=$($DIR/session/create-session.sh $(uuidgen) $BLOCKLIST_ID)
echo $SESSION | jq
SESSION_ID=$(echo $SESSION | jq -r '.id')
echo ""

echo "Listing test sessions"
$DIR/session/list-sessions.sh | jq
echo ""

echo "Finding test session $SESSION_ID"
$DIR/session/find-session.sh $SESSION_ID | jq
echo ""

echo "Updating test session $SESSION_ID"
$DIR/session/update-session.sh $SESSION_ID $(uuidgen) $BLOCKLIST_ID | jq
# Sleep for 1 seconds to allow proxy controller to update blocklist and Squid restart to complete
sleep 1
echo ""

echo "Asserting blocked domain denied while session active"
$DIR/proxy/proxy-deny-blocked.sh $DEVICE_NAME $DEVICE_PASSWORD
echo ""

echo "Asserting not blocked domain allowed while session active"
$DIR/proxy/proxy-allow.sh $DEVICE_NAME $DEVICE_PASSWORD
echo ""

#######################
#    Removal Tests    #
#######################

echo "Removing test blocklist $BLOCKLIST_ID"
$DIR/blocklist/remove-blocklist.sh ${BLOCKLIST_ID}
echo ""

echo "Removing test session $SESSION_ID"
$DIR/session/remove-session.sh ${SESSION_ID}
# Sleep for 1 seconds to allow proxy controller to update blocklist and Squid restart to complete
sleep 1
echo ""

echo "Asserting blocked domain allowed after session expires"
$DIR/proxy/proxy-allow-blocked.sh $DEVICE_NAME $DEVICE_PASSWORD
echo ""

echo "Removing test device $DEVICE_NAME"
$DIR/device/remove-device.sh ${DEVICE_NAME}
echo ""
