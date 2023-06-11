POSTGRESQL_URL = postgresql://postgres:password@localhost:5432/undercontrol?sslmode=disable

.PHONY: init-local-db
init-local-db:
	./scripts/mk-local-db.sh

migrate-db-up:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up

migrate-db-down:
	migrate -database ${POSTGRESQL_URL} -path db/migrations down

migrate-db-new:
	read -r -p "Name of the migration: " name
	# echo ${name}  TODO: makefile give user input
	# migrate create -ext sql -dir db/migrations -seq $name

gen-db-type: migrate-db-up
	jet -dsn=${POSTGRESQL_URL} -schema=public -path=./.gen
