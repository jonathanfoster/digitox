#!/bin/sh
set -e

/sbin/syslogd
/usr/sbin/incrond
/usr/sbin/squid
/usr/local/bin/digitox-apiserver $@
