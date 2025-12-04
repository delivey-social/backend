.PHONY: dev test coverage

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
