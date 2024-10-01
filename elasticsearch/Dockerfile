FROM elasticsearch:8.9.1

ENV discovery.type=single-node

EXPOSE 9200

ENTRYPOINT ["/bin/tini", "--", "/usr/local/bin/docker-entrypoint.sh"]
