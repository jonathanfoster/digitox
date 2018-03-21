#!/bin/bash

curl --proxy 0.0.0.0:8080 -i -N \
    -H "Connection: Upgrade" \
    -H "Upgrade: websocket" \
    -H "Host: echo.websocket.org" \
    -H "Origin: https://echo.websocket.org" \
    https://echo.websocket.org
