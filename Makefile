.PHONY: test coverage

build:
	docker build -t comida-app-backend .

run:
	docker run -it --rm -p 3001:3001  --name comida-app-backend-container comida-app-backend

test:
	go test ./...

coverage:
	go test -coverprofile cover.prof ./...
	covreport -o cover.html
	cmd.exe /c start cover.html
