include .env
export $(shell sed 's/=.*//' .env)

all:

up:
	docker compose up -d

down:
	docker compose down

build:
	docker compose build

logs:
	docker compose logs -f

re:
	docker compose down
	docker compose build
	docker compose up -d

init-replication:
	docker exec -it postgres-master psql -U ${POSTGRES_USER} -c "CREATE ROLE ${REPLICATION_USER} WITH REPLICATION PASSWORD '${REPLICATION_PASSWORD}' LOGIN;"
	docker exec -it postgres-slave bash -c "PGPASSWORD=${REPLICATION_PASSWORD} pg_basebackup -h postgres-master -D /var/lib/postgresql/data -U ${REPLICATION_USER} -v -P --wal-method=stream"
