#!/bin/sh
set -e

/usr/sbin/incrond
/usr/sbin/squid $@
