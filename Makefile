POSTGRESQL_URL = postgresql://postgres:password@localhost:5432/undercontrol?sslmode=disable
BINARY_NAME = utlgo

.PHONY: init-local-db
init-local-db:
	./scripts/mk-local-db.sh

migrate-db-up:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up

migrate-db-down:
	migrate -database ${POSTGRESQL_URL} -path db/migrations down

migrate-db-new:
	read -r -p "Name of the migration: " name; \
	migrate create -ext sql -dir db/migrations -seq $$name

gen-db-type: migrate-db-up
	jet -dsn=${POSTGRESQL_URL} -schema=public -path=./.gen

.PHONY: build-image
build-image:
	docker build -t utlgo:$$(git rev-parse --short HEAD) .

build-local-binary:
	go build -o ~/.local/bin/$(BINARY_NAME) -trimpath main.go
	chmod u+x ~/.local/bin/$(BINARY_NAME)

setup-local: init-local-db migrate-db-up build-local-binary
	utlgo help
