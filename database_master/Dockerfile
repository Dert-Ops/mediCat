FROM postgres:14

COPY ./master/postgresql.conf /etc/postgresql/postgresql.conf
COPY ./master/pg_hba.conf /etc/postgresql/pg_hba.conf
COPY ./master/setup-master.sh /docker-entrypoint-initdb.d/

CMD ["postgres", "-c", "config_file=/etc/postgresql/postgresql.conf"]
