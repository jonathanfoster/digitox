#!/bin/bash
set -e

curl -i 'http://localhost:8080/oauth/token?grant_type=client_credentials&client_id=admin&client_secret=Digitox123&redirect_uri=http://localhost'
