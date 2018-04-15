#!/bin/bash
set -e

curl -LO https://github.com/kubernetes/kops/releases/download/1.9.0/kops-linux-amd64
chmod +x kops-linux-amd64
sudo mv kops-linux-amd64 /usr/local/bin/kops
