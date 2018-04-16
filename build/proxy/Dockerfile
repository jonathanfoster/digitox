FROM alpine:latest

RUN apk add --no-cache squid incron && \
    mkdir -p /etc/digitox && \
    echo "" > /etc/digitox/active && \
    echo "root" > /etc/incron.allow && \
    echo "/etc/digitox/active IN_MODIFY /usr/sbin/squid -k reconfigure" > /var/spool/incron/root

COPY squid.conf /etc/squid/
COPY entrypoint.sh /usr/local/bin/

RUN chmod +x /usr/local/bin/entrypoint.sh

EXPOSE 3128

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
CMD ["-NYCd", "1"]
