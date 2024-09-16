FROM postgres:14

COPY ./slave/postgresql.conf /etc/postgresql/postgresql.conf
COPY ./slave/pg_hba.conf /etc/postgresql/pg_hba.conf
COPY ./slave/setup-slave.sh /docker-entrypoint-initdb.d/

CMD ["postgres", "-c", "config_file=/etc/postgresql/postgresql.conf"]
