# Security

Digitox uses a combination of [OAuth 2](https://tools.ietf.org/html/rfc6749) and [JWT tokens](https://tools.ietf.org/html/rfc7519)
for authorization for the API server and HTTP basic authentication for the proxy server.

## API Server

### Authorization

A client can request an access token using the [OAuth 2.0 Client Credentials grant flow](https://tools.ietf.org/html/rfc6749#section-4.4).

```bash
curl "http://localhost:8080/oauth/token?grant_type=client_credentials&client_id=${CLIENT_ID}&client_secret=${CLIENT_SECRET}&redirect_uri=http://localhost"
```

No other OAuth grant types are currently supported.

#### Client Credentials

Only one set of client credentials is supported and is set on startup using either the flags `--client-id` and
`--client-secret` or the environment variables `DIGITOX_CLIENT_ID` and `DIGITOX_CLIENT_SECRET`. If no credentials are
provided then default values will be used. These values are output to standard out on startup in case they're not
already known.

The redirect URI must be `http://localhost`.

### Tokens

An access token must be provided in either the Authorization header or in the access_token query string parameter for
most resource requests.

```bash
curl "http://localhost:8080/sessions/?access_token=$DIGITOX_ACCESS_TOKEN"
```


```bash
curl -H "Authorization: Bearer $DIGITOX_ACCESS_TOKEN" http://localhost:8080/sessions/
```

Digitox access tokens are JWT tokens that include claims for CID (i.e. client ID) and expiration (3,600 secords or 1
hour). Tokens are signed using the RSA256 algorithm.

#### Signing

Tokens are signed using an RSA public/private key pair. The private key is used for signing and the public key is used
for verification.

The key pair paths can be set on startup using the flags `signing-key` and `verifying-key` or the environment variables
`DIGITOX_SIGNING_KEY` and `DIGITOX_VERIFYING_KEY`. If no key pair paths are provided then a default key pair will be used.

You can generate your own key pair using the `token-keygen.sh` script.

```bash
./scripts/token-keygen.sh signing-key.pem verifying-key.pem
```

## Proxy Server

The proxy server is configured to use HTTP basic authentication and the credentials are stored in `/etc/digitox/passwd`
by default. Credentials are managed as device resources in the REST API. See [API doc](./api.md) for more information.
