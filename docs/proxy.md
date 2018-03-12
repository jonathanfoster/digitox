# Proxy

At the heart of Freedom is a forward proxy that restricts access to sites listed in block lists while a session is
active. The popular [Squid](http://www.squid-cache.org/) caching proxy is used along with cron-style jobs, access
control lists, and basic authentication to manage sessions, block lists, and devices.

## Sessions

Sessions are managed as cron-style jobs.

## Block Lists

Block lists are stored as flat files in `/etc/squid/blocklists/`. You can block a specific subdomain
(e.g. `www.reddit.com`) or all subdomains (e.g. `.reddit.com`). Each blocked domain must be listed on a separate line.

When a session starts, the included block lists are copied to `/etc/squid/blocklists/block`, and Squid denies access to
these domains resulting in a HTTP `403 Forbidden` error. When a session ends, `/etc/squid/blocklists/block` is removed
to grant access to all domains.

## Devices

Devices are managed as a set of credentials stored in `/etc/squid/passwd`. To create a new device, a call is made
`htpasswd` creating a set of credentials that includes the device name as the username and a secret key as the
password. The user must then configure a proxy connection on their device using the Squid proxy URL and credentials.
