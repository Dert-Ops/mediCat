include .env
export $(shell sed 's/=.*//' .env)

all:
	make up

init:
	docker network create medicat_network
	make up

up:
	make db
	make app

rre:
	make db-re
	make re

db:
	docker compose -f docker-compose-db.yml up -d

app:
	docker compose -f docker-compose-app.yml up -d

down:
	docker compose -f docker-compose-app.yml down

stop:
	docker compose -f docker-compose-app.yml down
	docker compose -f docker-compose-db.yml down

build:
	docker compose -f build docker-compose-app.yml
	docker compose -f build docker-compose-db.yml

db-re:
	docker compose -f docker-compose-db.yml down
	docker compose -f docker-compose-db.yml build
	docker compose -f docker-compose-db.yml up -d

re:
	docker compose -f docker-compose-app.yml down
	docker compose -f docker-compose-app.yml build
	docker compose -f docker-compose-app.yml up -d

