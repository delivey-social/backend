.PHONY: dev test coverage migrate-up

dev:
	docker compose up backend-dev db

build:
	docker compose up backend-build db

test:
	go test ./...

coverage:
	go test -coverprofile cover.prof ./...
	covreport -o cover.html
	cmd.exe /c start cover.html

migrate-up:
	MIGRATION_ACTION=up docker compose run migrate

migrate-down:
	MIGRATION_ACTION=down docker compose run migrate