# Proxy

At the heart of Digitox is a forward proxy that restricts access to domains listed in blocklists while a session is
active. The popular [Squid](http://www.squid-cache.org/) caching proxy is used to restrict access to blocked domains.

## Sessions

Sessions are stored as JSON files in `/etc/digitox/sessions/`. Sessions are associated with blocklists, which include
the list of domains to block.

## Blocklists

Blocklists are stored as JSON files in `/etc/digitox/blocklists/`. You can block a specific subdomain
(e.g. `www.reddit.com`) or all subdomains (e.g. `.reddit.com`). Each blocked domain must be listed on a separate line.

When a session starts, the included blocklists are copied to `/etc/squid/blocklist` and Squid denies access to
these domains resulting in a HTTP `403 Forbidden` error. When a session ends, `/etc/squid/blocklist` is removed
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

An Incron job monitors /etc/squid/blocklist for modifications and if a modification occurs, the job runs
`squid -k reconfigure` to reload the proxy configuration, which in turn reloads the blocklist.
