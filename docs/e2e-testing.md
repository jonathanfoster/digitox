# End-to-End Testing

End-to-end testing can be performed using the E2E shell scripts in `test/e2e/`.

## Prerequisites

* [jq](https://stedolan.github.io/jq/)

## Getting Started

Before getting started, you'll need to grant execute permissions to all E2E test scripts.

```bash
chmod +x ./test/e2e/**
```

## Full End-to-End Test

You can run a full e2e test using `test.sh`.

```bash
./test/e2e/test.sh
```

You can also use the Make task `make test-e2e`.

The full test will perform the following:

* Get access token
* Create test device
* List test devices
* Find test device
* Update test device
* Assert all domains not blocked before test session created
* Create test blocklist
* List test blocklists
* Find test blocklist
* Update test blocklist
* Create test session
* List test sessions
* Find test session
* Update test session
* Assert test domain blocked while session active
* Assert non-test domain not blocked while session active
* Remove test blocklist
* Remove test session
* Assert all domains not blocked after session expires
* Remove test device

## Access Token

Digitox REST API requires an OAuth access token. The E2E test scripts assume an access token is provided in the
`$DIGITOX_ACCESS_TOKEN` environment variable. Use `set-env-token.sh` to authenticate and set the environment variable.

```bash
./test/e2e/oauth/set-env-token.sh
```

You can also get a token by calling `get-token.sh`. This will authenticate and write the token to standard out.

```bash
./test/e2e/oauth/get-token.sh
```

Both scripts assume Digitox is running with the default client ID and secret.

## Blocklists

Blocklist E2E test scripts are found in `test/e2e/blocklist/`.

## Devices

Device E2E test scripts are found in `test/e2e/device/`.

## OAuth

OAuth E2E test scripts are found in `test/e2e/oauth/`.

## Proxy

Proxy E2E test scripts are found in `test/e2e/proxy/`.

## Sessions

Session E2E test scripts are found in `test/e2e/session/`.
