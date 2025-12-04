.PHONY: dev test coverage

dev:
	docker compose up

test:
	go test ./...

coverage:
	go test -coverprofile cover.prof ./...
	covreport -o cover.html
	cmd.exe /c start cover.html
