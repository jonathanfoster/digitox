# REST API

## Resources

Digitox REST API provides the following resources:

* Session
* Blocklist
* Device
* Proxy

## Authorization

Requests to the Digitox REST API require an access token, which can be granted using the client credentials OAuth flow.
The access token must be provided in either the `access_token` query string parameter
(e.g. `?access_token=$DIGITOX_ACCESS_TOKEN`) or in the `Authorization` header
(e.g. `Authorization: Bearer $DIGITOX_ACCESS_TOKEN`). See [Security](./security.md) for more information.

```bash
# Get access token
GET /oauth/token?grant_type=client_credentials&client_id=$DIGITOX_CLIENT_ID&client_secret=$DIGITOX_CLIENT_SECRET&redirect_uri=http://localhost
```

## Endpoints

### Sessions

Session endpoints provide CRUD operations on session resources.

```bash
# List sessions
GET /sessions/

# Find session
GET /sessions/{id}

# Create session
POST /sessions/

# Update session
PUT /sessions/{id}

# Remove session
DELETE /sessions/{id}
```

### Blocklists

Blocklist endpoints provide CRUD operations on blocklist resources.

```bash
# List blocklists
GET /blocklists/

# Find blocklist
GET /blocklists/{id}

# Create blocklist
POST /blocklists/

# Update blocklist
PUT /blocklists/{id}

# Remove blocklist
DELETE /blocklists/{id}
```

### Devices

Device endpoints provide CRUD operations on device resources.

```bash
# List devices
GET /devices/

# Find devices
GET /devices/{name}

# Create device
POST /devices/

# Update device
PUT /devices/{name}

# Remove device
DELETE /devices/{name}
```

### Proxy

Proxy endpoints are useful in debugging issues related to the proxy blocklist.

```bash
# Get active proxy blocklist
GET /proxy/active

# Get expected blocklist based on current session (should match active blocklist under normal circumstances)
GET /proxy/session

# Reload proxy configuration
POST /proxy/reconfigure
```
