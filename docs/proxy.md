# Proxy

At the heart of Freedom is a forward proxy that restricts access to sites listed in blocklists while a session is
active. The popular [Squid](http://www.squid-cache.org/) caching proxy to restrict access to blocked domains.

## Sessions

Sessions are stored as JSON files in `/etc/freedom/sessions/`. Sessions are associated with blocklists, which include
the list of domains to block.

## Blocklists

Blocklists are stored as JSON files in `/etc/freedom/blocklists/`. You can block a specific subdomain
(e.g. `www.reddit.com`) or all subdomains (e.g. `.reddit.com`). Each blocked domain must be listed on a separate line.

When a session starts, the included blocklists are copied to `/etc/squid/blocklists`, and Squid denies access to
these domains resulting in a HTTP `403 Forbidden` error. When a session ends, `/etc/squid/blocklists` is removed
to grant access to all domains.

## Proxy Controller

The proxy controller is responsible for updating the proxy blocklist when the state of a session changes from active to
inactive. The controller periodically polls for active sessions and constructs the expected blocklist. This expected
blocklist is then compared to the actual blocklist and adjusted accordingly. If a change was made, the controller
signals the proxy to reload the new configuration.
