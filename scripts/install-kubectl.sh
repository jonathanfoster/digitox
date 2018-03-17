#!/bin/bash
set -e

curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.9.4/bin/linux/amd64/kubectl
chmod +x ./kubectl
sudo mv ./kubectl /usr/local/bin/kubectl
