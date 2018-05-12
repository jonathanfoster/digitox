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
echo "Test device name: $DEVICE_NAME"
echo "Test device password: $DEVICE_PASSWORD"
echo ""

echo "Listing test devices"
$DIR/device/list-devices.sh | jq
echo ""

echo "Finding test device $DEVICE_NAME"
$DIR/device/find-device.sh $DEVICE_NAME | jq
echo ""

echo "Updating test device"
DEVICE=$($DIR/device/create-device.sh ${DEVICE_NAME} $(uuidgen))
echo $DEVICE | jq
DEVICE_PASSWORD=$(echo $DEVICE | jq -r '.password')
echo "New test device password: $DEVICE_PASSWORD"
echo ""

echo "Asserting all domains not blocked before test session created"
$DIR/proxy/proxy-deny.sh $DEVICE_NAME $DEVICE_PASSWORD
echo ""

#########################
#    Blocklist Tests    #
#########################

# Creating test blocklist
# Listing test blocklists
# Finding test blocklist
# Updating test blocklist

#######################
#    Session Tests    #
#######################

# Creating test session
# Listing test sessions
# Finding test session
# Updating test session

# Asserting test domain blocked while session active
# Asserting non-test domain not blocked while session active

#######################
#    Removal Tests    #
#######################

# Removing test blocklist
# Removing test session
# Asserting all domains not blocked after session expires

echo "Removing test device"
$DIR/device/remove-device.sh ${DEVICE_NAME}
