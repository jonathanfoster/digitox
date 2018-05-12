#!/bin/bash
set -e

CLIENT_ID=59f92849-b883-402c-b429-15a67663d4f3
CLIENT_SECRET=450a31ea-0c18-4925-97db-b9f981ca4a62

curl -s "http://localhost:8080/oauth/token?grant_type=client_credentials&client_id=${CLIENT_ID}&client_secret=${CLIENT_SECRET}&redirect_uri=http://localhost"