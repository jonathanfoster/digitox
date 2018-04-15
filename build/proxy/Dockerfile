FROM alpine:latest

RUN apk add --no-cache squid incron && \
    echo "" > /etc/squid/blocklist && \
    echo "root" > /etc/incron.allow && \
    echo "/etc/squid/blocklist IN_MODIFY /usr/sbin/squid -k reconfigure" > /var/spool/incron/root

COPY squid.conf /etc/squid/
COPY entrypoint.sh /usr/local/bin/

RUN chmod +x /usr/local/bin/entrypoint.sh

EXPOSE 3128

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
CMD ["-NYCd", "1"]
