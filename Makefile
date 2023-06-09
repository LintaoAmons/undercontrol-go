.PHONY: init-local-db
init-local-db:
	./scripts/mk-local-db.sh

gen-db-type:
	jet -dsn=postgresql://postgres:password@localhost:5432/undercontrol?sslmode=disable -schema=public -path=./.gen
