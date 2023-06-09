#!/bin/bash

docker run \
	--name undercontrol-db \
	-p 5432:5432 \
	-v undercontrol-db:/var/lib/postgresql/data \
	-e POSTGRES_PASSWORD=password \
	-d \
	postgres
