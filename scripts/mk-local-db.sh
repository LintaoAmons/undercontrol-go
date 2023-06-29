#!/bin/bash
container_name="undercontrol-db"

if docker ps -a --format '{{.Names}}' | grep -q "^$container_name$"; then
	echo "Container $container_name already exists."
else
	docker run \
		--name $container_name \
		-p 5432:5432 \
		-v undercontrol-db:/var/lib/postgresql/data \
		-e POSTGRES_PASSWORD=password \
		-d \
		postgres

	sleep 3
	PGPASSWORD=password psql -h localhost -U postgres -w -c "create database undercontrol;"
fi
