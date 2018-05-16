# Proxy

At the heart of Digitox is a forward proxy that restricts access to domains listed in blocklists while a session is
active. The popular [Squid](http://www.squid-cache.org/) caching proxy is used to restrict access to blocked domains.

## Sessions

Sessions are stored in `/etc/digitox/sessions.db`. Sessions are associated with blocklists, which include
the list of domains to block.

## Blocklists

Blocklists are stored in `/etc/digitox/sessions.db`. You can block a specific subdomain(e.g. `www.reddit.com`) or all
subdomains (e.g. `.reddit.com`). Each blocked domain must be listed on a separate line.

When a session starts, the included blocklists are copied to `/etc/squid/blocklist` and Squid denies access to
these domains resulting in a HTTP `403 Forbidden` error. When a session ends, `/etc/squid/blocklist` is cleared
to grant access to all domains.

## Proxy Controller

The proxy controller is responsible for updating the proxy blocklist when the state of a session changes from active to
inactive.

The controller periodically polls for active sessions and constructs the expected blocklist. This expected
blocklist is then compared to the actual blocklist and adjusted accordingly. If a change is made, the proxy
automatically reloads the new configuration.

## Proxy Config Reload

The proxy configuration reload is implemented using [Incron](http://inotify.aiken.cz/?section=incron&page=about&lang=en).
Incron is a cron-style job system that is triggered on file system events instead of a time periods.

An Incron job monitors `/etc/squid/blocklist` for modifications and if a modification occurs, the job runs
`/usr/sbin/squid -k reconfigure` to reload the proxy configuration, which in turn reloads the blocklist. Proxy
configuration can be manually reloaded by making an HTTP request `POST /proxy/reconfigure` to the REST API.

Please note that reloading the proxy configuration as negative side effects such as closing the listening port and
interrupting in-transit requests. Configuration reloads should be minimized to avoid impact. See [Squid's hot configuration](https://wiki.squid-cache.org/Features/HotConf)
reload feature request for more details.